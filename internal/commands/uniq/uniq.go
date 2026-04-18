package uniq

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Uniq struct{}

func New() *Uniq {
	return &Uniq{}
}

func (u *Uniq) Name() string {
	return "uniq"
}

func (u *Uniq) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("uniq", pflag.ContinueOnError)
	count := flags.BoolP("count", "c", false, "prefix lines by the number of occurrences")
	repeated := flags.BoolP("repeated", "d", false, "only print duplicate lines, one for each group")
	unique := flags.BoolP("unique", "u", false, "only print unique lines")
	ignoreCase := flags.BoolP("ignore-case", "i", false, "ignore differences in case when comparing")
	zero := flags.BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "uniq: %v\n", err)
		return 1
	}

	var input io.ReadCloser
	targets := flags.Args()

	if len(targets) == 0 {
		input = env.Stdin
	} else {
		fullPath := targets[0]
		if !filepath.IsAbs(fullPath) {
			fullPath = filepath.Join(env.Cwd, fullPath)
		}
		f, err := env.FS.Open(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "uniq: %v\n", err)
			return 1
		}
		input = f
	}
	defer func() {
		if input != env.Stdin {
			input.Close()
		}
	}()

	scanner := bufio.NewScanner(input)
	if *zero {
		scanner.Split(scanNull)
	}

	var prevLine string
	var currentCount int
	first := true

	compare := func(a, b string) bool {
		if *ignoreCase {
			return strings.EqualFold(a, b)
		}
		return a == b
	}

	terminator := "\n"
	if *zero {
		terminator = "\x00"
	}

	for scanner.Scan() {
		line := scanner.Text()
		if first {
			prevLine = line
			currentCount = 1
			first = false
			continue
		}

		if compare(line, prevLine) {
			currentCount++
		} else {
			u.outputLine(env.Stdout, prevLine, currentCount, *count, *repeated, *unique, terminator)
			prevLine = line
			currentCount = 1
		}
	}

	if !first {
		u.outputLine(env.Stdout, prevLine, currentCount, *count, *repeated, *unique, terminator)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(env.Stderr, "uniq: %v\n", err)
		return 1
	}

	return 0
}


func (u *Uniq) outputLine(w io.Writer, line string, count int, showCount, onlyRepeated, onlyUnique bool, terminator string) {
	if onlyRepeated && count <= 1 {
		return
	}
	if onlyUnique && count > 1 {
		return
	}

	if showCount {
		fmt.Fprintf(w, "%7d %s%s", count, line, terminator)
	} else {
		fmt.Fprintf(w, "%s%s", line, terminator)
	}
}

