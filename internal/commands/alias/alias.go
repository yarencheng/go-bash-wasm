package alias

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Alias struct{}

func New() *Alias {
	return &Alias{}
}

func (a *Alias) Name() string {
	return "alias"
}

func (a *Alias) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		keys := make([]string, 0, len(env.Aliases))
		for k := range env.Aliases {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Fprintf(env.Stdout, "alias %s='%s'\n", k, env.Aliases[k])
		}
		return 0
	}

	for _, arg := range args {
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			env.Aliases[parts[0]] = parts[1]
		} else {
			if val, ok := env.Aliases[arg]; ok {
				fmt.Fprintf(env.Stdout, "alias %s='%s'\n", arg, val)
			} else {
				fmt.Fprintf(env.Stderr, "alias: %s: not found\n", arg)
			}
		}
	}

	return 0
}
