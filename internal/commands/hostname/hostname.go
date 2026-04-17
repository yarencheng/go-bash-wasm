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
	if len(args) > 0 {
		env.EnvVars["HOSTNAME"] = args[0]
		return 0
	}

	hostname, ok := env.EnvVars["HOSTNAME"]
	if !ok {
		hostname = "wasm-host"
	}
	fmt.Fprintln(env.Stdout, hostname)
	return 0
}
