package printenv

import (
	"context"
	"fmt"
	"sort"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Printenv struct{}

func New() *Printenv {
	return &Printenv{}
}

func (p *Printenv) Name() string {
	return "printenv"
}

func (p *Printenv) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
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

	exitCode := 0
	for _, arg := range args {
		if val, ok := env.EnvVars[arg]; ok {
			fmt.Fprintln(env.Stdout, val)
		} else {
			exitCode = 1
		}
	}

	return exitCode
}
