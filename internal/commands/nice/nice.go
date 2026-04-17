package nice

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Nice struct{}

func New() *Nice {
	return &Nice{}
}

func (n *Nice) Name() string {
	return "nice"
}

func (n *Nice) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("nice", pflag.ContinueOnError)
	adjustment := flags.IntP("adjustment", "n", 10, "add adjustment to the priority")
	
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "nice: %v\n", err)
		return 125
	}

	targets := flags.Args()
	if len(targets) == 0 {
		// print current adjustment? not supported in sim
		fmt.Fprintln(env.Stdout, "0")
		return 0
	}

	cmdName := targets[0]
	cmdArgs := targets[1:]

	cmd, ok := env.Registry.Get(cmdName)
	if !ok {
		fmt.Fprintf(env.Stderr, "nice: %s: No such file or directory\n", cmdName)
		return 127
	}

	// We ignore adjustment for now
	_ = adjustment
	
	return cmd.Run(ctx, env, cmdArgs)
}
