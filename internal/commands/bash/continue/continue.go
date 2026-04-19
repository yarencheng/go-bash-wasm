package continuecmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Continue struct{}

func New() *Continue {
	return &Continue{}
}

func (c *Continue) Name() string {
	return "continue"
}

func (c *Continue) Run(ctx context.Context, env *commands.Environment, args []string) int {
	n := 1
	if len(args) > 0 {
		val, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(env.Stderr, "continue: %s: numeric argument required\n", args[0])
			return 1
		}
		if val < 1 {
			fmt.Fprintf(env.Stderr, "continue: %d: loop count must be greater than or equal to 1\n", val)
			return 1
		}
		n = val
	}

	env.ContinueRequested = n
	return 0
}
