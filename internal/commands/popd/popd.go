package popd

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Popd struct{}

func New() *Popd {
	return &Popd{}
}

func (p *Popd) Name() string {
	return "popd"
}

func (p *Popd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("popd", pflag.ContinueOnError)
	noChdir := flags.BoolP("no-chdir", "n", false, "suppress the normal change of directory")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "popd: %v\n", err)
		}
		return 1
	}

	if len(env.DirStack) == 0 {
		if env.Stderr != nil {
			fmt.Fprintln(env.Stderr, "popd: directory stack empty")
		}
		return 1
	}

	if *noChdir {
		// Remove from stack but don't change CWD
		// Usually removes the first element of DirStack (which is the one just before CWD)
		env.DirStack = env.DirStack[1:]
	} else {
		env.Cwd = env.DirStack[0]
		env.DirStack = env.DirStack[1:]
	}

	// Print stack (dirs style)
	stack := append([]string{env.Cwd}, env.DirStack...)
	for i, d := range stack {
		fmt.Fprint(env.Stdout, d)
		if i < len(stack)-1 {
			fmt.Fprint(env.Stdout, " ")
		}
	}
	fmt.Fprintln(env.Stdout, "")

	return 0
}
