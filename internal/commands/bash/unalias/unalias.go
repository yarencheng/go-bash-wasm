package unalias

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Unalias struct{}

func New() *Unalias {
	return &Unalias{}
}

func (u *Unalias) Name() string {
	return "unalias"
}

func (u *Unalias) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("unalias", pflag.ContinueOnError)
	all := flags.BoolP("all", "a", false, "remove all alias definitions")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "unalias: %v\n", err)
		}
		return 2
	}

	if *all {
		env.Aliases = make(map[string]string)
		return 0
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "unalias: usage: unalias [-a] name [name ...]\n")
		}
		return 2
	}

	exitCode := 0
	for _, arg := range remaining {
		if _, ok := env.Aliases[arg]; ok {
			delete(env.Aliases, arg)
		} else {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "unalias: %s: not found\n", arg)
			}
			exitCode = 1
		}
	}

	return exitCode
}
