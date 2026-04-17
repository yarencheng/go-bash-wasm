package envcmd

import (
	"context"
	"fmt"
	"sort"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Env struct{}

func New() *Env {
	return &Env{}
}

func (e *Env) Name() string {
	return "env"
}

func (e *Env) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// Simple implementation: list variables
	// In a real shell, 'env' can also run commands with modified env.
	
	keys := make([]string, 0, len(env.EnvVars))
	for k := range env.EnvVars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Fprintf(env.Stdout, "%s=%s\n", k, env.EnvVars[k])
	}
	
	return 0
}
