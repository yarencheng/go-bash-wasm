package fc

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Fc struct{}

func New() *Fc {
	return &Fc{}
}

func (f *Fc) Name() string {
	return "fc"
}

func (f *Fc) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("fc", pflag.ContinueOnError)
	list := flags.BoolP("list", "l", false, "list history")
	noNumbers := flags.BoolP("no-numbers", "n", false, "display without numbers")
	reverse := flags.BoolP("reverse", "r", false, "reverse listing")
	reExecute := flags.BoolP("re-execute", "s", false, "re-execute")
	editor := flags.StringP("editor", "e", "vi", "editor name")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "fc: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: fc [-e ename] [-lnr] [first] [last]\n")
		fmt.Fprintf(env.Stdout, "  or:  fc -s [pat=rep] [command]\n")
		fmt.Fprintf(env.Stdout, "Fix Command.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "fc")
		return 0
	}

	history := env.History
	if len(history) == 0 {
		return 0
	}

	if *reExecute {
		// Re-execute last command or specified command
		targetLine := history[len(history)-1]
		if flags.NArg() > 0 {
			arg := flags.Arg(0)
			if idx, err := strconv.Atoi(arg); err == nil {
				if idx > 0 && idx <= len(history) {
					targetLine = history[idx-1]
				} else if idx < 0 && len(history)+idx > 0 {
					targetLine = history[len(history)+idx]
				}
			} else {
				// Search for string prefix
				for i := len(history) - 1; i >= 0; i-- {
					if strings.HasPrefix(history[i], arg) {
						targetLine = history[i]
						break
					}
				}
			}
		}

		if env.Executor != nil {
			return env.Executor.Execute(ctx, targetLine)
		}
		return 0
	}

	if *list {
		first := 1
		last := len(history)

		if flags.NArg() > 0 {
			if f, err := strconv.Atoi(flags.Arg(0)); err == nil {
				first = f
				if first < 0 {
					first = len(history) + first + 1
				}
			}
		}
		if flags.NArg() > 1 {
			if l, err := strconv.Atoi(flags.Arg(1)); err == nil {
				last = l
				if last < 0 {
					last = len(history) + last + 1
				}
			}
		}

		if first > len(history) {
			first = len(history)
		}
		if first < 1 {
			first = 1
		}
		if last > len(history) {
			last = len(history)
		}
		if last < 1 {
			last = 1
		}

		printLine := func(i int) {
			if *noNumbers {
				fmt.Fprintln(env.Stdout, history[i-1])
			} else {
				fmt.Fprintf(env.Stdout, "%d\t%s\n", i, history[i-1])
			}
		}

		if *reverse {
			if first <= last {
				for i := last; i >= first; i-- {
					printLine(i)
				}
			} else {
				for i := last; i <= first; i++ {
					printLine(i)
				}
			}
		} else {
			if first <= last {
				for i := first; i <= last; i++ {
					printLine(i)
				}
			} else {
				for i := first; i >= last; i-- {
					printLine(i)
				}
			}
		}
		return 0
	}

	// Default: edit and re-execute
	// Interactive editing not supported in WASM, so we just re-execute for now
	// or print a warning
	fmt.Fprintf(env.Stderr, "fc: interactive editing with '%s' not supported in this environment\n", *editor)
	return 1
}
