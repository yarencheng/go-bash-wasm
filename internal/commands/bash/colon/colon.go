package colon

import (
	"context"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Colon struct{}

func New() *Colon {
	return &Colon{}
}

func (c *Colon) Name() string {
	return ":"
}

func (c *Colon) Run(ctx context.Context, env *commands.Environment, args []string) int {
	return 0
}
