package unalias

import (
	"context"
	"fmt"

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
	if len(args) == 0 {
		fmt.Fprintf(env.Stderr, "unalias: usage: unalias name [name ...]\n")
		return 2
	}

	exitCode := 0
	for _, arg := range args {
		if _, ok := env.Aliases[arg]; ok {
			delete(env.Aliases, arg)
		} else {
			fmt.Fprintf(env.Stderr, "unalias: %s: not found\n", arg)
			exitCode = 1
		}
	}

	return exitCode
}
