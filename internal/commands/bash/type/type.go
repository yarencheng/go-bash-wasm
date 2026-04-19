package typecmd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Type struct{}

func New() *Type {
	return &Type{}
}

func (t *Type) Name() string {
	return "type"
}

func (t *Type) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("type", pflag.ContinueOnError)
	all := flags.BoolP("all", "a", false, "display all locations that contain an executable named NAME")
	pathOnly := flags.BoolP("path", "p", false, "returns the name of the disk file that would be executed")
	typeOnly := flags.BoolP("type", "t", false, "print a single word which is one of alias, keyword, function, builtin, or file")
	flags.BoolP("skip-functions", "f", false, "suppress shell function lookup")
	forcePath := flags.BoolP("force-path", "P", false, "force a PATH search for each NAME")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "type: %v\n", err)
		}
		return 2
	}

	names := flags.Args()
	if len(names) == 0 {
		return 0
	}

	exitCode := 0
	for _, name := range names {
		found := false

		// 1. Check Alias (unless -P)
		if !*forcePath {
			if alias, ok := env.Aliases[name]; ok {
				found = true
				if *typeOnly {
					fmt.Fprintln(env.Stdout, "alias")
				} else if !*pathOnly {
					fmt.Fprintf(env.Stdout, "%s is aliased to `%s'\n", name, alias)
				}
				if !*all {
					continue
				}
			}
		}

		// 2. Check Builtin (unless -P)
		if !*forcePath {
			if _, ok := env.Registry.Get(name); ok {
				found = true
				if *typeOnly {
					fmt.Fprintln(env.Stdout, "builtin")
				} else if !*pathOnly {
					fmt.Fprintf(env.Stdout, "%s is a shell builtin\n", name)
				}
				if !*all {
					continue
				}
			}
		}

		// 3. Check PATH
		pathVar := env.EnvVars["PATH"]
		paths := filepath.SplitList(pathVar)
		for _, p := range paths {
			fullPath := filepath.Join(p, name)
			if stat, err := env.FS.Stat(fullPath); err == nil && !stat.IsDir() {
				found = true
				if *typeOnly {
					fmt.Fprintln(env.Stdout, "file")
				} else if *pathOnly {
					fmt.Fprintln(env.Stdout, fullPath)
				} else {
					fmt.Fprintf(env.Stdout, "%s is %s\n", name, fullPath)
				}
				if !*all {
					break
				}
			}
		}

		if !found {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "type: %s: not found\n", name)
			}
			exitCode = 1
		}
	}

	return exitCode
}
