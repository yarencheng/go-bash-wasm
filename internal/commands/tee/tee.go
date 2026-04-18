package tee

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Tee struct{}

func New() *Tee {
	return &Tee{}
}

func (t *Tee) Name() string {
	return "tee"
}

func (t *Tee) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("tee", pflag.ContinueOnError)
	appendFlag := flags.BoolP("append", "a", false, "append to the given FILEs, do not overwrite")
	_ = flags.BoolP("ignore-interrupts", "i", false, "ignore interrupt signals")
	_ = flags.BoolP("output-error", "p", false, "diagnose errors writing to non pipes")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "tee: %v\n", err)
		return 1
	}

	targets := flags.Args()
	var writers []io.Writer
	writers = append(writers, env.Stdout)

	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		openFlags := os.O_CREATE | os.O_WRONLY
		if *appendFlag {
			openFlags |= os.O_APPEND
		} else {
			openFlags |= os.O_TRUNC
		}

		f, err := env.FS.OpenFile(fullPath, openFlags, 0644)
		if err != nil {
			fmt.Fprintf(env.Stderr, "tee: %v\n", err)
			continue
		}
		defer f.Close()
		writers = append(writers, f)
	}

	multi := io.MultiWriter(writers...)
	if _, err := io.Copy(multi, env.Stdin); err != nil {
		fmt.Fprintf(env.Stderr, "tee: %v\n", err)
		return 1
	}

	return 0
}
