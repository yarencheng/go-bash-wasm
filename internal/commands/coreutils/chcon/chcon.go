package chcon

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Chcon struct{}

func New() *Chcon {
	return &Chcon{}
}

func (c *Chcon) Name() string {
	return "chcon"
}

func (c *Chcon) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("chcon", pflag.ContinueOnError)
	_ = flags.BoolP("no-dereference", "h", false, "affect symbolic links instead of any referenced file")
	_ = flags.BoolP("H", "H", false, "if a command line argument is a symbolic link to a directory, traverse it")
	_ = flags.BoolP("L", "L", false, "traverse every symbolic link to a directory encountered")
	_ = flags.BoolP("P", "P", false, "do not traverse any symbolic links (default)")
	_ = flags.BoolP("recursive", "R", false, "operate on files and directories recursively")
	_ = flags.StringP("user", "u", "", "set user USER in the target security context")
	_ = flags.StringP("role", "r", "", "set role ROLE in the target security context")
	_ = flags.StringP("type", "t", "", "set type TYPE in the target security context")
	_ = flags.StringP("range", "l", "", "set range RANGE in the target security context")
	_ = flags.String("reference", "", "use RFILE's security context rather than specifying a CONTEXT value")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "chcon: %v\n", err)
		return 1
	}

	// Stub: we don't support SELinux contexts.
	fmt.Fprintf(env.Stderr, "chcon: SELinux not supported in this environment\n")
	return 1
}
