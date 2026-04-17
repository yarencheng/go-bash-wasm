package jobs

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Jobs struct{}

func New() *Jobs {
	return &Jobs{}
}

func (j *Jobs) Name() string {
	return "jobs"
}

func (j *Jobs) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("jobs", pflag.ContinueOnError)
	long := flags.BoolP("long", "l", false, "show process IDs in addition to the normal information")
	onlyPids := flags.BoolP("pids", "p", false, "list only the process ID of the job's process group leader")
	runningOnly := flags.BoolP("running", "r", false, "display only running jobs")
	stoppedOnly := flags.BoolP("stopped", "s", false, "display only stopped jobs")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "jobs: %v\n", err)
		}
		return 1
	}

	if len(env.Jobs) == 0 {
		return 0
	}

	// Simple display for now
	for _, job := range env.Jobs {
		if runningOnly != nil && *runningOnly && job.Status != "Running" {
			continue
		}
		if stoppedOnly != nil && *stoppedOnly && job.Status != "Stopped" {
			continue
		}

		if *onlyPids {
			fmt.Fprintln(env.Stdout, job.PID)
			continue
		}

		symbol := " "
		if job.ID == len(env.Jobs) {
			symbol = "+"
		} else if job.ID == len(env.Jobs)-1 {
			symbol = "-"
		}

		if *long {
			fmt.Fprintf(env.Stdout, "[%d]%s %d %-24s %s\n", job.ID, symbol, job.PID, job.Status, job.Command)
		} else {
			fmt.Fprintf(env.Stdout, "[%d]%s  %-24s %s\n", job.ID, symbol, job.Status, job.Command)
		}
	}

	return 0
}
