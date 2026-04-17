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

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "rm: %v\n", err)
		return 1
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

		if info.IsDir() && !doRecursive {
			fmt.Fprintf(env.Stderr, "rm: cannot remove '%s': Is a directory\n", target)
			exitCode = 1
			continue
		}

		if info.IsDir() {
			err = env.FS.RemoveAll(fullPath)
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
