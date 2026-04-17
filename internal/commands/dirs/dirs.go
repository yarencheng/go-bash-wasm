package dirs

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Dirs struct{}

func New() *Dirs {
	return &Dirs{}
}

func (d *Dirs) Name() string {
	return "dirs"
}

func (d *Dirs) Run(ctx context.Context, env *commands.Environment, args []string) int {
	stack := append([]string{env.Cwd}, env.DirStack...)
	for i, dir := range stack {
		fmt.Fprint(env.Stdout, dir)
		if i < len(stack)-1 {
			fmt.Fprint(env.Stdout, " ")
		}
	}
	fmt.Fprintln(env.Stdout, "")
	return 0
}
