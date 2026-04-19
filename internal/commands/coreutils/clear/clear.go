package clear

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Clear struct{}

func New() *Clear {
	return &Clear{}
}

func (c *Clear) Name() string {
	return "clear"
}

func (c *Clear) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// ANSI escape code to clear screen and move cursor to home
	fmt.Fprint(env.Stdout, "\033[H\033[2J")
	return 0
}
