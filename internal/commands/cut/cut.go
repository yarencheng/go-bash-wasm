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

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "cut: %v\n", err)
		return 1
	}

	if *fields == "" && *bytes == "" {
		fmt.Fprintf(env.Stderr, "cut: you must specify a list of bytes, characters, or fields\n")
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
				fmt.Fprintf(env.Stderr, "cut: %v\n", err)
				continue
			}
			inputs = append(inputs, f)
		}
	}

	for _, input := range inputs {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			if *bytes != "" {
				c.cutBytes(env.Stdout, line, *bytes)
			} else {
				c.cutFields(env.Stdout, line, *fields, *delimiter)
			}
		}
		if input != env.Stdin {
			input.Close()
		}
	}

	return 0
}

func (c *Cut) cutBytes(w io.Writer, line, ranges string) {
	indices := c.parseRanges(ranges, len(line))
	var result strings.Builder
	for _, i := range indices {
		if i < len(line) {
			result.WriteByte(line[i])
		}
	}
	fmt.Fprintln(w, result.String())
}

func (c *Cut) cutFields(w io.Writer, line, ranges, delim string) {
	parts := strings.Split(line, delim)
	indices := c.parseRanges(ranges, len(parts))
	
	var result []string
	for _, i := range indices {
		if i < len(parts) {
			result = append(result, parts[i])
		}
	}
	fmt.Fprintln(w, strings.Join(result, delim))
}

func (c *Cut) parseRanges(ranges string, max int) []int {
	// Simple range parser for "1,2,5-8"
	var result []int
	parts := strings.Split(ranges, ",")
	for _, p := range parts {
		if strings.Contains(p, "-") {
			r := strings.Split(p, "-")
			start, _ := strconv.Atoi(r[0])
			end := max
			if r[1] != "" {
				end, _ = strconv.Atoi(r[1])
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
	return final
}
