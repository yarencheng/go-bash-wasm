package sleep

import (
	"context"
	"fmt"
	"strconv"
	"time"

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
	if len(args) == 0 {
		fmt.Fprintf(env.Stderr, "sleep: missing operand\n")
		return 1
	}

	seconds, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Fprintf(env.Stderr, "sleep: invalid time interval '%s'\n", args[0])
		return 1
	}

	duration := time.Duration(seconds * float64(time.Second))
	
	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-timer.C:
		return 0
	case <-ctx.Done():
		return 0
	}
}
