package du

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strconv"

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
	si            bool
	summarize     bool
	all           bool
	apparentSize  bool
	maxDepth      int
	kilobytes     bool
	megabytes     bool
	null          bool
	threshold     int64
	inodes        bool
}

func (d *Du) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("du", pflag.ContinueOnError)
	humanReadable := flags.BoolP("human-readable", "h", false, "print sizes in human readable format")
	si := flags.Bool("si", false, "like -h but use powers of 1000 not 1024")
	summarize := flags.BoolP("summarize", "s", false, "display only a total for each argument")
	all := flags.BoolP("all", "a", false, "write counts for all files")
	apparentSize := flags.BoolP("apparent-size", "b", false, "print apparent sizes")
	totalFlag := flags.BoolP("total", "c", false, "produce a grand total")
	maxDepth := flags.IntP("max-depth", "d", -1, "max depth to print")
	kilobytes := flags.BoolP("kilobytes", "k", false, "1K blocks")
	megabytes := flags.BoolP("megabytes", "m", false, "1M blocks")
	null := flags.BoolP("null", "0", false, "null terminated")
	thresholdStr := flags.StringP("threshold", "t", "0", "exclude entries smaller than SIZE if positive, or entries larger than SIZE if negative")
	inodes := flags.Bool("inodes", false, "list inode usage information instead of block usage")
	_ = flags.BoolP("dereference", "L", false, "dereference")
	_ = flags.BoolP("no-dereference", "P", false, "no dereference")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "du: %v\n", err)
		return 1
	}

	threshold, _ := strconv.ParseInt(*thresholdStr, 10, 64)

	opts := duOptions{
		humanReadable: *humanReadable,
		si:            *si,
		summarize:     *summarize,
		all:           *all,
		apparentSize:  *apparentSize,
		maxDepth:      *maxDepth,
		kilobytes:     *kilobytes,
		megabytes:     *megabytes,
		null:          *null,
		threshold:     threshold,
		inodes:        *inodes,
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

	val := info.Size()
	if opts.inodes {
		val = 1 // Simplified inode count
	}

	if !info.IsDir() {
		if (opts.all || depth == 0) && !opts.summarize {
			if opts.maxDepth == -1 || depth <= opts.maxDepth {
				d.printSize(env.Stdout, val, path, opts)
			}
		}
		return val, nil
	}

	var total int64
	total = val // include self size? coreutils du does block count, usually directory has some size.
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
	// Check threshold
	if opts.threshold > 0 && size < opts.threshold {
		return
	}
	if opts.threshold < 0 && size > -opts.threshold {
		return
	}

	sep := "\n"
	if opts.null {
		sep = "\x00"
	}

	if opts.humanReadable || opts.si {
		fmt.Fprintf(w, "%s\t%s%s", d.formatSize(size, opts.si), path, sep)
	} else {
		divisor := int64(1024)
		if opts.apparentSize || opts.inodes {
			divisor = 1
		} else if opts.megabytes {
			divisor = 1024 * 1024
		} else if opts.kilobytes {
			divisor = 1024
		}
		val := (size + divisor - 1) / divisor
		if val == 0 && size > 0 {
			val = 1
		}
		fmt.Fprintf(w, "%d\t%s%s", val, path, sep)
	}
}

func (d *Du) formatSize(size int64, si bool) string {
	unit := float64(1024)
	if si {
		unit = 1000
	}
	if float64(size) < unit {
		return fmt.Sprintf("%d", size)
	}
	f := float64(size)
	exp := 0
	for f >= unit {
		f /= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", f, "KMGTPE"[exp-1])
}
