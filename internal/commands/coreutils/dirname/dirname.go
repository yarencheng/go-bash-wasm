package dirname

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Dirname struct{}

func New() *Dirname {
	return &Dirname{}
}

func (d *Dirname) Name() string {
	return "dirname"
}

func (d *Dirname) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("dirname", pflag.ContinueOnError)
	zero := flags.BoolP("zero", "z", false, "end each output line with NUL, not newline")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "dirname: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "dirname: missing operand\n")
		return 1
	}

	lineEnd := "\n"
	if *zero {
		lineEnd = "\x00"
	}

	for _, target := range targets {
		dir := filepath.Dir(target)
		fmt.Fprint(env.Stdout, dir, lineEnd)
	}

	return 0
}
