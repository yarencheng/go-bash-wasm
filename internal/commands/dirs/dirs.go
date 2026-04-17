package dirs

import (
	"context"
	"fmt"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Dirs struct{}

func New() *Dirs {
	return &Dirs{}
}

func (d *Dirs) Name() string {
	return "dirs"
}

func (d *Dirs) Run(ctx context.Context, env *commands.Environment, args []string) int {
	clearStack := false
	onePerLine := false
	verbose := false

	for _, arg := range args {
		switch arg {
		case "-c":
			clearStack = true
		case "-p":
			onePerLine = true
		case "-v":
			verbose = true
		}
	}

	if clearStack {
		env.DirStack = []string{}
		return 0
	}

	stack := append([]string{env.Cwd}, env.DirStack...)

	if verbose {
		for i, dir := range stack {
			fmt.Fprintf(env.Stdout, "%2d  %s\n", i, dir)
		}
		return 0
	}

	if onePerLine {
		for _, dir := range stack {
			fmt.Fprintln(env.Stdout, dir)
		}
		return 0
	}

	for i, dir := range stack {
		fmt.Fprint(env.Stdout, dir)
		if i < len(stack)-1 {
			fmt.Fprint(env.Stdout, " ")
		}
	}
	fmt.Fprintln(env.Stdout, "")
	return 0
}
