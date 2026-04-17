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
	if len(args) == 0 {
		return g.printGroups(env, env.User)
	}

	exitCode := 0
	for _, user := range args {
		if user == env.User {
			if err := g.printGroups(env, user); err != 0 {
				exitCode = err
			}
		} else {
			fmt.Fprintf(env.Stderr, "groups: '%s': no such user\n", user)
			exitCode = 1
		}
	}
	return exitCode
}

func (g *Groups) printGroups(env *commands.Environment, user string) int {
	groupStrs := make([]string, 0, len(env.Groups))
	for range env.Groups {
		groupStrs = append(groupStrs, user) // Simplified: group name = user name
	}
	if len(groupStrs) == 0 {
		groupStrs = append(groupStrs, user)
	}
	if len(env.Groups) > 0 {
		// Output format when multiple users: USER : GROUP...
		// But GNU groups output: USER : GROUP... if multiple args?
		// Actually if single arg, it prints: GROUP...
		// If multiple args, it prints: USER : GROUP... NO, GNU groups prints each on a new line.
	}
	fmt.Fprintln(env.Stdout, strings.Join(groupStrs, " "))
	return 0
}
