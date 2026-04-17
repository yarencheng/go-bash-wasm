package groups

import (
	"context"
	"fmt"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Groups struct{}

func New() *Groups {
	return &Groups{}
}

func (g *Groups) Name() string {
	return "groups"
}

func (g *Groups) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// Simple implementation for current user only
	groupStrs := make([]string, 0, len(env.Groups))
	for range env.Groups {
		groupStrs = append(groupStrs, env.User) // Simplified: group name = user name
	}
	fmt.Fprintln(env.Stdout, strings.Join(groupStrs, " "))
	return 0
}
