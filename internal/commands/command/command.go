package command

import (
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c *Command) Name() string {
	return "command"
}

func (c *Command) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("command", pflag.ContinueOnError)
	useDefaultPath := flags.BoolP("path", "p", false, "use a default value for PATH")
	identify := flags.BoolP("verbose", "v", false, "print a description of COMMAND")
	verboseIdentify := flags.BoolP("Verbose", "V", false, "print a more verbose description of COMMAND")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "command: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return 0
	}

	name := remaining[0]

	if *identify || *verboseIdentify {
		cmd, ok := env.Registry.Get(name)
		if ok {
			if *verboseIdentify {
				fmt.Fprintf(env.Stdout, "%s is a builtin command\n", name)
			} else {
				fmt.Fprintln(env.Stdout, name)
			}
			return 0
		}
		// If not found in registry, search in PATH
		// (In this simulator, we might not have external binaries, so we just check registry)
		fmt.Fprintf(env.Stderr, "command: %s: not found\n", name)
		return 1
	}

	if *useDefaultPath {
		// Save old PATH
		oldPath := env.EnvVars["PATH"]
		env.EnvVars["PATH"] = "/usr/local/bin:/usr/bin:/bin"
		defer func() { env.EnvVars["PATH"] = oldPath }()
	}

	cmd, ok := env.Registry.Get(name)
	if !ok {
		fmt.Fprintf(env.Stderr, "command: %s: not found\n", name)
		return 127
	}

	return cmd.Run(ctx, env, remaining[1:])
}

