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
	lower := flags.BoolP("lowercase", "l", false, "convert NAMEs to lowercase on assignment")
	upper := flags.BoolP("uppercase", "u", false, "convert NAMEs to uppercase on assignment")
	_ = flags.BoolP("nameref", "n", false, "make NAME a reference to the variable named by its value")
	_ = flags.BoolP("trace", "t", false, "make NAMEs have the trace attribute")
	_ = flags.BoolP("function", "f", false, "restrict action or display to function names and definitions")
	_ = flags.BoolP("funcname", "F", false, "restrict display to function names only")
	_ = flags.BoolP("global", "g", false, "create NAMEs in the global scope")
	_ = flags.BoolP("inherit", "I", false, "inherit attributes from name in surrounding scope")
	_ = flags.BoolP("array", "a", false, "make NAMEs indexed arrays")
	_ = flags.BoolP("assoc", "A", false, "make NAMEs associative arrays")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "declare: %v\n", err)
		return 2
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: declare [-aAfFgilnrtux] [-p] [name[=value] ...]\n")
		fmt.Fprintf(env.Stdout, "Declare variables and give them attributes.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "declare")
		return 0
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
			val := parts[1]
			if *lower {
				val = strings.ToLower(val)
			}
			if *upper {
				val = strings.ToUpper(val)
			}
			env.EnvVars[parts[0]] = val
		}
	}

	return 0
}
