package unlink

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Unlink struct{}

func New() *Unlink {
	return &Unlink{}
}

func (u *Unlink) Name() string {
	return "unlink"
}

func (u *Unlink) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("unlink", pflag.ContinueOnError)
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "unlink: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) != 1 {
		fmt.Fprintf(env.Stderr, "unlink: missing operand or too many operands\n")
		return 1
	}

	target := remaining[0]
	fullPath := target
	if !filepath.IsAbs(target) {
		fullPath = filepath.Join(env.Cwd, target)
	}

	err := env.FS.Remove(fullPath)
	if err != nil {
		fmt.Fprintf(env.Stderr, "unlink: cannot unlink '%s': %v\n", target, err)
		return 1
	}

	return 0
}
