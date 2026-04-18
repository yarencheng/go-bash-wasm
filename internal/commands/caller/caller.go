package caller

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Caller struct{}

func New() *Caller {
	return &Caller{}
}

func (c *Caller) Name() string {
	return "caller"
}

func (c *Caller) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// caller [expr]
	// Returns the context of the active subroutine call.
	// Since we don't have a call stack yet, we just return 1 (not in a subroutine).
	return 1
}
