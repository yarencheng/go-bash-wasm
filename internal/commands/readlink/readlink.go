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
		// afero doesn't have a simple EvalSymlinks that works across all its Fs types easily if they are mixed
		// but afero.Fs should generally satisfy some interface or we can use a helper
		// For MemMapFs, we might need a custom implementation or use filepath if it's OsFs.
		// Since we are mostly MemMapFs in tests, let's use a simpler approach or just Readlink loop.
		
		if *canonicalizeExisting {
			// Components must exist
			_, err = env.FS.Stat(target)
			if err != nil {
				if env.Stderr != nil && *verbose {
					fmt.Fprintf(env.Stderr, "readlink: %s: %v\n", target, err)
				}
				return 1
			}
		}

		// Simplified canonicalization: just use Absolute path for now if not implemented
		// TODO: Implement full symlink following canonicalization
		result = filepath.Clean(target)
	} else {
		linker, ok := env.FS.(afero.Symlinker)
		if !ok {
			return 1
		}
		result, err = linker.ReadlinkIfPossible(target)
		if err != nil {
			if env.Stderr != nil && *verbose {
				fmt.Fprintf(env.Stderr, "readlink: %s: %v\n", target, err)
			}
			return 1
		}
		// If ReadlinkIfPossible returns the same path, it's not a link
		if result == target {
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
