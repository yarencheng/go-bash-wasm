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

type duOptions struct {
	humanReadable bool
	summarize     bool
	all           bool
	apparentSize  bool
	maxDepth      int
	kilobytes     bool
	megabytes     bool
	null          bool
}

func (d *Du) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("du", pflag.ContinueOnError)
	humanReadable := flags.BoolP("human-readable", "h", false, "print sizes in human readable format")
	summarize := flags.BoolP("summarize", "s", false, "display only a total for each argument")
	all := flags.BoolP("all", "a", false, "write counts for all files")
	apparentSize := flags.BoolP("apparent-size", "b", false, "print apparent sizes")
	totalFlag := flags.BoolP("total", "c", false, "produce a grand total")
	maxDepth := flags.IntP("max-depth", "d", -1, "max depth to print")
	kilobytes := flags.BoolP("kilobytes", "k", false, "1K blocks")
	megabytes := flags.BoolP("megabytes", "m", false, "1M blocks")
	null := flags.BoolP("null", "0", false, "null terminated")
	_ = flags.BoolP("dereference", "L", false, "dereference")
	_ = flags.BoolP("no-dereference", "P", false, "no dereference")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "du: %v\n", err)
		return 1
	}

	opts := duOptions{
		humanReadable: *humanReadable,
		summarize:     *summarize,
		all:           *all,
		apparentSize:  *apparentSize,
		maxDepth:      *maxDepth,
		kilobytes:     *kilobytes,
		megabytes:     *megabytes,
		null:          *null,
	}

	targets := flags.Args()
	if len(targets) == 0 {
		targets = append(targets, ".")
	}

	var grandTotal int64
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		size, err := d.calculateSize(env.FS, fullPath, 0, opts, env)
		if err != nil {
			fmt.Fprintf(env.Stderr, "du: %v\n", err)
			continue
		}
		grandTotal += size

		if opts.summarize {
			d.printSize(env.Stdout, size, target, opts)
		}
	}

	if *totalFlag {
		d.printSize(env.Stdout, grandTotal, "total", opts)
	}

	return 0
}

func (d *Du) calculateSize(fs afero.Fs, path string, depth int, opts duOptions, env *commands.Environment) (int64, error) {
	info, err := fs.Stat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		if (opts.all || depth == 0) && !opts.summarize {
			if opts.maxDepth == -1 || depth <= opts.maxDepth {
				d.printSize(env.Stdout, info.Size(), path, opts)
			}
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
		size, _ := d.calculateSize(fs, subPath, depth+1, opts, env)
		total += size
	}

	if !opts.summarize {
		if opts.maxDepth == -1 || depth <= opts.maxDepth {
			d.printSize(env.Stdout, total, path, opts)
		}
	}

	return total, nil
}

func (d *Du) printSize(w io.Writer, size int64, path string, opts duOptions) {
	sep := "\n"
	if opts.null {
		sep = "\x00"
	}

	if opts.humanReadable {
		fmt.Fprintf(w, "%s\t%s%s", d.formatSize(size), path, sep)
	} else {
		divisor := int64(1024)
		if opts.apparentSize {
			divisor = 1
		} else if opts.megabytes {
			divisor = 1024 * 1024
		} else if opts.kilobytes {
			divisor = 1024
		}
		fmt.Fprintf(w, "%d\t%s%s", (size+divisor-1)/divisor, path, sep)
	}
}

func (d *Du) formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(size)/float64(div), "KMGTPE"[exp])
}
