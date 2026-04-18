package returncmd

import (
	"context"
	"strconv"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Return struct{}

func New() *Return {
	return &Return{}
}

func (r *Return) Name() string {
	return "return"
}

func (r *Return) Run(ctx context.Context, env *commands.Environment, args []string) int {
	code := env.ExitCode
	if len(args) > 0 {
		c, err := strconv.Atoi(args[0])
		if err == nil {
			code = c
		}
	}

	env.ReturnRequested = true
	env.ReturnCode = code
	return code
}
