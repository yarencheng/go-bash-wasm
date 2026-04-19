package unset

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Unset struct{}

func New() *Unset {
	return &Unset{}
}

func (u *Unset) Name() string {
	return "unset"
}

func (u *Unset) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("unset", pflag.ContinueOnError)
	varsOnly := flags.BoolP("vars", "v", false, "unset variables")
	funcsOnly := flags.BoolP("funcs", "f", false, "unset functions")
	_ = flags.BoolP("nameref", "n", false, "treat each name as a nameref (ignored)")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "unset: %v\n", err)
		}
		return 1
	}

	targets := flags.Args()
	for _, arg := range targets {
		if *funcsOnly {
			if env.Functions != nil {
				delete(env.Functions, arg)
			}
		} else if *varsOnly {
			delete(env.EnvVars, arg)
			if env.Arrays != nil {
				delete(env.Arrays, arg)
			}
		} else {
			// Default: try variable first, then function
			if _, ok := env.EnvVars[arg]; ok {
				delete(env.EnvVars, arg)
			} else if env.Arrays != nil {
				if _, ok := env.Arrays[arg]; ok {
					delete(env.Arrays, arg)
				}
			} else if env.Functions != nil {
				delete(env.Functions, arg)
			}
		}
	}
	return 0
}
