package nohup

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Nohup struct{}

func New() *Nohup {
	return &Nohup{}
}

func (n *Nohup) Name() string {
	return "nohup"
}

func (n *Nohup) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("nohup", pflag.ContinueOnError)
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "nohup: %v\n", err)
		return 125
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "nohup: missing operand\n")
		return 125
	}

	cmdName := targets[0]
	cmdArgs := targets[1:]

	cmd, ok := env.Registry.Get(cmdName)
	if !ok {
		fmt.Fprintf(env.Stderr, "nohup: failed to run command '%s': No such file or directory\n", cmdName)
		return 127
	}

	// In real nohup, we would redirect stdout/stderr to nohup.out if they are terminals.
	// For simulator, we just run the command.
	// "nohup: ignoring input and appending output to 'nohup.out'" message skipped for now.
	
	return cmd.Run(ctx, env, cmdArgs)
}
