package tty

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type TTY struct{}

func New() *TTY {
	return &TTY{}
}

func (t *TTY) Name() string {
	return "tty"
}

func (t *TTY) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("tty", pflag.ContinueOnError)
	silent := flags.BoolP("silent", "s", false, "print nothing, only return an exit status")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "tty: %v\n", err)
		return 3 // GNU tty uses 3 for error
	}

	// In a real system, we'd check if stdin is a terminal.
	// For simulator, we always return /dev/tty for now.

	if !*silent {
		fmt.Fprintln(env.Stdout, "/dev/tty")
	}

	return 0
}
