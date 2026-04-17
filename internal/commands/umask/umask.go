package umask

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Umask struct{}

func New() *Umask {
	return &Umask{}
}

func (u *Umask) Name() string {
	return "umask"
}

func (u *Umask) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("umask", pflag.ContinueOnError)
	symbolic := flags.BoolP("symbolic", "S", false, "display umask in symbolic format")
	printCode := flags.BoolP("print", "p", false, "display umask in a form that can be reused as input")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "umask: %v\n", err)
		}
		return 2
	}

	remaining := flags.Args()

	if len(remaining) == 0 {
		if *symbolic {
			u0 := (env.Umask >> 6) & 7
			u1 := (env.Umask >> 3) & 7
			u2 := env.Umask & 7
			fmt.Fprintf(env.Stdout, "u=%s,g=%s,o=%s\n", formatSymbolic(u0), formatSymbolic(u1), formatSymbolic(u2))
		} else if *printCode {
			fmt.Fprintf(env.Stdout, "umask %04o\n", env.Umask)
		} else {
			fmt.Fprintf(env.Stdout, "%04o\n", env.Umask)
		}
		return 0
	}

	// Set umask
	newMask, err := strconv.ParseUint(remaining[0], 8, 32)
	if err != nil {
		// Try symbolic or more complex parsing if needed, but for now octal only
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "umask: %s: invalid octal number\n", remaining[0])
		}
		return 1
	}

	if newMask > 0777 {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "umask: %s: octal number out of range\n", remaining[0])
		}
		return 1
	}

	env.Umask = uint32(newMask)
	return 0
}

func formatSymbolic(val uint32) string {
	// umask is mask, so 0 means allowed, 7 means blocked.
	// Actually umask 022 means 755 permissions.
	// So R=4, W=2, X=1.
	// If val=0 (allowed), permissions=7 (rwx).
	// If val=2 (w blocked), permissions=5 (rx).
	perm := 7 - val
	res := ""
	if perm&4 != 0 {
		res += "r"
	}
	if perm&2 != 0 {
		res += "w"
	}
	if perm&1 != 0 {
		res += "x"
	}
	return res
}
