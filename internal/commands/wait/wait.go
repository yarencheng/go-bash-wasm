package wait

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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
	pidVar := flags.StringP("pid-var", "p", "", "store the process ID or job management ID of the job for which the exit status is returned in the variable VARNAME")

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
		var jobID int
		var pid int
		isJobSpec := strings.HasPrefix(target, "%")

		if isJobSpec {
			id, err := strconv.Atoi(target[1:])
			if err == nil {
				jobID = id
			}
		} else {
			p, err := strconv.Atoi(target)
			if err == nil {
				pid = p
			}
		}

		found := false
		for _, job := range env.Jobs {
			if (isJobSpec && job.ID == jobID) || (!isJobSpec && job.PID == pid) {
				found = true
				if *pidVar != "" {
					env.EnvVars[*pidVar] = strconv.Itoa(job.PID)
				}
				if job.Status == "Running" {
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
