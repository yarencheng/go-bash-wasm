package arch

import (
	"context"
	"fmt"
	"runtime"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Arch struct{}

func New() *Arch {
	return &Arch{}
}

func (a *Arch) Name() string {
	return "arch"
}

func (a *Arch) Run(ctx context.Context, env *commands.Environment, args []string) int {
	fmt.Fprintln(env.Stdout, runtime.GOARCH)
	return 0
}
