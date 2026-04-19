package pathchk

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Pathchk struct{}

func New() *Pathchk {
	return &Pathchk{}
}

func (p *Pathchk) Name() string {
	return "pathchk"
}

const (
	posixPathMax = 4096
	posixNameMax = 255
	posixChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789._-/"
)

func (p *Pathchk) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("pathchk", pflag.ContinueOnError)
	posix := flags.BoolP("posix", "p", false, "check for POSIX portability")
	noEmpty := flags.BoolP("no-empty", "P", false, "check for empty paths")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "pathchk: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		return 0
	}

	exitCode := 0
	for _, path := range targets {
		if path == "" {
			if *noEmpty {
				fmt.Fprintf(env.Stderr, "pathchk: empty path\n")
				exitCode = 1
			}
			continue
		}

		if *posix {
			if len(path) > posixPathMax {
				fmt.Fprintf(env.Stderr, "pathchk: path too long\n")
				exitCode = 1
			}
			for _, c := range path {
				if !strings.ContainsRune(posixChars, c) {
					fmt.Fprintf(env.Stderr, "pathchk: invalid character '%c' in path\n", c)
					exitCode = 1
					break
				}
			}
			components := strings.Split(path, string(filepath.Separator))
			for _, comp := range components {
				if len(comp) > posixNameMax {
					fmt.Fprintf(env.Stderr, "pathchk: path component too long\n")
					exitCode = 1
					break
				}
			}
		}
	}

	return exitCode
}
