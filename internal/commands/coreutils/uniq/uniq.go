package uniq

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"

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

	_ = flags.BoolP("all-repeated", "D", false, "print all duplicate lines (ignored)")
	_ = flags.IntP("skip-fields", "f", 0, "avoid comparing the first N fields (ignored)")
	_ = flags.IntP("skip-chars", "s", 0, "avoid comparing the first N characters (ignored)")
	_ = flags.IntP("check-chars", "w", 0, "compare no more than N characters in lines (ignored)")
	_ = flags.String("group", "", "show all items, separating groups with an empty line (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "uniq: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: uniq [OPTION]... [INPUT [OUTPUT]]\n")
		fmt.Fprintf(env.Stdout, "Filter adjacent matching lines from INPUT (or standard input), writing to OUTPUT (or standard output).\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "uniq")
		return 0
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

func scanNull(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\x00'); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
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
