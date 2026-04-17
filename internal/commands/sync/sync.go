package sync

import (
	"context"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Sync struct{}

func New() *Sync {
	return &Sync{}
}

func (s *Sync) Name() string {
	return "sync"
}

func (s *Sync) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// In our simulator, sync is a no-op as everything is in memory or handled by afero.
	return 0
}
