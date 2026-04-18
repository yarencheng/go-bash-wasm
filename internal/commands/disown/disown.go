package disown

import (
	"context"
	"fmt"
	"strconv"

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
	_ = flags.BoolP("remove-running", "r", false, "remove running jobs (ignored)")
	_ = flags.BoolP("hup", "h", false, "mark to not receive SIGHUP (ignored)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "disown: %v\n", err)
		return 1
	}

	if *all {
		env.Jobs = nil
		return 0
	}

	targets := flags.Args()
	if len(targets) == 0 {
		// Disown current job? For now just remove the last one if it exists.
		if len(env.Jobs) > 0 {
			env.Jobs = env.Jobs[:len(env.Jobs)-1]
		}
		return 0
	}

	for _, target := range targets {
		id, err := strconv.Atoi(target)
		if err != nil {
			fmt.Fprintf(env.Stderr, "disown: %s: invalid job id\n", target)
			continue
		}

		newJobs := make([]*commands.Job, 0, len(env.Jobs))
		found := false
		for _, job := range env.Jobs {
			if job.ID == id {
				found = true
				continue
			}
			newJobs = append(newJobs, job)
		}
		if !found {
			fmt.Fprintf(env.Stderr, "disown: %d: no such job\n", id)
			return 1
		}
		env.Jobs = newJobs
	}

	return 0
}
