package fg

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Fg struct{}

func New() *Fg {
	return &Fg{}
}

func (f *Fg) Name() string {
	return "fg"
}

func (f *Fg) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(env.Jobs) == 0 {
		fmt.Fprintf(env.Stderr, "fg: no current job\n")
		return 1
	}

	targetJob := env.Jobs[len(env.Jobs)-1]
	if len(args) > 0 {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(env.Stderr, "fg: %s: invalid job id\n", args[0])
			return 1
		}
		found := false
		for _, job := range env.Jobs {
			if job.ID == id {
				targetJob = job
				found = true
				break
			}
		}
		if !found {
			fmt.Fprintf(env.Stderr, "fg: %d: no such job\n", id)
			return 1
		}
	}

	targetJob.Status = "Running"
	fmt.Fprintf(env.Stdout, "%s\n", targetJob.Command)
	
	// In a real shell, we would wait for it.
	// In the simulator, if it's "running" in the foreground, we just return.
	return 0
}
