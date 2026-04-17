package userscmd

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Users struct{}

func New() *Users {
	return &Users{}
}

func (u *Users) Name() string {
	return "users"
}

func (u *Users) Run(ctx context.Context, env *commands.Environment, args []string) int {
	fmt.Fprintln(env.Stdout, env.User)
	return 0
}
