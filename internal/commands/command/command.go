package command

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c *Command) Name() string {
	return "command"
}

func (c *Command) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		return 0
	}

	// For now we just bypass aliases (if we had functions we'd bypass them too)
	name := args[0]
	cmd, ok := env.Registry.Get(name)
	if !ok {
		fmt.Fprintf(env.Stderr, "command: %s: not found\n", name)
		return 127
	}

	return cmd.Run(ctx, env, args[1:])
}
