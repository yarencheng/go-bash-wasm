package hostid

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Hostid struct{}

func New() *Hostid {
	return &Hostid{}
}

func (h *Hostid) Name() string {
	return "hostid"
}

func (h *Hostid) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// hostid usually returns a 32-bit identifier.
	// We'll return a fixed value for the simulator.
	fmt.Fprintln(env.Stdout, "007f0101")
	return 0
}
