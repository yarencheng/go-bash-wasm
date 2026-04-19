package who

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Who struct{}

func New() *Who {
	return &Who{}
}

func (w *Who) Name() string {
	return "who"
}

func (w *Who) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("who", pflag.ContinueOnError)
	_ = flags.BoolP("all", "a", false, "print all information")
	_ = flags.BoolP("boot", "b", false, "time of last system boot")
	_ = flags.BoolP("dead", "d", false, "print dead processes")
	_ = flags.BoolP("login", "l", false, "print system login processes")
	_ = flags.BoolP("process", "p", false, "print active processes spawned by init")
	_ = flags.BoolP("count", "q", false, "all login names and number of users logged on")
	_ = flags.BoolP("runlevel", "r", false, "current runlevel")
	_ = flags.BoolP("short", "s", false, "print only name, line, and time (default)")
	_ = flags.BoolP("time", "t", false, "print last system clock change")
	_ = flags.BoolP("users", "u", false, "list users logged in")
	heading := flags.BoolP("heading", "H", false, "print line of column headings")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "who: %v\n", err)
		return 1
	}

	if *heading {
		fmt.Fprintf(env.Stdout, "%-10s %-10s %s\n", "NAME", "LINE", "TIME")
	}

	// Simulation for virtual environment
	fmt.Fprintf(env.Stdout, "%-10s %-10s %s\n", env.User, "tty1", env.StartTime.Format("2006-01-02 15:04"))
	return 0
}
