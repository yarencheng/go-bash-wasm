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
	groupsFlag := flags.BoolP("groups", "G", false, "print all group IDs")
	nameFlag := flags.BoolP("name", "n", false, "print a name instead of a number")
	realFlag := flags.BoolP("real", "r", false, "print the real ID instead of the effective ID")
	contextFlag := flags.BoolP("context", "Z", false, "print only the security context of the process")
	_ = flags.BoolP("all", "a", false, "ignore, for backward compatibility")
	zeroFlag := flags.BoolP("zero", "z", false, "delimit entries with NUL characters, not whitespace")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "id: %v\n", err)
		}
		return 1
	}

	// Use realFlag to satisfy compiler; currently real == effective in this simulator
	_ = *realFlag

	if *contextFlag {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "id: context not supported\n")
		}
		return 1
	}

	sep := " "
	if *zeroFlag {
		sep = "\x00"
	}

	if *userFlag {
		if *nameFlag {
			fmt.Fprint(env.Stdout, env.User)
		} else {
			fmt.Fprint(env.Stdout, env.Uid)
		}
		if !*zeroFlag {
			fmt.Fprintln(env.Stdout)
		} else {
			fmt.Fprint(env.Stdout, "\x00")
		}
		return 0
	}

	if *groupFlag {
		if *nameFlag {
			fmt.Fprint(env.Stdout, env.User) // Simplification: common name
		} else {
			fmt.Fprint(env.Stdout, env.Gid)
		}
		if !*zeroFlag {
			fmt.Fprintln(env.Stdout)
		} else {
			fmt.Fprint(env.Stdout, "\x00")
		}
		return 0
	}

	if *groupsFlag {
		strs := []string{}
		for _, g := range env.Groups {
			if *nameFlag {
				strs = append(strs, env.User)
			} else {
				strs = append(strs, fmt.Sprintf("%d", g))
			}
		}
		fmt.Fprint(env.Stdout, strings.Join(strs, sep))
		if !*zeroFlag {
			fmt.Fprintln(env.Stdout)
		} else {
			fmt.Fprint(env.Stdout, "\x00")
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
