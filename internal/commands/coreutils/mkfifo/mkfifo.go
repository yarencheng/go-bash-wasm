package mkfifo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Mkfifo struct{}

func New() *Mkfifo {
	return &Mkfifo{}
}

func (m *Mkfifo) Name() string {
	return "mkfifo"
}

func (m *Mkfifo) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("mkfifo", pflag.ContinueOnError)
	modeStr := flags.StringP("mode", "m", "0666", "set file mode (as in chmod), not a=rw - umask")
	_ = flags.StringP("context", "Z", "", "set SELinux security context (ignored)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "mkfifo: %v\n", err)
		return 1
	}

	mode := uint64(0666)
	if *modeStr != "" {
		parsed, err := strconv.ParseUint(*modeStr, 8, 32)
		if err != nil {
			fmt.Fprintf(env.Stderr, "mkfifo: invalid mode '%s'\n", *modeStr)
			return 1
		}
		mode = parsed
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "mkfifo: missing operand\n")
		return 1
	}

	exitCode := 0
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		// In MemMapFs, we can't really create a FIFO.
		// We'll just create a file with the FIFO bit if possible,
		// but afero.MemMapFs doesn't really care about the FIFO bit for operations.
		// We'll just create a normal file as a stub.
		err := env.FS.MkdirAll(filepath.Dir(fullPath), 0755)
		if err != nil {
			fmt.Fprintf(env.Stderr, "mkfifo: cannot create directory for '%s': %v\n", target, err)
			exitCode = 1
			continue
		}

		f, err := env.FS.OpenFile(fullPath, os.O_CREATE|os.O_EXCL, os.FileMode(mode)|os.ModeNamedPipe)
		if err != nil {
			fmt.Fprintf(env.Stderr, "mkfifo: cannot create fifo '%s': %v\n", target, err)
			exitCode = 1
			continue
		}
		f.Close()
	}

	return exitCode
}
