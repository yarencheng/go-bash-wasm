package disown

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Disown struct{}

func New() *Disown {
	return &Disown{}
}

func (d *Disown) Name() string {
	return "disown"
}

func (d *Disown) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("disown", pflag.ContinueOnError)
	all := flags.BoolP("all", "a", false, "remove all jobs")
	runningOnly := flags.BoolP("remove-running", "r", false, "remove running jobs")
	hup := flags.BoolP("hup", "h", false, "mark to not receive SIGHUP")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "disown: %v\n", err)
		return 1
	}

	if *all {
		if *runningOnly {
			newJobs := make([]*commands.Job, 0, len(env.Jobs))
			for _, job := range env.Jobs {
				if job.Status != "Running" {
					newJobs = append(newJobs, job)
				}
			}
			env.Jobs = newJobs
		} else {
			env.Jobs = nil
		}
		return 0
	}

	targets := flags.Args()
	if len(targets) == 0 {
		// Disown current job? For now just remove the last one if it exists.
		if len(env.Jobs) > 0 {
			targetIdx := len(env.Jobs) - 1
			if *runningOnly && env.Jobs[targetIdx].Status != "Running" {
				return 0
			}
			if *hup {
				// We don't have a field for this, but we simulate it by doing nothing 
				// since our jobs are already persistent in the simulation.
				return 0
			}
			env.Jobs = env.Jobs[:targetIdx]
		}
		return 0
	}

	for _, target := range targets {
		var jobID int
		isJobSpec := strings.HasPrefix(target, "%")
		if isJobSpec {
			id, err := strconv.Atoi(target[1:])
			if err == nil {
				jobID = id
			}
		} else {
			id, err := strconv.Atoi(target)
			if err == nil {
				jobID = id
			}
		}

		newJobs := make([]*commands.Job, 0, len(env.Jobs))
		found := false
		for _, job := range env.Jobs {
			if job.ID == jobID {
				found = true
				if *runningOnly && job.Status != "Running" {
					newJobs = append(newJobs, job)
					continue
				}
				if *hup {
					newJobs = append(newJobs, job)
					continue
				}
				continue
			}
			newJobs = append(newJobs, job)
		}
		if !found {
			fmt.Fprintf(env.Stderr, "disown: %d: no such job\n", jobID)
			return 1
		}
		env.Jobs = newJobs
	}

	return 0
}
