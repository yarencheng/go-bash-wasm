package hostname

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Hostname struct{}

func New() *Hostname {
	return &Hostname{}
}

func (h *Hostname) Name() string {
	return "hostname"
}

func (h *Hostname) Run(ctx context.Context, env *commands.Environment, args []string) int {
	fmt.Fprintln(env.Stdout, "wasm-host")
	return 0
}
