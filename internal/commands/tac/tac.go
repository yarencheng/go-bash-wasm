package tac

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Tac struct{}

func New() *Tac {
	return &Tac{}
}

func (t *Tac) Name() string {
	return "tac"
}

func (t *Tac) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("tac", pflag.ContinueOnError)
	_ = flags.BoolP("before", "b", false, "attach the separator before instead of after")
	_ = flags.BoolP("regex", "r", false, "interpret the separator as a regular expression")
	_ = flags.StringP("separator", "s", "\n", "use STRING as the separator instead of newline")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "tac: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return t.process(env, env.Stdin)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "tac: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := t.process(env, f); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

func (t *Tac) process(env *commands.Environment, r io.Reader) int {
	// Simple implementation: read all lines into memory
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(env.Stderr, "tac: %v\n", err)
		return 1
	}

	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Fprintln(env.Stdout, lines[i])
	}

	return 0
}
