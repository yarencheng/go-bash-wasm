package logout

import (
	"context"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Logout struct{}

func New() *Logout {
	return &Logout{}
}

func (l *Logout) Name() string {
	return "logout"
}

func (l *Logout) Run(ctx context.Context, env *commands.Environment, args []string) int {
	env.ExitRequested = true
	return 0
}
