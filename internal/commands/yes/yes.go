package yes

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Yes struct{}

func New() *Yes {
	return &Yes{}
}

func (y *Yes) Name() string {
	return "yes"
}

func (y *Yes) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("yes", pflag.ContinueOnError)
	// pflag handles --help by default if we don't define it, but we can define it to be explicit.
	
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "yes: %v\n", err)
		return 1
	}

	text := "y"
	remaining := flags.Args()
	if len(remaining) > 0 {
		text = strings.Join(remaining, " ")
	}

	for {
		select {
		case <-ctx.Done():
			return 0
		default:
			// Optimization: writes can be slow, but for a simulator this is fine.
			_, err := fmt.Fprintln(env.Stdout, text)
			if err != nil {
				return 0 // Broken pipe or similar
			}
		}
	}
}
