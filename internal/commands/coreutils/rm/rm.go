package rm

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Rm struct{}

func New() *Rm {
	return &Rm{}
}

func (r *Rm) Name() string {
	return "rm"
}

func (r *Rm) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("rm", pflag.ContinueOnError)
	recursive := flags.BoolP("recursive", "r", false, "remove directories and their contents recursively")
	recursiveUpper := flags.BoolP("recursive-upper", "R", false, "identical to -r")
	force := flags.BoolP("force", "f", false, "ignore nonexistent files and arguments, never prompt")
	verbose := flags.BoolP("verbose", "v", false, "explain what is being done")
	dir := flags.BoolP("dir", "d", false, "remove empty directories")
	_ = flags.BoolP("interactive", "i", false, "prompt before every removal (stub)")
	_ = flags.Bool("one-file-system", false, "when removing a hierarchy recursively, skip any directory that is on a file system different from that of the corresponding command line argument (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "rm: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: rm [OPTION]... [FILE]...\n")
		fmt.Fprintf(env.Stdout, "Remove (unlink) the FILE(ies).\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "rm")
		return 0
	}

	targets := flags.Args()
	if len(targets) == 0 && !*force {
		fmt.Fprintf(env.Stderr, "rm: missing operand\n")
		return 1
	}

	doRecursive := *recursive || *recursiveUpper
	exitCode := 0

	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		info, err := env.FS.Stat(fullPath)
		if err != nil {
			if !*force {
				fmt.Fprintf(env.Stderr, "rm: cannot remove '%s': No such file or directory\n", target)
				exitCode = 1
			}
			continue
		}

		if info.IsDir() && !doRecursive && !*dir {
			fmt.Fprintf(env.Stderr, "rm: cannot remove '%s': Is a directory\n", target)
			exitCode = 1
			continue
		}

		if info.IsDir() {
			if doRecursive {
				err = env.FS.RemoveAll(fullPath)
			} else if *dir {
				err = env.FS.Remove(fullPath)
			}
		} else {
			err = env.FS.Remove(fullPath)
		}

		if err != nil {
			fmt.Fprintf(env.Stderr, "rm: cannot remove '%s': %v\n", target, err)
			exitCode = 1
			continue
		}

		if *verbose {
			fmt.Fprintf(env.Stdout, "removed '%s'\n", target)
		}
	}

	return exitCode
}
