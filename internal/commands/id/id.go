package id

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Id struct{}

func New() *Id {
	return &Id{}
}

func (i *Id) Name() string {
	return "id"
}

func (i *Id) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("id", pflag.ContinueOnError)
	userFlag := flags.BoolP("user", "u", false, "print only the effective user ID")
	groupFlag := flags.BoolP("group", "g", false, "print only the effective group ID")
	nameFlag := flags.BoolP("name", "n", false, "print a name instead of a number")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "id: %v\n", err)
		return 1
	}

	if *userFlag {
		if *nameFlag {
			fmt.Fprintln(env.Stdout, env.User)
		} else {
			fmt.Fprintln(env.Stdout, env.Uid)
		}
		return 0
	}

	if *groupFlag {
		if *nameFlag {
			fmt.Fprintln(env.Stdout, env.User) // Simplification: common name
		} else {
			fmt.Fprintln(env.Stdout, env.Gid)
		}
		return 0
	}

	// Default: uid=1000(wasm) gid=1000(wasm) groups=1000(wasm)
	groupsStr := []string{}
	for _, g := range env.Groups {
		groupsStr = append(groupsStr, fmt.Sprintf("%d(%s)", g, env.User))
	}

	fmt.Fprintf(env.Stdout, "uid=%d(%s) gid=%d(%s) groups=%s\n", 
		env.Uid, env.User, env.Gid, env.User, strings.Join(groupsStr, ","))
	return 0
}
