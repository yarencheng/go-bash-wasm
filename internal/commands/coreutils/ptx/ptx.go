package ptx

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Ptx struct{}

func New() *Ptx {
	return &Ptx{}
}

func (p *Ptx) Name() string {
	return "ptx"
}

func (p *Ptx) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("ptx", pflag.ContinueOnError)
	_ = flags.BoolP("auto-reference", "A", false, "output automatically generated references")
	_ = flags.StringP("flag-truncation", "F", "", "use STRING for flag truncation")
	_ = flags.BoolP("gnu-extensions", "G", true, "use GNU extensions")
	_ = flags.StringP("macro-name", "M", "", "use STRING for macro name")
	_ = flags.StringP("format", "O", "roff", "output format (roff or tex)")
	_ = flags.BoolP("right-side-refs", "R", false, "put references at right side")
	_ = flags.StringP("sentence-regexp", "S", "", "use REGEXP for sentence break")
	_ = flags.StringP("word-regexp", "W", "", "use REGEXP for word break")
	_ = flags.StringP("break-file", "b", "", "use FILE for break characters")
	_ = flags.BoolP("ignore-case", "f", false, "ignore case")
	_ = flags.IntP("gap-size", "g", 3, "gap size between columns")
	_ = flags.StringP("ignore-file", "i", "", "use FILE for ignore words")
	_ = flags.StringP("only-file", "o", "", "use FILE for only words")
	_ = flags.BoolP("references", "r", false, "use first word of each line as reference")
	_ = flags.BoolP("typeset-mode", "t", false, "typeset mode")
	_ = flags.IntP("width", "w", 72, "output width")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "ptx: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		targets = []string{"-"}
	}

	// Stub: we just print the lines back for now in a simple way
	for _, target := range targets {
		var reader io.Reader
		if target == "-" {
			reader = env.Stdin
		} else {
			f, err := env.FS.Open(env.Cwd + "/" + target)
			if err != nil {
				fmt.Fprintf(env.Stderr, "ptx: %v\n", err)
				return 1
			}
			defer f.Close()
			reader = f
		}

		data, _ := io.ReadAll(reader)
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if line != "" {
				fmt.Fprintf(env.Stdout, "  %s\n", line)
			}
		}
	}

	return 0
}
