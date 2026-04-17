package wait

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Wait struct{}

func New() *Wait {
	return &Wait{}
}

func (w *Wait) Name() string {
	return "wait"
}

func (w *Wait) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("wait", pflag.ContinueOnError)
	_ = flags.BoolP("wait-all", "f", false, "wait for each ID to terminate before returning (ignored)")
	_ = flags.BoolP("next", "n", false, "wait for the next job to terminate and return its exit status (ignored)")
	
	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "wait: %v\n", err)
		}
		return 1
	}

	targets := flags.Args()
	
	// Since we don't have real background processes in this simulator yet,
	// wait will mostly return 0 if the job is not found or already done.
	
	if len(targets) == 0 {
		// Wait for all jobs? Our jobs are just a list in env.Jobs
		// For now, if no targets, just return 0.
		return 0
	}

	lastStatus := 0
	for _, target := range targets {
		pid, err := strconv.Atoi(target)
		if err != nil {
			// Try job spec %n?
			// For now, just integer PIDs
		}
		
		found := false
		for _, job := range env.Jobs {
			if job.PID == pid {
				found = true
				if job.Status == "Running" {
					// We can't really wait in WASM easily without a scheduler,
					// so we assume it completes immediately for this simulation.
					job.Status = "Done"
				}
				break
			}
		}
		
		if !found {
			// Bash returns 127 if PID not found and not a job
			lastStatus = 127
		}
	}

	return lastStatus
}
