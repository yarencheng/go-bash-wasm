package compgen

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Compgen struct{}

func New() *Compgen {
	return &Compgen{}
}

func (c *Compgen) Name() string {
	return "compgen"
}

func (c *Compgen) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("compgen", pflag.ContinueOnError)
	_ = flags.BoolP("alias", "a", false, "alias names")
	_ = flags.BoolP("builtin", "b", false, "builtin names")
	_ = flags.BoolP("command", "c", false, "command names")
	_ = flags.BoolP("directory", "d", false, "directory names")
	_ = flags.BoolP("file", "f", false, "file names")
	_ = flags.StringP("wordlist", "W", "", "wordlist")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "compgen: %v\n", err)
		return 1
	}

	word := ""
	if flags.NArg() > 0 {
		word = flags.Arg(0)
	}

	// Mock generation
	var matches []string

	// If wordlist is provided
	w, _ := flags.GetString("wordlist")
	if w != "" {
		for _, item := range strings.Fields(w) {
			if strings.HasPrefix(item, word) {
				matches = append(matches, item)
			}
		}
	}

	// Basic command matching if -c is set
	if cFlag, _ := flags.GetBool("command"); cFlag {
		for _, name := range env.Registry.List() {
			if strings.HasPrefix(name, word) {
				matches = append(matches, name)
			}
		}
	}

	for _, m := range matches {
		fmt.Fprintln(env.Stdout, m)
	}

	return 0
}
