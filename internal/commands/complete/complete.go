package complete

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Complete struct{}

func New() *Complete {
	return &Complete{}
}

func (c *Complete) Name() string {
	return "complete"
}

func (c *Complete) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("complete", pflag.ContinueOnError)
	
	printFlag := flags.BoolP("print", "p", false, "print current completion settings")
	removeFlag := flags.BoolP("remove", "r", false, "remove completion settings")
	
	// Define other flags as boolean for simplicity in this simulator
	_ = flags.BoolP("alias", "a", false, "alias names")
	_ = flags.BoolP("builtin", "b", false, "builtin names")
	_ = flags.BoolP("command", "c", false, "command names")
	_ = flags.BoolP("directory", "d", false, "directory names")
	_ = flags.BoolP("export", "e", false, "exported variable names")
	_ = flags.BoolP("file", "f", false, "file names")
	_ = flags.BoolP("group", "g", false, "group names")
	_ = flags.BoolP("job", "j", false, "job names")
	_ = flags.BoolP("keyword", "k", false, "shell reserved word names")
	_ = flags.BoolP("service", "s", false, "service names")
	_ = flags.BoolP("user", "u", false, "user names")
	_ = flags.BoolP("variable", "v", false, "variable names")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "complete: %v\n", err)
		return 1
	}

	targets := flags.Args()

	if *printFlag {
		if len(targets) == 0 {
			for name, settings := range env.Completions {
				fmt.Fprintf(env.Stdout, "complete %s %s\n", renderSettings(settings), name)
			}
		} else {
			for _, name := range targets {
				if settings, ok := env.Completions[name]; ok {
					fmt.Fprintf(env.Stdout, "complete %s %s\n", renderSettings(settings), name)
				}
			}
		}
		return 0
	}

	if len(targets) == 0 {
		return 0
	}

	for _, name := range targets {
		if *removeFlag {
			delete(env.Completions, name)
			continue
		}

		// Store settings
		settings := make(map[string]string)
		flags.Visit(func(f *pflag.Flag) {
			if f.Name != "print" && f.Name != "remove" {
				settings[f.Name] = "true"
			}
		})
		env.Completions[name] = settings
	}

	return 0
}

func renderSettings(s map[string]string) string {
	res := ""
	for k := range s {
		res += "-" + k + " "
	}
	return res
}
