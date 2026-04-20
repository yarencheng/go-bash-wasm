package getopts

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Getopts struct{}

func New() *Getopts {
	return &Getopts{}
}

func (g *Getopts) Name() string {
	return "getopts"
}

func (g *Getopts) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) < 2 {
		return 1
	}

	optstring := args[0]
	varname := args[1]

	// Arguments to parse are from args[2:]
	cmdArgs := args[2:]
	if len(cmdArgs) == 0 {
		cmdArgs = env.PositionalArgs
	}

	optindStr, ok := env.EnvVars["OPTIND"]
	if !ok {
		optindStr = "1"
	}
	optind, _ := strconv.Atoi(optindStr)

	if optind > len(cmdArgs) {
		env.EnvVars[varname] = "?"
		return 1
	}

	arg := cmdArgs[optind-1]
	if !strings.HasPrefix(arg, "-") || arg == "-" || arg == "--" {
		if arg == "--" {
			env.EnvVars["OPTIND"] = strconv.Itoa(optind + 1)
		}
		env.EnvVars[varname] = "?"
		return 1
	}

	option := arg[1:2]
	idx := strings.Index(optstring, option)
	if idx == -1 {
		env.EnvVars[varname] = "?"
		env.EnvVars["OPTARG"] = option
		env.EnvVars["OPTIND"] = strconv.Itoa(optind + 1)
		if !strings.HasPrefix(optstring, ":") {
			fmt.Fprintf(env.Stderr, "getopts: illegal option -- %s\n", option)
		}
		return 0
	}

	env.EnvVars[varname] = option
	if idx+1 < len(optstring) && optstring[idx+1] == ':' {
		// Option requires an argument
		if optind < len(cmdArgs) {
			env.EnvVars["OPTARG"] = cmdArgs[optind]
			env.EnvVars["OPTIND"] = strconv.Itoa(optind + 2)
		} else {
			env.EnvVars[varname] = "?"
			if strings.HasPrefix(optstring, ":") {
				env.EnvVars[varname] = ":"
				env.EnvVars["OPTARG"] = option
			} else {
				fmt.Fprintf(env.Stderr, "getopts: option requires an argument -- %s\n", option)
			}
			env.EnvVars["OPTIND"] = strconv.Itoa(optind + 1)
		}
	} else {
		env.EnvVars["OPTIND"] = strconv.Itoa(optind + 1)
	}

	return 0
}
