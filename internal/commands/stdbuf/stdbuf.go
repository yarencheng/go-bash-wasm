package stdbuf

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Stdbuf struct{}

func New() *Stdbuf {
	return &Stdbuf{}
}

func (s *Stdbuf) Name() string {
	return "stdbuf"
}

func (s *Stdbuf) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("stdbuf", pflag.ContinueOnError)
	_ = flags.StringP("input", "i", "", "adjust standard input stream buffering")
	_ = flags.StringP("output", "o", "", "adjust standard output stream buffering")
	_ = flags.StringP("error", "e", "", "adjust standard error stream buffering")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "stdbuf: %v\n", err)
		return 125
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		fmt.Fprintf(env.Stderr, "stdbuf: missing operand\n")
		return 125
	}

	// Stub: we just execute the command without actually changing buffering.
	// In a real system this would use LD_PRELOAD.
	
	cmdName := remaining[0]
	cmdArgs := remaining[1:]

	// Find the command in the registry
	if env.Registry == nil {
		fmt.Fprintf(env.Stderr, "stdbuf: registry not initialized\n")
		return 1
	}

	cmd, exists := env.Registry.Get(cmdName)
	if !exists {
		fmt.Fprintf(env.Stderr, "stdbuf: %s: command not found\n", cmdName)
		return 127
	}

	return cmd.Run(ctx, env, cmdArgs)
}
