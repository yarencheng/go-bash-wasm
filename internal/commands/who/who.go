package who

import (
	"context"
	"fmt"
	"time"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Who struct{}

func New() *Who {
	return &Who{}
}

func (w *Who) Name() string {
	return "who"
}

func (w *Who) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// Simulation for virtual environment
	fmt.Fprintf(env.Stdout, "%-10s %-10s %s\n", env.User, "tty1", env.StartTime.Format(time.UnixDate))
	return 0
}
