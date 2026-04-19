package exec

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Exec struct{}

func New() *Exec {
	return &Exec{}
}

func (e *Exec) Name() string {
	return "exec"
}

func (e *Exec) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("exec", pflag.ContinueOnError)
	login := flags.BoolP("login", "l", false, "login shell")
	_ = flags.StringP("name", "a", "", "use name as the zeroth argument of the executed command")
	_ = flags.BoolP("clean", "c", false, "execute command with an empty environment")

	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "exec: %v\n", err)
		}
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: exec [-cl] [-a name] [command [arguments]]\n")
		fmt.Fprintf(env.Stdout, "Replace the shell with the given command.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "exec")
		return 0
	}

	targets := flags.Args()
	if len(targets) == 0 {
		// Just setting login shell or something? Not really supported.
		return 0
	}

	cmdName := targets[0]
	cmdArgs := targets[1:]

	env.ExitRequested = true
	_ = login // TODO: use this

	if cmd, ok := env.Registry.Get(cmdName); ok {
		return cmd.Run(ctx, env, cmdArgs)
	}

	if env.Stderr != nil {
		fmt.Fprintf(env.Stderr, "exec: %s: command not found\n", cmdName)
	}
	return 127
}
