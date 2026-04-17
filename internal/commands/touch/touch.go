package touch

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Touch struct{}

func New() *Touch {
	return &Touch{}
}

func (t *Touch) Name() string {
	return "touch"
}

func (t *Touch) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("touch", pflag.ContinueOnError)
	noCreate := flags.BoolP("no-create", "c", false, "do not create any files")
	
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "touch: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "touch: missing file operand\n")
		return 1
	}

	exitCode := 0
	currentTime := time.Now()

	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		exists, err := afero.Exists(env.FS, fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "touch: cannot stat '%s': %v\n", target, err)
			exitCode = 1
			continue
		}

		if !exists {
			if *noCreate {
				continue
			}
			f, err := env.FS.Create(fullPath)
			if err != nil {
				fmt.Fprintf(env.Stderr, "touch: cannot create '%s': %v\n", target, err)
				exitCode = 1
				continue
			}
			f.Close()
		} else {
			err = env.FS.Chtimes(fullPath, currentTime, currentTime)
			if err != nil {
				// Some filesystems might not support Chtimes, ignore or report
			}
		}
	}

	return exitCode
}
