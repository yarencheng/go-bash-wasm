package pwd

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Pwd struct{}

func New() *Pwd {
	return &Pwd{}
}

func (p *Pwd) Name() string {
	return "pwd"
}

func (p *Pwd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// TODO: Handle -L and -P if necessary. For now, we return Cwd.
	fmt.Fprintln(env.Stdout, env.Cwd)
	return 0
}
