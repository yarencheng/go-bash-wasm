package export

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/pflag"
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
	flags := pflag.NewFlagSet("export", pflag.ContinueOnError)
	_ = flags.BoolP("functions", "f", false, "refer to shell functions (ignored)")
	_ = flags.BoolP("remove", "n", false, "remove the export property from variables (ignored)")
	_ = flags.BoolP("print", "p", false, "list exported variables (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "export: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: export [-fn] [name[=value] ...] or export -p\n")
		fmt.Fprintf(env.Stdout, "Set export attribute for shell variables.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "export")
		return 0
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
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

	for _, arg := range remaining {
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			env.EnvVars[parts[0]] = parts[1]
		}
	}

	return 0
}
