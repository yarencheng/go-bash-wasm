package logname

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Logname struct{}

func New() *Logname {
	return &Logname{}
}

func (l *Logname) Name() string {
	return "logname"
}

func (l *Logname) Run(ctx context.Context, env *commands.Environment, args []string) int {
	fmt.Fprintln(env.Stdout, env.User)
	return 0
}
