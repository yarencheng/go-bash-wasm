package stat

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Stat struct{}

func New() *Stat {
	return &Stat{}
}

func (s *Stat) Name() string {
	return "stat"
}

func (s *Stat) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("stat", pflag.ContinueOnError)
	_ = flags.BoolP("file-system", "f", false, "display file system status instead of file status (ignored)")
	_ = flags.BoolP("terse", "t", false, "print the information in terse form (ignored)")
	
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "stat: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "stat: missing operand\n")
		return 1
	}

	exitCode := 0
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		info, err := env.FS.Stat(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "stat: cannot stat '%s': %v\n", target, err)
			exitCode = 1
			continue
		}

		fmt.Fprintf(env.Stdout, "  File: %s\n", target)
		fmt.Fprintf(env.Stdout, "  Size: %-10d Mode: %s\n", info.Size(), info.Mode())
		fmt.Fprintf(env.Stdout, "Modify: %s\n", info.ModTime().Format("2006-01-02 15:04:05.000000000 -0700"))
	}

	return exitCode
}
