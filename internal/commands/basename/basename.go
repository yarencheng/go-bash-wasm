package basename

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Basename struct{}

func New() *Basename {
	return &Basename{}
}

func (b *Basename) Name() string {
	return "basename"
}

func (b *Basename) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("basename", pflag.ContinueOnError)
	multiple := flags.BoolP("multiple", "a", false, "support multiple arguments and treat each as a NAME")
	suffix := flags.StringP("suffix", "s", "", "remove a trailing SUFFIX")
	zero := flags.BoolP("zero", "z", false, "end each output line with NUL, not newline")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "basename: %v\n", err)
		return 1
	}

	posArgs := flags.Args()
	if len(posArgs) == 0 {
		fmt.Fprintf(env.Stderr, "basename: missing operand\n")
		return 1
	}

	targets := posArgs
	suffixToTrim := *suffix

	if !*multiple && suffixToTrim == "" {
		if len(posArgs) > 1 {
			targets = posArgs[:1]
			suffixToTrim = posArgs[1]
		}
	}

	lineEnd := "\n"
	if *zero {
		lineEnd = "\x00"
	}

	for _, target := range targets {
		name := filepath.Base(target)
		if suffixToTrim != "" && strings.HasSuffix(name, suffixToTrim) {
			name = name[:len(name)-len(suffixToTrim)]
		}
		fmt.Fprint(env.Stdout, name, lineEnd)
	}

	return 0
}
