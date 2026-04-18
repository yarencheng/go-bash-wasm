package set

import (
	"context"
	"fmt"
	"sort"
	"strings"

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
	_ = flags.BoolP("posix", "o", false, "change the behavior of the shell to follow the POSIX standard")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "set: %v\n", err)
		return 2
	}

	remaining := flags.Args()
	if len(remaining) > 0 {
		// Bash set command also allows setting positional parameters
		// env.PositionalArgs = remaining
	}

	return 0
}
