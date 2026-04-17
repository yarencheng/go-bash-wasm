package exit

import (
	"context"
	"strconv"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Exit struct{}

func New() *Exit {
	return &Exit{}
}

func (e *Exit) Name() string {
	return "exit"
}

func (e *Exit) Run(ctx context.Context, env *commands.Environment, args []string) int {
	code := 0
	if len(args) > 0 {
		c, err := strconv.Atoi(args[0])
		if err == nil {
			code = c
		}
	}

	env.ExitRequested = true
	env.ExitCode = code
	return code
}
