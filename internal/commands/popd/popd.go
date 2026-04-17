package popd

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Popd struct{}

func New() *Popd {
	return &Popd{}
}

func (p *Popd) Name() string {
	return "popd"
}

func (p *Popd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(env.DirStack) == 0 {
		fmt.Fprintln(env.Stderr, "popd: directory stack empty")
		return 1
	}

	env.Cwd = env.DirStack[len(env.DirStack)-1]
	env.DirStack = env.DirStack[:len(env.DirStack)-1]

	// Print stack (dirs style)
	stack := append([]string{env.Cwd}, env.DirStack...)
	for i, d := range stack {
		fmt.Fprint(env.Stdout, d)
		if i < len(stack)-1 {
			fmt.Fprint(env.Stdout, " ")
		}
	}
	fmt.Fprintln(env.Stdout, "")

	return 0
}
