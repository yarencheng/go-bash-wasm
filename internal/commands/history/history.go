package history

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type History struct{}

func New() *History {
	return &History{}
}

func (h *History) Name() string {
	return "history"
}

func (h *History) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("history", pflag.ContinueOnError)
	clear := flags.BoolP("clear", "c", false, "clear the history list")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "history: %v\n", err)
		return 1
	}

	if *clear {
		env.History = nil
		return 0
	}

	targets := flags.Args()
	limit := len(env.History)
	if len(targets) > 0 {
		n, err := strconv.Atoi(targets[0])
		if err == nil && n < limit {
			limit = n
		}
	}

	start := len(env.History) - limit
	if start < 0 {
		start = 0
	}

	for i := start; i < len(env.History); i++ {
		fmt.Fprintf(env.Stdout, "%5d  %s\n", i+1, env.History[i])
	}

	return 0
}
