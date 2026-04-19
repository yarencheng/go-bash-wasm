package readlink

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Readlink struct{}

func New() *Readlink {
	return &Readlink{}
}

func (r *Readlink) Name() string {
	return "readlink"
}

func (r *Readlink) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("readlink", pflag.ContinueOnError)
	noNewline := flags.BoolP("no-newline", "n", false, "do not output the trailing newline")
	zero := flags.BoolP("zero", "z", false, "end each output line with NUL, not newline")
	quiet := flags.BoolP("quiet", "q", false, "do not print error messages (default)")
	silent := flags.BoolP("silent", "s", false, "same as -q")
	verbose := flags.BoolP("verbose", "v", false, "report error messages")
	canonicalize := flags.BoolP("canonicalize", "f", false, "canonicalize by following every symlink")
	canonicalizeExisting := flags.BoolP("canonicalize-existing", "e", false, "canonicalize by following every symlink; all components must exist")
	canonicalizeMissing := flags.BoolP("canonicalize-missing", "m", false, "canonicalize by following every symlink; missing components are ignored")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil && (*verbose || !(*quiet || *silent)) {
			fmt.Fprintf(env.Stderr, "readlink: %v\n", err)
		}
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		return 1
	}

	target := targets[0]
	if !filepath.IsAbs(target) {
		target = filepath.Join(env.Cwd, target)
	}

	var result string
	var err error

	if *canonicalize || *canonicalizeExisting || *canonicalizeMissing {
		fullPath := target
		if *canonicalizeExisting {
			_, err = env.FS.Stat(fullPath)
			if err != nil {
				if env.Stderr != nil && *verbose {
					fmt.Fprintf(env.Stderr, "readlink: %s: %v\n", target, err)
				}
				return 1
			}
		}

		// Simplified canonicalization for MemMapFs
		result = filepath.Clean(fullPath)
		// TODO: follow symlinks with afero once we have a good helper
	} else {
		linker, ok := env.FS.(afero.Symlinker)
		if !ok {
			if env.Stderr != nil && *verbose {
				fmt.Fprintln(env.Stderr, "readlink: filesystem does not support symlinks")
			}
			return 1
		}
		result, err = linker.ReadlinkIfPossible(target)
		if err != nil {
			if env.Stderr != nil && *verbose {
				fmt.Fprintf(env.Stderr, "readlink: %s: %v\n", target, err)
			}
			return 1
		}
		if result == target {
			// Not a symlink
			return 1
		}
	}

	fmt.Fprint(env.Stdout, result)
	if *zero {
		fmt.Fprint(env.Stdout, "\x00")
	} else if !*noNewline {
		fmt.Fprintln(env.Stdout)
	}

	return 0
}
