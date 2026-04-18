package set

import (
	"context"
	"fmt"
	"sort"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Set struct{}

func New() *Set {
	return &Set{}
}

func (s *Set) Name() string {
	return "set"
}

func (s *Set) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// If no arguments, print all variables
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

	flags := pflag.NewFlagSet("set", pflag.ContinueOnError)
	_ = flags.BoolP("errexit", "e", false, "exit immediately if a command exits with a non-zero status")
	_ = flags.BoolP("xtrace", "x", false, "print commands and their arguments as they are executed")
	_ = flags.BoolP("nounset", "u", false, "treat unset variables as an error when substituting")
	_ = flags.BoolP("noglob", "f", false, "disable pathname expansion")
	_ = flags.BoolP("noclobber", "C", false, "do not overwrite existing files with >")
	_ = flags.BoolP("option-name", "o", false, "change the behavior of the shell")
	_ = flags.BoolP("allexport", "a", false, "mark variables which are modified or created for export")
	_ = flags.BoolP("notify", "b", false, "report the status of terminated background jobs immediately")
	_ = flags.BoolP("hashall", "h", false, "remember the location of commands as they are looked up")
	_ = flags.BoolP("keyword", "k", false, "all arguments in the form of assignment statements are placed in the environment")
	_ = flags.BoolP("monitor", "m", false, "job control is enabled")
	_ = flags.BoolP("noexec", "n", false, "read commands but do not execute them")
	_ = flags.BoolP("privileged", "p", false, "privileged mode")
	_ = flags.BoolP("onecmd", "t", false, "exit after reading and executing one command")
	_ = flags.BoolP("verbose", "v", false, "print shell input lines as they are read")
	_ = flags.BoolP("braceexpand", "B", false, "the shell will perform brace expansion")
	_ = flags.BoolP("errtrace", "E", false, "if set, any trap on ERR is inherited by shell functions")
	_ = flags.BoolP("histexpand", "H", false, "enable ! style history substitution")
	_ = flags.BoolP("physical", "P", false, "do not follow symbolic links when executing commands such as cd")
	_ = flags.BoolP("functrace", "T", false, "if set, any trap on DEBUG and RETURN are inherited by shell functions")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "set: %v\n", err)
		return 2
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: set [-abefhkmnptuvBCGHPT] [-o option-name] [arg ...]\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "set")
		return 0
	}

	remaining := flags.Args()
	if len(remaining) > 0 {
		// Bash set command also allows setting positional parameters
		env.PositionalArgs = remaining
	}

	return 0
}
