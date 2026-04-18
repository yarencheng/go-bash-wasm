package sortcmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strings"

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
	unique := flags.BoolP("unique", "u", false, "output only the first of an equal run")
	ignoreCase := flags.BoolP("ignore-case", "f", false, "fold lower case to upper case characters")
	numeric := flags.BoolP("numeric-sort", "n", false, "compare according to string numerical value")
	check := flags.BoolP("check", "c", false, "check for sorted input; do not sort")
	outputFile := flags.StringP("output", "o", "", "write result to FILE instead of standard output")

	_ = flags.StringP("key", "k", "", "sort via a key (ignored)")
	_ = flags.BoolP("merge", "m", false, "merge already sorted files (ignored)")
	_ = flags.BoolP("stable", "s", false, "stabilize sort (ignored)")
	_ = flags.StringP("buffer-size", "S", "", "use SIZE for main memory buffer (ignored)")
	_ = flags.StringP("field-separator", "t", "", "use SEP instead of blank to blank transition (ignored)")
	_ = flags.StringP("temporary-directory", "T", "", "use DIR for temporaries (ignored)")
	_ = flags.Int("parallel", 1, "change the number of sorts run concurrently to N (ignored)")
	_ = flags.Bool("debug", false, "annotate the part of the line used to sort (ignored)")
	_ = flags.String("files0-from", "", "read input from the files specified by NUL-terminated names in file F (ignored)")
	_ = flags.BoolP("ignore-leading-blanks", "b", false, "ignore leading blanks (ignored)")
	_ = flags.BoolP("dictionary-order", "d", false, "consider only blanks and alphanumeric characters (ignored)")
	_ = flags.BoolP("general-numeric-sort", "g", false, "compare according to general numerical value (ignored)")
	_ = flags.BoolP("human-numeric-sort", "h", false, "compare human readable numbers (e.g., 2K 1G) (ignored)")
	_ = flags.BoolP("ignore-nonprinting", "i", false, "consider only printable characters (ignored)")
	_ = flags.BoolP("month-sort", "M", false, "compare (unknown) < 'JAN' < ... < 'DEC' (ignored)")
	_ = flags.BoolP("random-sort", "R", false, "shuffle, but group identical keys (ignored)")
	_ = flags.BoolP("version-sort", "V", false, "natural sort of (version) numbers (ignored)")
	_ = flags.BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "sort: %v\n", err)
		}
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: sort [OPTION]... [FILE]...\n")
		fmt.Fprintf(env.Stdout, "Write sorted concatenation of all FILE(s) to standard output.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "sort")
		return 0
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
				if env.Stderr != nil {
					fmt.Fprintf(env.Stderr, "sort: %v\n", err)
				}
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

	compare := func(i, j int) bool {
		s1, s2 := lines[i], lines[j]
		if *ignoreCase {
			s1 = strings.ToLower(s1)
			s2 = strings.ToLower(s2)
		}

		res := 0
		if *numeric {
			var n1, n2 float64
			fmt.Sscanf(s1, "%f", &n1)
			fmt.Sscanf(s2, "%f", &n2)
			if n1 < n2 {
				res = -1
			} else if n1 > n2 {
				res = 1
			} else {
				res = strings.Compare(s1, s2)
			}
		} else {
			res = strings.Compare(s1, s2)
		}

		if *reverse {
			return res > 0
		}
		return res < 0
	}

	if *check {
		for i := 1; i < len(lines); i++ {
			if compare(i, i-1) { // if lines[i] < lines[i-1]
				if env.Stderr != nil {
					fmt.Fprintf(env.Stderr, "sort: disorder: %s\n", lines[i])
				}
				return 1
			}
		}
		return 0
	}

	sort.SliceStable(lines, compare)

	if *unique {
		if len(lines) > 0 {
			uniqueLines := []string{lines[0]}
			for i := 1; i < len(lines); i++ {
				// Compare function can be used here too if we want "equal" check
				if lines[i] != lines[i-1] {
					uniqueLines = append(uniqueLines, lines[i])
				}
			}
			lines = uniqueLines
		}
	}

	var out io.Writer = env.Stdout
	if *outputFile != "" {
		fullPath := *outputFile
		if !filepath.IsAbs(fullPath) {
			fullPath = filepath.Join(env.Cwd, fullPath)
		}
		f, err := env.FS.Create(fullPath)
		if err != nil {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "sort: %v\n", err)
			}
			return 1
		}
		defer f.Close()
		out = f
	}

	for _, line := range lines {
		fmt.Fprintln(out, line)
	}

	return 0
}
