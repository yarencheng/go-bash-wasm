package comm

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Comm struct{}

func New() *Comm {
	return &Comm{}
}

func (c *Comm) Name() string {
	return "comm"
}

func (c *Comm) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("comm", pflag.ContinueOnError)
	suppress1 := flags.BoolP("suppress-1", "1", false, "suppress column 1 (lines unique to FILE1)")
	suppress2 := flags.BoolP("suppress-2", "2", false, "suppress column 2 (lines unique to FILE2)")
	suppress3 := flags.BoolP("suppress-3", "3", false, "suppress column 3 (lines that appear in both files)")
	checkOrder := flags.Bool("check-order", false, "check that the input is correctly sorted")
	noCheckOrder := flags.Bool("nocheck-order", false, "do not check that the input is correctly sorted")
	outputDelimiter := flags.String("output-delimiter", "\t", "separate columns with STR")
	total := flags.Bool("total", false, "output a summary")
	zero := flags.BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "comm: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) < 2 {
		fmt.Fprintf(env.Stderr, "comm: missing operand\n")
		return 1
	}

	f1Path := targets[0]
	if !filepath.IsAbs(f1Path) {
		f1Path = filepath.Join(env.Cwd, f1Path)
	}
	f1, err := env.FS.Open(f1Path)
	if err != nil {
		fmt.Fprintf(env.Stderr, "comm: %v\n", err)
		return 1
	}
	defer f1.Close()

	f2Path := targets[1]
	if !filepath.IsAbs(f2Path) {
		f2Path = filepath.Join(env.Cwd, f2Path)
	}
	f2, err := env.FS.Open(f2Path)
	if err != nil {
		fmt.Fprintf(env.Stderr, "comm: %v\n", err)
		return 1
	}
	defer f2.Close()

	s1 := bufio.NewScanner(f1)
	s2 := bufio.NewScanner(f2)
	if *zero {
		s1.Split(scanNull)
		s2.Split(scanNull)
	}

	v1 := ""
	v2 := ""
	ok1 := s1.Scan()
	if ok1 {
		v1 = s1.Text()
	}
	ok2 := s2.Scan()
	if ok2 {
		v2 = s2.Text()
	}

	var count1, count2, count3 int
	last1, last2 := "", ""

	terminator := "\n"
	if *zero {
		terminator = "\x00"
	}

	for ok1 || ok2 {
		if ok1 && *checkOrder && !*noCheckOrder && last1 != "" && v1 < last1 {
			fmt.Fprintf(env.Stderr, "comm: %s is not sorted\n", f1Path)
			if *checkOrder {
				return 1
			}
		}
		if ok2 && *checkOrder && !*noCheckOrder && last2 != "" && v2 < last2 {
			fmt.Fprintf(env.Stderr, "comm: %s is not sorted\n", f2Path)
			if *checkOrder {
				return 1
			}
		}

		if ok1 && (!ok2 || v1 < v2) {
			count1++
			if !*suppress1 {
				fmt.Fprintf(env.Stdout, "%s%s", v1, terminator)
			}
			last1 = v1
			ok1 = s1.Scan()
			if ok1 {
				v1 = s1.Text()
			}
		} else if ok2 && (!ok1 || v2 < v1) {
			count2++
			if !*suppress2 {
				prefix := ""
				if !*suppress1 {
					prefix = *outputDelimiter
				}
				fmt.Fprintf(env.Stdout, "%s%s%s", prefix, v2, terminator)
			}
			last2 = v2
			ok2 = s2.Scan()
			if ok2 {
				v2 = s2.Text()
			}
		} else {
			// v1 == v2
			count3++
			if !*suppress3 {
				prefix := ""
				if !*suppress1 {
					prefix += *outputDelimiter
				}
				if !*suppress2 {
					prefix += *outputDelimiter
				}
				fmt.Fprintf(env.Stdout, "%s%s%s", prefix, v1, terminator)
			}
			last1, last2 = v1, v2
			ok1 = s1.Scan()
			if ok1 {
				v1 = s1.Text()
			}
			ok2 = s2.Scan()
			if ok2 {
				v2 = s2.Text()
			}
		}
	}

	if *total {
		fmt.Fprintf(env.Stdout, "%d\t%d\t%d\ttotal%s", count1, count2, count3, terminator)
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

