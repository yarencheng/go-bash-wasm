package suspend

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Suspend struct{}

func New() *Suspend {
	return &Suspend{}
}

func (s *Suspend) Name() string {
	return "suspend"
}

func (s *Suspend) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("suspend", pflag.ContinueOnError)
	force := flags.BoolP("force", "f", false, "force suspend even if it is a login shell")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "suspend: %v\n", err)
		return 1
	}

	// Stub: we can't really suspend the WASM process this way.
	fmt.Fprintf(env.Stderr, "suspend: not supported in this environment\n")
	if *force {
		return 0
	}
	return 1
}
