package unset

import (
	"context"

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
	for _, arg := range args {
		delete(env.EnvVars, arg)
		delete(env.Aliases, arg)
	}
	return 0
}
