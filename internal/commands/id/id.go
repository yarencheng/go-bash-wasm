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
	f := i.defineFlags(flags)

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "id: %v\n", err)
		}
		return 1
	}

	if *f.contextFlag {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "id: context not supported\n")
		}
		return 1
	}

	if *f.userFlag {
		return i.printUser(env, &f)
	}
	if *f.groupFlag {
		return i.printGroup(env, &f)
	}
	if *f.groupsFlag {
		return i.printGroups(env, &f)
	}

	return i.printDefault(env)
}

type idFlags struct {
	userFlag    *bool
	groupFlag   *bool
	groupsFlag  *bool
	nameFlag    *bool
	realFlag    *bool
	contextFlag *bool
	zeroFlag    *bool
}

func (i *Id) defineFlags(flags *pflag.FlagSet) idFlags {
	f := idFlags{
		userFlag:    flags.BoolP("user", "u", false, "print only the effective user ID"),
		groupFlag:   flags.BoolP("group", "g", false, "print only the effective group ID"),
		groupsFlag:  flags.BoolP("groups", "G", false, "print all group IDs"),
		nameFlag:    flags.BoolP("name", "n", false, "print a name instead of a number"),
		realFlag:    flags.BoolP("real", "r", false, "print the real ID instead of the effective ID"),
		contextFlag: flags.BoolP("context", "Z", false, "print only the security context of the process"),
		zeroFlag:    flags.BoolP("zero", "z", false, "delimit entries with NUL characters, not whitespace"),
	}
	_ = flags.BoolP("all", "a", false, "ignore, for backward compatibility")
	return f
}

func (i *Id) printUser(env *commands.Environment, f *idFlags) int {
	if *f.nameFlag {
		fmt.Fprint(env.Stdout, env.User)
	} else {
		fmt.Fprint(env.Stdout, env.Uid)
	}
	i.printTerminator(env, f)
	return 0
}

func (i *Id) printGroup(env *commands.Environment, f *idFlags) int {
	if *f.nameFlag {
		fmt.Fprint(env.Stdout, env.User)
	} else {
		fmt.Fprint(env.Stdout, env.Gid)
	}
	i.printTerminator(env, f)
	return 0
}

func (i *Id) printGroups(env *commands.Environment, f *idFlags) int {
	sep := " "
	if *f.zeroFlag {
		sep = "\x00"
	}
	strs := []string{}
	for _, g := range env.Groups {
		if *f.nameFlag {
			strs = append(strs, env.User)
		} else {
			strs = append(strs, fmt.Sprintf("%d", g))
		}
	}
	fmt.Fprint(env.Stdout, strings.Join(strs, sep))
	i.printTerminator(env, f)
	return 0
}

func (i *Id) printDefault(env *commands.Environment) int {
	groupsStr := []string{}
	for _, g := range env.Groups {
		groupsStr = append(groupsStr, fmt.Sprintf("%d(%s)", g, env.User))
	}

	fmt.Fprintf(env.Stdout, "uid=%d(%s) gid=%d(%s) groups=%s\n",
		env.Uid, env.User, env.Gid, env.User, strings.Join(groupsStr, ","))
	return 0
}

func (i *Id) printTerminator(env *commands.Environment, f *idFlags) {
	if !*f.zeroFlag {
		fmt.Fprintln(env.Stdout)
	} else {
		fmt.Fprint(env.Stdout, "\x00")
	}
}
