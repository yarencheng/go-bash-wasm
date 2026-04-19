package nproccmd

import (
	"context"
	"fmt"
	"runtime"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Nproc struct{}

func New() *Nproc {
	return &Nproc{}
}

func (n *Nproc) Name() string {
	return "nproc"
}

func (n *Nproc) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("nproc", pflag.ContinueOnError)
	all := flags.Bool("all", false, "print the number of installed processors")
	ignore := flags.Int("ignore", 0, "if possible, exclude N processors")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "nproc: %v\n", err)
		return 1
	}

	num := runtime.NumCPU()
	if *all {
		// In Go/WASM they might be same
	}

	result := num - *ignore
	if result < 1 {
		result = 1
	}

	fmt.Fprintln(env.Stdout, result)
	return 0
}
