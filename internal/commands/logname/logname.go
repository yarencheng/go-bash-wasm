package logname

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Logname struct{}

func New() *Logname {
	return &Logname{}
}

func (l *Logname) Name() string {
	return "logname"
}

func (l *Logname) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("logname", pflag.ContinueOnError)
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "logname: %v\n", err)
		return 1
	}

	fmt.Fprintln(env.Stdout, env.User)
	return 0
}
