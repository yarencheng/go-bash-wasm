package shift

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Shift struct{}

func New() *Shift {
	return &Shift{}
}

func (s *Shift) Name() string {
	return "shift"
}

func (s *Shift) Run(ctx context.Context, env *commands.Environment, args []string) int {
	n := 1
	if len(args) > 0 {
		val, err := strconv.Atoi(args[0])
		if err != nil || val < 0 {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "shift: %s: invalid number\n", args[0])
			}
			return 1
		}
		n = val
	}

	if n > len(env.PositionalArgs) {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "shift: %d: shift count out of range\n", n)
		}
		return 1
	}

	env.PositionalArgs = env.PositionalArgs[n:]
	return 0
}
