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

	totalDuration := time.Duration(0)
	for _, arg := range targets {
		unit := time.Second
		numPart := arg
		if len(arg) > 0 {
			last := arg[len(arg)-1]
			switch last {
			case 's':
				unit = time.Second
				numPart = arg[:len(arg)-1]
			case 'm':
				unit = time.Minute
				numPart = arg[:len(arg)-1]
			case 'h':
				unit = time.Hour
				numPart = arg[:len(arg)-1]
			case 'd':
				unit = time.Hour * 24
				numPart = arg[:len(arg)-1]
			}
		}

		val, err := strconv.ParseFloat(numPart, 64)
		if err != nil {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "sleep: invalid time interval '%s'\n", arg)
			}
			return 1
		}
		totalDuration += time.Duration(val * float64(unit))
	}

	timer := time.NewTimer(totalDuration)
	defer timer.Stop()

	select {
	case <-timer.C:
		return 0
	case <-ctx.Done():
		return 0
	}
}
