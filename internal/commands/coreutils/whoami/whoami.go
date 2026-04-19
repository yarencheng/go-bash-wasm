package whoami

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Whoami struct{}

func New() *Whoami {
	return &Whoami{}
}

func (w *Whoami) Name() string {
	return "whoami"
}

func (w *Whoami) Run(ctx context.Context, env *commands.Environment, args []string) int {
	fmt.Fprintln(env.Stdout, env.User)
	return 0
}
