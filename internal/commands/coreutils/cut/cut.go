package cut

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Cut struct{}

func New() *Cut {
	return &Cut{}
}

func (c *Cut) Name() string {
	return "cut"
}

func (c *Cut) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("cut", pflag.ContinueOnError)
	fields := flags.StringP("fields", "f", "", "select only these fields")
	delimiter := flags.StringP("delimiter", "d", "\t", "use DELIM instead of TAB for field delimiter")
	bytes := flags.StringP("bytes", "b", "", "select only these bytes")
	complement := flags.Bool("complement", false, "complement the set of selected bytes, characters or fields")
	outputDelimiter := flags.String("output-delimiter", "", "use STRING as the output delimiter")
	zero := flags.BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")
	_ = flags.BoolP("ignored-n", "n", false, "ignored")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "cut: %v\n", err)
		return 1
	}

	if *fields == "" && *bytes == "" {
		fmt.Fprintf(env.Stderr, "cut: you must specify a list of bytes, characters, or fields\n")
		return 1
	}

	var outDelim string
	if *outputDelimiter != "" {
		outDelim = *outputDelimiter
	} else {
		outDelim = *delimiter
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
				fmt.Fprintf(env.Stderr, "cut: %v\n", err)
				continue
			}
			inputs = append(inputs, f)
		}
	}

	terminator := "\n"
	if *zero {
		terminator = "\x00"
	}

	for _, input := range inputs {
		scanner := bufio.NewScanner(input)
		if *zero {
			scanner.Split(scanNull)
		}
		for scanner.Scan() {
			line := scanner.Text()
			if *bytes != "" {
				c.cutBytes(env.Stdout, line, *bytes, *complement, terminator)
			} else {
				c.cutFields(env.Stdout, line, *fields, *delimiter, outDelim, *complement, terminator)
			}
		}
		if input != env.Stdin {
			input.Close()
		}
	}

	return 0
}

func (c *Cut) cutBytes(w io.Writer, line, ranges string, complement bool, terminator string) {
	indices := c.parseRanges(ranges, len(line), complement)
	var result strings.Builder
	for _, i := range indices {
		if i < len(line) {
			result.WriteByte(line[i])
		}
	}
	fmt.Fprintf(w, "%s%s", result.String(), terminator)
}

func (c *Cut) cutFields(w io.Writer, line, ranges, delim, outDelim string, complement bool, terminator string) {
	parts := strings.Split(line, delim)
	indices := c.parseRanges(ranges, len(parts), complement)

	var result []string
	for _, i := range indices {
		if i < len(parts) {
			result = append(result, parts[i])
		}
	}
	fmt.Fprintf(w, "%s%s", strings.Join(result, outDelim), terminator)
}

func scanNull(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.IndexByte(string(data), '\x00'); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func (c *Cut) parseRanges(ranges string, max int, complement bool) []int {
	// Simple range parser for "1,2,5-8"
	var result []int
	parts := strings.Split(ranges, ",")
	for _, p := range parts {
		if strings.Contains(p, "-") {
			r := strings.Split(p, "-")
			startStr := r[0]
			endStr := r[1]

			start := 1
			if startStr != "" {
				start, _ = strconv.Atoi(startStr)
			}

			end := max
			if endStr != "" {
				end, _ = strconv.Atoi(endStr)
			}

			if start < 1 {
				start = 1
			}
			for i := start; i <= end && i <= max; i++ {
				result = append(result, i-1)
			}
		} else {
			i, _ := strconv.Atoi(p)
			if i > 0 && i <= max {
				result = append(result, i-1)
			}
		}
	}

	// Deduplicate and sort
	unique := make(map[int]bool)
	var final []int
	for _, i := range result {
		if !unique[i] {
			unique[i] = true
			final = append(final, i)
		}
	}
	sort.Ints(final)

	if complement {
		var complemented []int
		selectedMap := make(map[int]bool)
		for _, i := range final {
			selectedMap[i] = true
		}
		for i := 0; i < max; i++ {
			if !selectedMap[i] {
				complemented = append(complemented, i)
			}
		}
		return complemented
	}

	return final
}
