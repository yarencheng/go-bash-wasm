package sortcmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"sort"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Sort struct{}

func New() *Sort {
	return &Sort{}
}

func (s *Sort) Name() string {
	return "sort"
}

func (s *Sort) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("sort", pflag.ContinueOnError)
	reverse := flags.BoolP("reverse", "r", false, "reverse the result of comparisons")
	unique := flags.BoolP("unique", "u", false, "with -c, check for strict ordering; otherwise, output only the first of an equal run")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "sort: %v\n", err)
		return 1
	}

	var inputs []io.ReadCloser
	targets := flags.Args()

	if len(targets) == 0 {
		inputs = append(inputs, env.Stdin)
	} else {
		for _, target := range targets {
			fullPath := target
			if !filepath.IsAbs(target) {
				fullPath = filepath.Join(env.Cwd, target)
			}
			f, err := env.FS.Open(fullPath)
			if err != nil {
				fmt.Fprintf(env.Stderr, "sort: %v\n", err)
				continue
			}
			inputs = append(inputs, f)
		}
	}

	var lines []string
	for _, input := range inputs {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if input != env.Stdin {
			input.Close()
		}
	}

	if *unique {
		lineMap := make(map[string]bool)
		var uniqueLines []string
		for _, line := range lines {
			if !lineMap[line] {
				lineMap[line] = true
				uniqueLines = append(uniqueLines, line)
			}
		}
		lines = uniqueLines
	}

	if *reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(lines)))
	} else {
		sort.Strings(lines)
	}

	for _, line := range lines {
		fmt.Fprintln(env.Stdout, line)
	}

	return 0
}
