package runcon

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Runcon struct{}

func New() *Runcon {
	return &Runcon{}
}

func (r *Runcon) Name() string {
	return "runcon"
}

func (r *Runcon) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("runcon", pflag.ContinueOnError)
	_ = flags.BoolP("compute", "c", false, "compute process transition context before modifying")
	_ = flags.StringP("user", "u", "", "user identity")
	_ = flags.StringP("role", "r", "", "role")
	_ = flags.StringP("type", "t", "", "type")
	_ = flags.StringP("range", "l", "", "level/range")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "runcon: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		// If no command, just print context if context was given? 
		// Actually runcon requires a command.
		fmt.Fprintf(env.Stderr, "runcon: missing operand\n")
		return 1
	}

	// Stub: we don't support SELinux contexts.
	// We just execute the command directly.
	cmdName := remaining[0]
	cmdArgs := remaining[1:]

	if env.Registry == nil {
		fmt.Fprintf(env.Stderr, "runcon: registry not initialized\n")
		return 1
	}

	cmd, exists := env.Registry.Get(cmdName)
	if !exists {
		fmt.Fprintf(env.Stderr, "runcon: %s: command not found\n", cmdName)
		return 127
	}

	return cmd.Run(ctx, env, cmdArgs)
}
