package sync

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
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
	flags := pflag.NewFlagSet("sync", pflag.ContinueOnError)
	_ = flags.BoolP("data", "d", false, "sync only file data, no unneeded metadata")
	_ = flags.BoolP("file-system", "f", false, "sync the file systems containing the files")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "sync: %v\n", err)
		return 1
	}

	// In our simulator, sync is a no-op as everything is in memory or handled by afero.
	return 0
}
