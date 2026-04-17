package read

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Read struct{}

func New() *Read {
	return &Read{}
}

func (r *Read) Name() string {
	return "read"
}

func (r *Read) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("read", pflag.ContinueOnError)
	prompt := flags.StringP("prompt", "p", "", "display PROMPT without a trailing newline, before attempting to read")
	// -s is tricky in WASM/browser, skipped for now but defined
	_ = flags.BoolP("silent", "s", false, "do not echo input coming from a terminal")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "read: %v\n", err)
		return 1
	}

	if *prompt != "" {
		fmt.Fprint(env.Stdout, *prompt)
	}

	reader := bufio.NewReader(env.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && err != fmt.Errorf("EOF") && line == "" {
		return 1
	}

	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")

	fields := flags.Args()
	if len(fields) == 0 {
		env.EnvVars["REPLY"] = line
		return 0
	}

	words := strings.Fields(line)
	for i, field := range fields {
		if i < len(words) {
			if i == len(fields)-1 {
				// Last field gets the rest of the line
				// Find start index of this word in original line
				// This is a bit simplified, but close enough for now
				env.EnvVars[field] = strings.Join(words[i:], " ")
			} else {
				env.EnvVars[field] = words[i]
			}
		} else {
			env.EnvVars[field] = ""
		}
	}

	return 0
}
