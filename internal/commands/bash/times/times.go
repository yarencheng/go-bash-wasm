package times

import (
	"context"
	"fmt"
	"time"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Times struct{}

func New() *Times {
	return &Times{}
}

func (t *Times) Name() string {
	return "times"
}

func (t *Times) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// Print shell's user and system times
	// For simulation, we'll use a portion of the time since startup as user time
	uptime := time.Since(env.StartTime)
	userTime := uptime / 10 // Mock user time
	sysTime := uptime / 20  // Mock system time

	// Format:
	// user_time system_time
	// user_time_children system_time_children
	fmt.Fprintf(env.Stdout, "%dm%d.%03ds %dm%d.%03ds\n",
		int(userTime.Minutes()), int(userTime.Seconds())%60, userTime.Milliseconds()%1000,
		int(sysTime.Minutes()), int(sysTime.Seconds())%60, sysTime.Milliseconds()%1000)

	// For children, we'll just show 0 for now as we don't track their resource usage
	fmt.Fprintf(env.Stdout, "0m0.000s 0m0.000s\n")

	return 0
}
