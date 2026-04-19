package sum

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/cksum"
)

type Sum struct {
	name string
}

func New(name string) *Sum {
	return &Sum{name: name}
}

func (s *Sum) Name() string {
	return s.name
}

func (s *Sum) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet(s.name, pflag.ContinueOnError)
	check := flags.BoolP("check", "c", false, "read checksums from the FILEs and check them")
	zero := flags.BoolP("zero", "z", false, "end each output line with NUL, not newline")
	quiet := flags.Bool("quiet", false, "don't print OK for each successfully verified file")
	status := flags.Bool("status", false, "don't output anything, status code shows success")
	strict := flags.Bool("strict", false, "exit non-zero for improperly formatted checksum lines")
	warn := flags.BoolP("warn", "w", false, "warn about improperly formatted checksum lines")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "%s: %v\n", s.name, err)
		return 1
	}

	algo := strings.TrimSuffix(s.name, "sum")
	opts := cksum.CksumOptions{
		Algorithm: algo,
		Check:     *check,
		Zero:      *zero,
		Quiet:     *quiet,
		Status:    *status,
		Strict:    *strict,
		Warn:      *warn,
	}

	c := cksum.New()
	targets := flags.Args()
	if len(targets) == 0 {
		return c.Process(env, env.Stdin, "", opts)
	}

	exitCode := 0
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		f, err := env.FS.Open(fullPath)
		if err != nil {
			if !opts.Status {
				fmt.Fprintf(env.Stderr, "%s: %s: %v\n", s.name, target, err)
			}
			exitCode = 1
			continue
		}

		if res := c.Process(env, f, target, opts); res != 0 {
			exitCode = res
		}
		f.Close()
	}

	return exitCode
}
