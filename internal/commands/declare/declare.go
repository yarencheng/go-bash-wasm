package declare

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Declare struct{}

func New() *Declare {
	return &Declare{}
}

func (d *Declare) Name() string {
	return "declare"
}

func (d *Declare) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("declare", pflag.ContinueOnError)
	printFlag := flags.BoolP("print", "p", false, "display the attributes and value of each NAME")
	_ = flags.BoolP("export", "x", false, "make NAMEs export")
	_ = flags.BoolP("readonly", "r", false, "make NAMEs readonly")
	_ = flags.BoolP("integer", "i", false, "make NAMEs have the integer attribute")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "declare: %v\n", err)
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
			fmt.Fprintf(env.Stdout, "declare %s=\"%s\"\n", k, env.EnvVars[k])
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
