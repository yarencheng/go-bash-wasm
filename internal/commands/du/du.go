package du

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Du struct{}

func New() *Du {
	return &Du{}
}

func (d *Du) Name() string {
	return "du"
}

func (d *Du) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("du", pflag.ContinueOnError)
	humanReadable := flags.BoolP("human-readable", "h", false, "print sizes in human readable format (e.g., 1K 234M 2G)")
	summarize := flags.BoolP("summarize", "s", false, "display only a total for each argument")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "du: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		targets = append(targets, ".")
	}

	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		totalSize, err := d.calculateSize(env.FS, fullPath, *summarize, *humanReadable, env)
		if err != nil {
			fmt.Fprintf(env.Stderr, "du: %v\n", err)
			continue
		}

		if *summarize {
			d.printSize(env.Stdout, totalSize, target, *humanReadable)
		}
	}

	return 0
}

func (d *Du) calculateSize(fs afero.Fs, path string, summarize, human bool, env *commands.Environment) (int64, error) {
	info, err := fs.Stat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		if !summarize {
			d.printSize(env.Stdout, info.Size(), path, human)
		}
		return info.Size(), nil
	}

	var total int64
	entries, err := afero.ReadDir(fs, path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		subPath := filepath.Join(path, entry.Name())
		size, _ := d.calculateSize(fs, subPath, summarize, human, env)
		total += size
	}

	if !summarize {
		d.printSize(env.Stdout, total, path, human)
	}

	return total, nil
}

func (d *Du) printSize(w io.Writer, size int64, path string, human bool) {
	if human {
		fmt.Fprintf(w, "%s\t%s\n", d.formatSize(size), path)
	} else {
		fmt.Fprintf(w, "%d\t%s\n", size/1024, path) // du usually reports in blocks or KB
	}
}

func (d *Du) formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(size)/float64(div), "KMGTPE"[exp])
}
