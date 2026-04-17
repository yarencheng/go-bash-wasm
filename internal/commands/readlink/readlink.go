package readlink

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Readlink struct{}

func New() *Readlink {
	return &Readlink{}
}

func (r *Readlink) Name() string {
	return "readlink"
}

func (r *Readlink) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("readlink", pflag.ContinueOnError)
	noNewline := flags.BoolP("no-newline", "n", false, "do not output the trailing newline")
	_ = flags.BoolP("canonicalize", "f", false, "canonicalize by following every symlink")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "readlink: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		return 1
	}

	target := targets[0]
	if !filepath.IsAbs(target) {
		target = filepath.Join(env.Cwd, target)
	}

	linker, ok := env.FS.(afero.Symlinker)
	if !ok {
		// If not a symlinker, it can't have symlinks
		return 1
	}

	linkTarget, err := linker.ReadlinkIfPossible(target)
	if err != nil {
		return 1
	}

	fmt.Fprint(env.Stdout, linkTarget)
	if !*noNewline {
		fmt.Fprintln(env.Stdout)
	}

	return 0
}
