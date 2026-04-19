package compgen

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/afero"
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

	// File and directory matching
	fFlag, _ := flags.GetBool("file")
	dFlag, _ := flags.GetBool("directory")
	if fFlag || dFlag {
		dir := "."
		base := word
		if strings.Contains(word, "/") {
			idx := strings.LastIndex(word, "/")
			dir = word[:idx]
			if dir == "" {
				dir = "/"
			}
			base = word[idx+1:]
		}

		entries, err := afero.ReadDir(env.FS, dir)
		if err == nil {
			for _, entry := range entries {
				if strings.HasPrefix(entry.Name(), base) {
					if dFlag && !entry.IsDir() {
						continue
					}
					match := entry.Name()
					if dir != "." && dir != "/" {
						match = dir + "/" + entry.Name()
					} else if dir == "/" {
						match = "/" + entry.Name()
					}
					if entry.IsDir() {
						match += "/"
					}
					matches = append(matches, match)
				}
			}
		}
	}

	for _, m := range matches {
		fmt.Fprintln(env.Stdout, m)
	}

	return 0
}
