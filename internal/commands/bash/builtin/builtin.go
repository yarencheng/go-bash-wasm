package builtin

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Builtin struct{}

func New() *Builtin {
	return &Builtin{}
}

func (b *Builtin) Name() string {
	return "builtin"
}

func (b *Builtin) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		return 0
	}

	name := args[0]
	cmd, ok := env.Registry.Get(name)
	if !ok {
		fmt.Fprintf(env.Stderr, "builtin: %s: not a shell builtin\n", name)
		return 1
	}

	return cmd.Run(ctx, env, args[1:])
}
