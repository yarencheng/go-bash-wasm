package mkdir

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Mkdir struct{}

func New() *Mkdir {
	return &Mkdir{}
}

func (m *Mkdir) Name() string {
	return "mkdir"
}

func (m *Mkdir) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("mkdir", pflag.ContinueOnError)
	parents := flags.BoolP("parents", "p", false, "no error if existing, make parent directories as needed")
	verbose := flags.BoolP("verbose", "v", false, "print a message for each created directory")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "mkdir: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "mkdir: missing operand\n")
		return 1
	}

	exitCode := 0
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		var err error
		if *parents {
			err = env.FS.MkdirAll(fullPath, 0755)
		} else {
			err = env.FS.Mkdir(fullPath, 0755)
		}

		if err != nil {
			if strings.Contains(err.Error(), "already exists") && *parents {
				// According to GNU mkdir -p, this is not an error
			} else {
				fmt.Fprintf(env.Stderr, "mkdir: cannot create directory '%s': %v\n", target, err)
				exitCode = 1
				continue
			}
		}

		if *verbose {
			fmt.Fprintf(env.Stdout, "mkdir: created directory '%s'\n", target)
		}
	}

	return exitCode
}
