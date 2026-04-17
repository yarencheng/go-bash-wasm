package comm

import (
	"bufio"
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

	for ok1 || ok2 {
		if ok1 && (!ok2 || v1 < v2) {
			if !*suppress1 {
				fmt.Fprintln(env.Stdout, v1)
			}
			ok1 = s1.Scan()
			if ok1 {
				v1 = s1.Text()
			}
		} else if ok2 && (!ok1 || v2 < v1) {
			if !*suppress2 {
				prefix := ""
				if !*suppress1 {
					prefix = "\t"
				}
				fmt.Fprintf(env.Stdout, "%s%s\n", prefix, v2)
			}
			ok2 = s2.Scan()
			if ok2 {
				v2 = s2.Text()
			}
		} else {
			// v1 == v2
			if !*suppress3 {
				prefix := ""
				if !*suppress1 {
					prefix += "\t"
				}
				if !*suppress2 {
					prefix += "\t"
				}
				fmt.Fprintf(env.Stdout, "%s%s\n", prefix, v1)
			}
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

	return 0
}
