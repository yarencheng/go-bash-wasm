package fold

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Fold struct{}

func New() *Fold {
	return &Fold{}
}

func (f *Fold) Name() string {
	return "fold"
}

func (f *Fold) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("fold", pflag.ContinueOnError)
	bytesFlag := flags.BoolP("bytes", "b", false, "count bytes instead of columns")
	charsFlag := flags.BoolP("characters", "c", false, "count characters instead of columns")
	spaces := flags.BoolP("spaces", "s", false, "break at spaces")
	width := flags.StringP("width", "w", "80", "maximum line width")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "fold: %v\n", err)
		return 1
	}

	maxWidth := 80
	if *width != "" {
		if val, err := strconv.Atoi(*width); err == nil {
			maxWidth = val
		}
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return f.process(env, env.Stdin, maxWidth, *bytesFlag || *charsFlag, *spaces)
	}

	exitCode := 0
	for _, arg := range remaining {
		file, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "fold: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := f.process(env, file, maxWidth, *bytesFlag || *charsFlag, *spaces); status != 0 {
			exitCode = status
		}
		file.Close()
	}

	return exitCode
}

func (f *Fold) process(env *commands.Environment, r io.Reader, maxWidth int, bytesMode, breakAtSpaces bool) int {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		f.foldLine(env, line, maxWidth, bytesMode, breakAtSpaces)
	}
	return 0
}

func (f *Fold) foldLine(env *commands.Environment, line string, maxWidth int, bytesMode, breakAtSpaces bool) {
	if len(line) == 0 {
		fmt.Fprintln(env.Stdout)
		return
	}

	if bytesMode {
		data := []byte(line)
		for len(data) > 0 {
			if len(data) <= maxWidth {
				env.Stdout.Write(data)
				fmt.Fprintln(env.Stdout)
				return
			}

			breakIdx := maxWidth
			if breakAtSpaces {
				lastSpace := bytes.LastIndex(data[:maxWidth+1], []byte(" "))
				if lastSpace != -1 {
					breakIdx = lastSpace + 1
				}
			}

			env.Stdout.Write(data[:breakIdx])
			fmt.Fprintln(env.Stdout)
			data = data[breakIdx:]
		}
	} else {
		runes := []rune(line)
		for len(runes) > 0 {
			if len(runes) <= maxWidth {
				fmt.Fprintln(env.Stdout, string(runes))
				return
			}

			breakIdx := maxWidth
			if breakAtSpaces {
				lastSpace := -1
				for i := 0; i <= maxWidth && i < len(runes); i++ {
					if runes[i] == ' ' {
						lastSpace = i
					}
				}
				if lastSpace != -1 {
					breakIdx = lastSpace + 1
				}
			}

			fmt.Fprintln(env.Stdout, string(runes[:breakIdx]))
			runes = runes[breakIdx:]
		}
	}
}

