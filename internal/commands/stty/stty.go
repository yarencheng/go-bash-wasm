package stty

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Stty struct{}

func New() *Stty {
	return &Stty{}
}

func (s *Stty) Name() string {
	return "stty"
}

func (s *Stty) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("stty", pflag.ContinueOnError)
	_ = flags.StringP("file", "F", "", "open and use the specified DEVICE instead of stdin")
	_ = flags.BoolP("all", "a", false, "print all current settings in human-readable form")
	_ = flags.BoolP("save", "g", false, "print all current settings in a stty-readable form")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "stty: %v\n", err)
		return 1
	}

	// Stub: in WASM we don't really have a TTY that we can configure this way.

	if len(args) == 0 || (len(args) == 1 && (args[0] == "-a" || args[0] == "--all")) {
		fmt.Fprintf(env.Stdout, "speed 38400 baud; line = 0;\n")
		fmt.Fprintf(env.Stdout, "-brkint -imaxbel\n")
	}

	return 0
}
