package enable

import (
	"context"
	"fmt"
	"sort"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Enable struct{}

func New() *Enable {
	return &Enable{}
}

func (e *Enable) Name() string {
	return "enable"
}

func (e *Enable) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("enable", pflag.ContinueOnError)
	all := flags.BoolP("all", "a", false, "display all builtins")
	disable := flags.BoolP("disable", "n", false, "disable each NAME or display disabled builtins")
	print := flags.BoolP("print", "p", false, "print in reusable format")
	posix := flags.BoolP("special", "s", false, "print only POSIX special builtins")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "enable: %v\n", err)
		return 1
	}

	targets := flags.Args()

	if len(targets) == 0 || *print {
		// List builtins
		names := env.Registry.List()
		sort.Strings(names)

		for _, name := range names {
			isEnabled := env.Registry.IsEnabled(name)

			if *posix {
				// Special stubs - for now consider all special
			}

			if *all {
				// Print all
				e.printStatus(env, name, isEnabled, *print)
			} else if *disable {
				// Print only disabled
				if !isEnabled {
					e.printStatus(env, name, isEnabled, *print)
				}
			} else {
				// Print only enabled
				if isEnabled {
					e.printStatus(env, name, isEnabled, *print)
				}
			}
		}
		return 0
	}

	exitCode := 0
	for _, target := range targets {
		if *disable {
			if err := env.Registry.Disable(target); err != nil {
				fmt.Fprintf(env.Stderr, "enable: %s: not a shell builtin\n", target)
				exitCode = 1
			}
		} else {
			if err := env.Registry.Enable(target); err != nil {
				fmt.Fprintf(env.Stderr, "enable: %s: not a shell builtin\n", target)
				exitCode = 1
			}
		}
	}

	return exitCode
}

func (e *Enable) printStatus(env *commands.Environment, name string, isEnabled bool, reusable bool) {
	if reusable {
		if isEnabled {
			fmt.Fprintf(env.Stdout, "enable %s\n", name)
		} else {
			fmt.Fprintf(env.Stdout, "enable -n %s\n", name)
		}
	} else {
		if isEnabled {
			fmt.Fprintf(env.Stdout, "enable %s\n", name)
		} else {
			fmt.Fprintf(env.Stdout, "enable -n %s\n", name)
		}
	}
}
