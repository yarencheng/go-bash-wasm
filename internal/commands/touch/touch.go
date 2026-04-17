package touch

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

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
	access := flags.BoolP("access", "a", false, "change only the access time")
	modification := flags.BoolP("modification", "m", false, "change only the modification time")
	reference := flags.StringP("reference", "r", "", "use this file's times instead of current time")
	
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
	atime := time.Now()
	mtime := time.Now()

	if *reference != "" {
		fullRefPath := *reference
		if !filepath.IsAbs(fullRefPath) {
			fullRefPath = filepath.Join(env.Cwd, fullRefPath)
		}
		info, err := env.FS.Stat(fullRefPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "touch: cannot stat '%s': %v\n", *reference, err)
			return 1
		}
		atime = info.ModTime() // In many systems ModTime is same for both if not specified
		mtime = info.ModTime()
	}

	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		info, err := env.FS.Stat(fullPath)
		exists := err == nil
		
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
			info, _ = env.FS.Stat(fullPath)
		}

		targetAtime := atime
		targetMtime := mtime

		if *access && !*modification {
			targetMtime = info.ModTime()
		} else if *modification && !*access {
			targetAtime = info.ModTime() // Afero doesn't give us Atime easily, so we use mtime if we must
		}

		err = env.FS.Chtimes(fullPath, targetAtime, targetMtime)
		if err != nil {
			// Some filesystems might not support Chtimes
		}
	}

	return exitCode
}
