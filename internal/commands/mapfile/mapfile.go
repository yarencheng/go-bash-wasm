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

type Mapfile struct{}

func New() *Mapfile {
	return &Mapfile{}
}

func (m *Mapfile) Name() string {
	return "mapfile"
}

func (m *Mapfile) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("mapfile", pflag.ContinueOnError)
	trim := flags.BoolP("trim", "t", false, "remove trailing newline from each line")
	
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "mapfile: %v\n", err)
		return 1
	}

	arrayName := "MAPFILE"
	if len(flags.Args()) > 0 {
		arrayName = flags.Args()[0]
	}

	reader := bufio.NewReader(env.Stdin)
	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			if *trim {
				line = strings.TrimSuffix(line, "\n")
				line = strings.TrimSuffix(line, "\r")
			}
			lines = append(lines, line)
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
