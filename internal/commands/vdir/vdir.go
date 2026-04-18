package vdir

import (
	"context"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ls"
)

type Vdir struct {
	ls *ls.Ls
}

func New() *Vdir {
	return &Vdir{ls: ls.New()}
}

func (v *Vdir) Name() string {
	return "vdir"
}

func (v *Vdir) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// 'vdir' is equivalent to 'ls -l -b' (long format, escaped)
	return v.ls.Run(ctx, env, append([]string{"-l"}, args...))
}
