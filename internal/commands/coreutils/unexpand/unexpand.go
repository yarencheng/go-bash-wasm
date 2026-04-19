package unexpand

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Unexpand struct{}

func New() *Unexpand {
	return &Unexpand{}
}

func (u *Unexpand) Name() string {
	return "unexpand"
}

func (u *Unexpand) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("unexpand", pflag.ContinueOnError)
	all := flags.BoolP("all", "a", false, "convert all blanks, instead of just initial blanks")
	tabs := flags.StringP("tabs", "t", "8", "have tabs N characters apart")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "unexpand: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: unexpand [OPTION]... [FILE]...\n")
		fmt.Fprintf(env.Stdout, "Convert spaces in each FILE to tabs, writing to standard output.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "unexpand")
		return 0
	}

	tabSize := 8
	if *tabs != "" {
		// simplify: only support single number for now
		if val, err := fmt.Sscanf(*tabs, "%d", &tabSize); err != nil || val != 1 {
			// ignore error or handle?
		}
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return u.process(env, env.Stdin, *all, tabSize)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "unexpand: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := u.process(env, f, *all, tabSize); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

func (u *Unexpand) process(env *commands.Environment, r io.Reader, all bool, tabSize int) int {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(env.Stdout, u.unexpandLine(line, all, tabSize))
	}
	return 0
}

func (u *Unexpand) unexpandLine(line string, all bool, tabSize int) string {
	var builder strings.Builder
	column := 0
	spaces := 0

	onlyInitial := !all

	for i := 0; i < len(line); i++ {
		c := line[i]
		if c == ' ' {
			spaces++
			column++
			if column%tabSize == 0 && spaces > 1 {
				builder.WriteByte('\t')
				spaces = 0
			}
		} else {
			// Flush spaces
			for j := 0; j < spaces; j++ {
				builder.WriteByte(' ')
			}
			spaces = 0
			builder.WriteByte(c)
			column++
			if onlyInitial {
				// Copy the rest of the line
				builder.WriteString(line[i+1:])
				break
			}
		}
	}

	// Final flush
	for j := 0; j < spaces; j++ {
		builder.WriteByte(' ')
	}

	return builder.String()
}
