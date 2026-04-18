package tac

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

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
	before := flags.BoolP("before", "b", false, "attach the separator before instead of after")
	_ = flags.BoolP("regex", "r", false, "interpret the separator as a regular expression (not implemented)")
	separator := flags.StringP("separator", "s", "\n", "use STRING as the separator instead of newline")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "tac: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return t.process(env, env.Stdin, *separator, *before)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "tac: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := t.process(env, f, *separator, *before); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

func (t *Tac) process(env *commands.Environment, r io.Reader, sep string, before bool) int {
	var chunks []string

	if sep == "\n" {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			chunks = append(chunks, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(env.Stderr, "tac: %v\n", err)
			return 1
		}
		for i := len(chunks) - 1; i >= 0; i-- {
			fmt.Fprintln(env.Stdout, chunks[i])
		}
		return 0
	}

	// Custom separator
	data, err := io.ReadAll(r)
	if err != nil {
		fmt.Fprintf(env.Stderr, "tac: %v\n", err)
		return 1
	}

	parts := strings.Split(string(data), sep)
	// If it ends with sep, the last part is empty.
	// GNU tac: if input ends with sep, we should keep it.

	for i := len(parts) - 1; i >= 0; i-- {
		if i == len(parts)-1 && parts[i] == "" {
			continue
		}
		if before {
			fmt.Fprint(env.Stdout, sep+parts[i])
		} else {
			fmt.Fprint(env.Stdout, parts[i]+sep)
		}
	}

	return 0
}
