package eval

import (
	"context"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Eval struct{}

func New() *Eval {
	return &Eval{}
}

func (e *Eval) Name() string {
	return "eval"
}

func (e *Eval) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		return 0
	}

	// join arguments into a single string
	line := strings.Join(args, " ")
	
	// Very simple shell execution logic for now, similar to shell.go
	// In the future, this should probably use the full shell parser/executor
	subArgs := strings.Fields(line)
	if len(subArgs) == 0 {
		return 0
	}

	cmdName := subArgs[0]
	cmdArgs := subArgs[1:]

	if cmd, ok := env.Registry.Get(cmdName); ok {
		return cmd.Run(ctx, env, cmdArgs)
	}
	
	return 127
}
