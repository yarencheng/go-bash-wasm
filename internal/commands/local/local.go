package local

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Local struct{}

func New() *Local {
	return &Local{}
}

func (l *Local) Name() string {
	return "local"
}

func (l *Local) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("local", pflag.ContinueOnError)

	// local accepts same flags as declare
	_ = flags.BoolP("integer", "i", false, "make NAMEs have the integer attribute")
	_ = flags.BoolP("readonly", "r", false, "make NAMEs readonly")
	_ = flags.BoolP("export", "x", false, "make NAMEs export")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "local: %v\n", err)
		return 2
	}

	targets := flags.Args()
	if len(targets) == 0 {
		// In a real shell, would list local vars. For now, just success.
		return 0
	}

	for _, arg := range targets {
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			env.EnvVars[parts[0]] = parts[1]
		} else {
			// Declaration without value
			if _, exists := env.EnvVars[arg]; !exists {
				env.EnvVars[arg] = ""
			}
		}
	}

	return 0
}
