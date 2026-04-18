package dircolors

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Dircolors struct{}

func New() *Dircolors {
	return &Dircolors{}
}

func (d *Dircolors) Name() string {
	return "dircolors"
}

func (d *Dircolors) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("dircolors", pflag.ContinueOnError)
	_ = flags.BoolP("sh", "b", false, "output Bourne shell code")
	_ = flags.BoolP("csh", "c", false, "output C shell code")
	_ = flags.BoolP("print-database", "p", false, "print default database")
	_ = flags.Bool("print-ls-colors", false, "print LS_COLORS")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "dircolors: %v\n", err)
		return 1
	}

	// Stub: we don't really support color configuration via dircolors yet.
	// We just print a default LS_COLORS if requested, or empty code.

	return 0
}
