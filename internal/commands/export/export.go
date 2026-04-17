package export

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Export struct{}

func New() *Export {
	return &Export{}
}

func (e *Export) Name() string {
	return "export"
}

func (e *Export) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		// List exported variables in POSIX format
		keys := make([]string, 0, len(env.EnvVars))
		for k := range env.EnvVars {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Fprintf(env.Stdout, "export %s=\"%s\"\n", k, env.EnvVars[k])
		}
		return 0
	}

	for _, arg := range args {
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			env.EnvVars[parts[0]] = parts[1]
		} else {
			// In a real shell, this would mark an existing variable as exported.
			// Here, everything is "exported" if it's in the map.
		}
	}

	return 0
}
