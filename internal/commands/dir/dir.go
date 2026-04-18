package dir

import (
	"context"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ls"
)

type Dir struct {
	ls *ls.Ls
}

func New() *Dir {
	return &Dir{ls: ls.New()}
}

func (d *Dir) Name() string {
	return "dir"
}

func (d *Dir) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// 'dir' is equivalent to 'ls -C -b' (multi-column, escaped)
	// For simulator, we just call ls. Default for ls in simulator is often one-line or similar.
	// We'll prepend -C if no other format is specified.
	return d.ls.Run(ctx, env, append([]string{"-C"}, args...))
}
