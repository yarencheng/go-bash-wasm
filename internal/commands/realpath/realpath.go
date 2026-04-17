package realpath

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Realpath struct{}

func New() *Realpath {
	return &Realpath{}
}

func (r *Realpath) Name() string {
	return "realpath"
}

func (r *Realpath) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("realpath", pflag.ContinueOnError)
	canonicalizeExisting := flags.BoolP("canonicalize-existing", "e", false, "all components of the path must exist")
	_ = flags.BoolP("canonicalize-missing", "m", false, "no components of the path need exist")
	quiet := flags.BoolP("quiet", "q", false, "suppress most error messages")
	strip := flags.BoolP("strip", "s", false, "no components of the path need exist")
	zero := flags.BoolP("zero", "z", false, "end each output line with NUL, not newline")
	
	// -L and -P are usually flags for following symlinks
	_ = flags.BoolP("logical", "L", false, "resolve '..' components before symlinks")
	_ = flags.BoolP("physical", "P", false, "resolve symlinks (default)")

	if err := flags.Parse(args); err != nil {
		if !*quiet {
			fmt.Fprintf(env.Stderr, "realpath: %v\n", err)
		}
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		if !*quiet {
			fmt.Fprintf(env.Stderr, "realpath: missing operand\n")
		}
		return 1
	}

	exitCode := 0
	for _, target := range remaining {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		if *canonicalizeExisting {
			_, err := env.FS.Stat(fullPath)
			if err != nil {
				if !*quiet {
					fmt.Fprintf(env.Stderr, "realpath: %s: %v\n", target, err)
				}
				exitCode = 1
				continue
			}
		}

		// Simple canonicalization for MemMapFs
		cleanPath := filepath.Clean(fullPath)
		
		sep := "\n"
		if *zero {
			sep = "\x00"
		}

		if !*strip {
			fmt.Fprint(env.Stdout, cleanPath+sep)
		} else {
			fmt.Fprint(env.Stdout, cleanPath+sep)
		}
	}

	return exitCode
}
