package boolcmd

import (
	"context"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type True struct{}

func NewTrue() *True {
	return &True{}
}

func (t *True) Name() string {
	return "true"
}

func (t *True) Run(ctx context.Context, env *commands.Environment, args []string) int {
	return 0
}

type False struct{}

func NewFalse() *False {
	return &False{}
}

func (f *False) Name() string {
	return "false"
}

func (f *False) Run(ctx context.Context, env *commands.Environment, args []string) int {
	return 1
}
