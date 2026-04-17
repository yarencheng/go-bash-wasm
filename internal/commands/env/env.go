package envcmd

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Env struct{}

func New() *Env {
	return &Env{}
}

func (e *Env) Name() string {
	return "env"
}

func (e *Env) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("env", pflag.ContinueOnError)
	flags.SetOutput(env.Stderr)

	ignoreEnv := flags.BoolP("ignore-environment", "i", false, "start with an empty environment")
	nullTerminated := flags.BoolP("null", "0", false, "end each output line with NUL, not newline")
	chdir := flags.StringP("chdir", "C", "", "change working directory to DIR")
	var unsetVars []string
	flags.StringArrayVarP(&unsetVars, "unset", "u", nil, "remove variable from the environment")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "env: %v\n", err)
		return 1
	}

	remainingArgs := flags.Args()

	newEnvVars := make(map[string]string)
	if !*ignoreEnv {
		for k, v := range env.EnvVars {
			newEnvVars[k] = v
		}
	}

	for _, name := range unsetVars {
		delete(newEnvVars, name)
	}

	// Parse NAME=VALUE assignments
	idx := 0
	for ; idx < len(remainingArgs); idx++ {
		arg := remainingArgs[idx]
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			newEnvVars[parts[0]] = parts[1]
		} else {
			break
		}
	}

	cmdArgs := remainingArgs[idx:]

	if len(cmdArgs) == 0 {
		keys := make([]string, 0, len(newEnvVars))
		for k := range newEnvVars {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		terminator := "\n"
		if *nullTerminated {
			terminator = "\x00"
		}

		for _, k := range keys {
			fmt.Fprintf(env.Stdout, "%s=%s%s", k, newEnvVars[k], terminator)
		}
		return 0
	}

	// Run command
	cmdName := cmdArgs[0]
	if cmd, ok := env.Registry.Get(cmdName); ok {
		// Create a sub-environment
		subEnv := *env
		subEnv.EnvVars = newEnvVars
		if *chdir != "" {
			subEnv.Cwd = *chdir
		}
		return cmd.Run(ctx, &subEnv, cmdArgs[1:])
	}

	fmt.Fprintf(env.Stderr, "env: ‘%s’: No such file or directory\n", cmdName)
	return 127
}
