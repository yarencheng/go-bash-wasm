package mkdir

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Mkdir struct{}

func New() *Mkdir {
	return &Mkdir{}
}

func (m *Mkdir) Name() string {
	return "mkdir"
}

func (m *Mkdir) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("mkdir", pflag.ContinueOnError)
	parents := flags.BoolP("parents", "p", false, "no error if existing, make parent directories as needed")
	verbose := flags.BoolP("verbose", "v", false, "print a message for each created directory")
	modeStr := flags.StringP("mode", "m", "0755", "set file mode (as in chmod), not a=rwx - umask")
	_ = flags.StringP("context", "Z", "", "set SELinux security context of each created directory to CTX (ignored)")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "mkdir: %v\n", err)
		}
		return 1
	}

	mode := uint64(0755)
	if *modeStr != "" {
		parsed, err := strconv.ParseUint(*modeStr, 8, 32)
		if err != nil {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "mkdir: invalid mode '%s'\n", *modeStr)
			}
			return 1
		}
		mode = parsed
	}

	targets := flags.Args()
	if len(targets) == 0 {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "mkdir: missing operand\n")
		}
		return 1
	}

	exitCode := 0
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		var err error
		if *parents {
			err = env.FS.MkdirAll(fullPath, os.FileMode(mode))
		} else {
			err = env.FS.Mkdir(fullPath, os.FileMode(mode))
		}

		if err != nil {
			if strings.Contains(err.Error(), "already exists") && *parents {
				// According to GNU mkdir -p, this is not an error
			} else {
				if env.Stderr != nil {
					fmt.Fprintf(env.Stderr, "mkdir: cannot create directory '%s': %v\n", target, err)
				}
				exitCode = 1
				continue
			}
		}

		if *verbose {
			fmt.Fprintf(env.Stdout, "mkdir: created directory '%s'\n", target)
		}
	}

	return exitCode
}
