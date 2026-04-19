package bg

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Bg struct{}

func New() *Bg {
	return &Bg{}
}

func (b *Bg) Name() string {
	return "bg"
}

func (b *Bg) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(env.Jobs) == 0 {
		fmt.Fprintf(env.Stderr, "bg: no current job\n")
		return 1
	}

	targetJob := env.Jobs[len(env.Jobs)-1]
	if len(args) > 0 {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(env.Stderr, "bg: %s: invalid job id\n", args[0])
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
			fmt.Fprintf(env.Stderr, "bg: %d: no such job\n", id)
			return 1
		}
	}

	targetJob.Status = "Running"
	fmt.Fprintf(env.Stdout, "[%d]+ %s &\n", targetJob.ID, targetJob.Command)
	return 0
}
