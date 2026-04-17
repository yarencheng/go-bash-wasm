package yes

import (
	"context"
	"fmt"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Yes struct{}

func New() *Yes {
	return &Yes{}
}

func (y *Yes) Name() string {
	return "yes"
}

func (y *Yes) Run(ctx context.Context, env *commands.Environment, args []string) int {
	text := "y"
	if len(args) > 0 {
		text = strings.Join(args, " ")
	}

	for {
		select {
		case <-ctx.Done():
			return 0
		default:
			fmt.Fprintln(env.Stdout, text)
		}
	}
}
