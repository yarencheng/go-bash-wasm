package help

import (
	"context"
	"fmt"
	"sort"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Help struct{}

func New() *Help {
	return &Help{}
}

func (h *Help) Name() string {
	return "help"
}

func (h *Help) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("help", pflag.ContinueOnError)
	short := flags.BoolP("short", "d", false, "output short description of each topic")
	man := flags.BoolP("man", "m", false, "display usage in pseudo-manpage format")
	syntax := flags.BoolP("syntax", "s", false, "output only a short usage synopsis for each topic")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "help: %v\n", err)
		return 1
	}

	targets := flags.Args()
	_ = short
	_ = man
	_ = syntax
	if len(targets) == 0 {
		fmt.Fprintln(env.Stdout, "This is a Go-based Bash simulator (WASM).")
		fmt.Fprintln(env.Stdout, "These shell commands are defined internally.  Type `help' to see this list.")
		fmt.Fprintln(env.Stdout, "")
		
		cmds := env.Registry.List()
		sort.Strings(cmds)
		
		// Print in columns
		for i, name := range cmds {
			fmt.Fprintf(env.Stdout, "%-15s", name)
			if (i+1)%5 == 0 {
				fmt.Fprintln(env.Stdout, "")
			}
		}
		if len(cmds)%5 != 0 {
			fmt.Fprintln(env.Stdout, "")
		}
		return 0
	}

	for _, name := range targets {
		if cmd, ok := env.Registry.Get(name); ok {
			fmt.Fprintf(env.Stdout, "%s: %s help info...\n", name, name)
			// Ideally we would have a 'Help()' method on Command interface
			// but for now we just acknowledge it exists.
			_ = cmd.Run(ctx, env, []string{"--help"})
		} else {
			fmt.Fprintf(env.Stderr, "help: no help topics match '%s'.  Try 'help help' or 'man -k %s' or 'info %s'.\n", name, name, name)
		}
	}

	return 0
}
