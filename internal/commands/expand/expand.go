package expand

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Expand struct{}

func New() *Expand {
	return &Expand{}
}

func (e *Expand) Name() string {
	return "expand"
}

func (e *Expand) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("expand", pflag.ContinueOnError)
	initial := flags.BoolP("initial", "i", false, "convert only leading tabs")
	tabs := flags.StringP("tabs", "t", "8", "have tabs N characters apart")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "expand: %v\n", err)
		return 1
	}

	tabSize := 8
	if *tabs != "" {
		if val, err := strconv.Atoi(*tabs); err == nil {
			tabSize = val
		}
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return e.process(env, env.Stdin, *initial, tabSize)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "expand: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := e.process(env, f, *initial, tabSize); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

func (e *Expand) process(env *commands.Environment, r io.Reader, initial bool, tabSize int) int {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(env.Stdout, e.expandLine(line, initial, tabSize))
	}
	return 0
}

func (e *Expand) expandLine(line string, initial bool, tabSize int) string {
	var builder strings.Builder
	column := 0
	
	onlyInitial := initial

	for i := 0; i < len(line); i++ {
		c := line[i]
		if c == '\t' {
			spaces := tabSize - (column % tabSize)
			for j := 0; j < spaces; j++ {
				builder.WriteByte(' ')
			}
			column += spaces
		} else {
			builder.WriteByte(c)
			column++
			if onlyInitial && c != ' ' {
				// Copy the rest of the line
				builder.WriteString(line[i+1:])
				break
			}
		}
	}

	return builder.String()
}
