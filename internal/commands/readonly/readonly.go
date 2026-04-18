package readonly

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Readonly struct{}

func New() *Readonly {
	return &Readonly{}
}

func (r *Readonly) Name() string {
	return "readonly"
}

func (r *Readonly) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("readonly", pflag.ContinueOnError)
	printFlag := flags.BoolP("print", "p", false, "display the attributes and value of each NAME")
	_ = flags.BoolP("function", "f", false, "restrict action or display to function names and definitions")
	_ = flags.BoolP("array", "a", false, "restrict action or display to indexed array variables")
	_ = flags.BoolP("assoc", "A", false, "restrict action or display to associative array variables")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "readonly: %v\n", err)
		return 2
	}

	targets := flags.Args()

	if len(targets) == 0 || *printFlag {
		keys := make([]string, 0, len(env.EnvVars))
		for k := range env.EnvVars {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Fprintf(env.Stdout, "readonly %s=\"%s\"\n", k, env.EnvVars[k])
		}
		return 0
	}

	for _, arg := range targets {
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			env.EnvVars[parts[0]] = parts[1]
		}
	}

	return 0
}
