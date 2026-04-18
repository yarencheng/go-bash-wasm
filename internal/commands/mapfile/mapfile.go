package mapfile

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Mapfile struct {
	name string
}

func New() *Mapfile {
	return &Mapfile{name: "mapfile"}
}

func NewWithName(name string) *Mapfile {
	return &Mapfile{name: name}
}

func (m *Mapfile) Name() string {
	return m.name
}

func (m *Mapfile) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet(m.name, pflag.ContinueOnError)
	trim := flags.BoolP("trim", "t", false, "remove trailing newline from each line")
	count := flags.IntP("count", "n", 0, "copy at most COUNT lines")
	origin := flags.IntP("origin", "O", 0, "begin assigning to array at index ORIGIN")
	fd := flags.IntP("fd", "u", 0, "read from file descriptor FD")
	callback := flags.StringP("callback", "C", "", "evaluate CALLBACK each time QUANTUM lines are read")
	quantum := flags.IntP("quantum", "c", 1, "number of lines to read between each call to CALLBACK")
	delim := flags.StringP("delimiter", "d", "\n", "use DELIM to terminate lines, instead of newline")
	skip := flags.IntP("skip", "s", 0, "discard the first SKIP lines read")
	
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "mapfile: %v\n", err)
		return 1
	}

	arrayName := "MAPFILE"
	if len(flags.Args()) > 0 {
		arrayName = flags.Args()[0]
	}

	var input io.Reader = env.Stdin
	if *fd != 0 {
		// Mock: we only support FD 0 for now.
		// In a real WASM env, FD might be tricky.
	}

	reader := bufio.NewReader(input)
	var lines []string

	existingLines := env.Arrays[arrayName]
	if *origin > 0 {
		if *origin > len(existingLines) {
			// Pad with empty strings
			padding := make([]string, *origin-len(existingLines))
			existingLines = append(existingLines, padding...)
		}
		lines = existingLines[:*origin]
	}

	d := byte('\n')
	if *delim != "" {
		d = (*delim)[0]
	}

	lineCount := 0
	readCount := 0
	for {
		if *count > 0 && lineCount >= *count {
			break
		}

		line, err := reader.ReadString(d)
		if line != "" {
			readCount++
			if readCount <= *skip {
				// Skip this line
			} else {
				if *trim {
					line = strings.TrimSuffix(line, string(d))
					if d == '\n' {
						line = strings.TrimSuffix(line, "\r")
					}
				}
				lines = append(lines, line)
				lineCount++

				if *callback != "" && lineCount%(*quantum) == 0 {
					// Mock: evaluate callback
					_ = env.Executor.Execute(ctx, *callback)
				}
			}
		}
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(env.Stderr, "mapfile: %v\n", err)
				return 1
			}
			break
		}
	}

	env.Arrays[arrayName] = lines
	return 0
}
