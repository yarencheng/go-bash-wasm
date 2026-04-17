package sleep

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Sleep struct{}

func New() *Sleep {
	return &Sleep{}
}

func (s *Sleep) Name() string {
	return "sleep"
}

func (s *Sleep) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("sleep", pflag.ContinueOnError)
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "sleep: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "sleep: missing operand\n")
		return 1
	}

	totalSeconds := 0.0
	for _, arg := range targets {
		seconds, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(env.Stderr, "sleep: invalid time interval '%s'\n", arg)
			return 1
		}
		totalSeconds += seconds
	}

	duration := time.Duration(totalSeconds * float64(time.Second))
	
	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-timer.C:
		return 0
	case <-ctx.Done():
		return 0
	}
}
