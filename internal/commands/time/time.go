package timecmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Time struct{}

func New() *Time {
	return &Time{}
}

func (t *Time) Name() string {
	return "time"
}

func (t *Time) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "\nreal 0m0.000s\nuser 0m0.000s\nsys 0m0.000s\n")
		}
		return 0
	}

	start := time.Now()
	
	// Re-execute the rest of the arguments as a command
	line := strings.Join(args, " ")
	exitCode := env.Executor.Execute(ctx, line)
	
	duration := time.Since(start)

	if env.Stderr != nil {
		fmt.Fprintf(env.Stderr, "\nreal %dm%.3fs\nuser 0m0.000s\nsys 0m0.000s\n", 
			int(duration.Minutes()), duration.Seconds()-float64(int(duration.Minutes())*60))
	}

	return exitCode
}
