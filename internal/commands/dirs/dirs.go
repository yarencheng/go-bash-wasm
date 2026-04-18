package dirs

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
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
	flags := pflag.NewFlagSet("dirs", pflag.ContinueOnError)
	clearStack := flags.BoolP("clear", "c", false, "clear the directory stack")
	onePerLine := flags.BoolP("print", "p", false, "print the directory stack with one entry per line")
	verbose := flags.BoolP("verbose", "v", false, "print the directory stack with one entry per line, prefixed with its position in the stack")
	_ = flags.BoolP("long", "l", false, "do not print the tilde-prefix (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "dirs: %v\n", err)
		}
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: dirs [-clpv] [+N] [-N]\n")
		fmt.Fprintf(env.Stdout, "Display directory stack.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "dirs")
		return 0
	}

	if *clearStack {
		env.DirStack = []string{}
		return 0
	}

	stack := append([]string{env.Cwd}, env.DirStack...)

	if *verbose {
		for i, dir := range stack {
			fmt.Fprintf(env.Stdout, "%2d  %s\n", i, dir)
		}
		return 0
	}

	if *onePerLine {
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
