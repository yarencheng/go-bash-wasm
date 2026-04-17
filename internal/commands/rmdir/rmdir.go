package rmdir

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Rmdir struct{}

func New() *Rmdir {
	return &Rmdir{}
}

func (r *Rmdir) Name() string {
	return "rmdir"
}

func (r *Rmdir) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("rmdir", pflag.ContinueOnError)
	parents := flags.BoolP("parents", "p", false, "remove each directory and its ancestors")
	verbose := flags.BoolP("verbose", "v", false, "output a diagnostic for every directory processed")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "rmdir: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "rmdir: missing operand\n")
		return 1
	}

	exitCode := 0
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		info, err := env.FS.Stat(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "rmdir: failed to remove '%s': %v\n", target, err)
			exitCode = 1
			continue
		}

		if !info.IsDir() {
			fmt.Fprintf(env.Stderr, "rmdir: '%s': Not a directory\n", target)
			exitCode = 1
			continue
		}

		// Check if empty
		f, err := env.FS.Open(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "rmdir: failed to remove '%s': %v\n", target, err)
			exitCode = 1
			continue
		}
		names, err := f.Readdirnames(1)
		f.Close()
		if err != nil && err != io.EOF {
			fmt.Fprintf(env.Stderr, "rmdir: failed to remove '%s': %v\n", target, err)
			exitCode = 1
			continue
		}
		if len(names) > 0 {
			fmt.Fprintf(env.Stderr, "rmdir: failed to remove '%s': Directory not empty\n", target)
			exitCode = 1
			continue
		}

		err = env.FS.Remove(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "rmdir: failed to remove '%s': %v\n", target, err)
			exitCode = 1
			continue
		}

		if *verbose {
			fmt.Fprintf(env.Stdout, "rmdir: removing directory, '%s'\n", target)
		}

		if *parents {
			parent := filepath.Dir(fullPath)
			for parent != "/" && parent != "." {
				err := env.FS.Remove(parent)
				if err != nil {
					break // Stop if parent is not empty or other error
				}
				if *verbose {
					fmt.Fprintf(env.Stdout, "rmdir: removing directory, '%s'\n", parent)
				}
				parent = filepath.Dir(parent)
			}
		}
	}

	return exitCode
}
