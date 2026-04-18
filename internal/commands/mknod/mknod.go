package mknod

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Mknod struct{}

func New() *Mknod {
	return &Mknod{}
}

func (m *Mknod) Name() string {
	return "mknod"
}

func (m *Mknod) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("mknod", pflag.ContinueOnError)
	modeStr := flags.StringP("mode", "m", "0666", "set file mode (as in chmod), not a=rw - umask")
	_ = flags.StringP("context", "Z", "", "set SELinux security context (ignored)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "mknod: %v\n", err)
		return 1
	}

	mode := uint64(0666)
	if *modeStr != "" {
		parsed, err := strconv.ParseUint(*modeStr, 8, 32)
		if err != nil {
			fmt.Fprintf(env.Stderr, "mknod: invalid mode '%s'\n", *modeStr)
			return 1
		}
		mode = parsed
	}

	targets := flags.Args()
	if len(targets) < 2 {
		fmt.Fprintf(env.Stderr, "mknod: missing operand\n")
		return 1
	}

	name := targets[0]
	typ := targets[1]

	fullPath := name
	if !filepath.IsAbs(name) {
		fullPath = filepath.Join(env.Cwd, name)
	}

	var fileMode os.FileMode = os.FileMode(mode)
	switch typ {
	case "p":
		fileMode |= os.ModeNamedPipe
	case "b":
		fileMode |= os.ModeDevice
	case "c", "u":
		fileMode |= os.ModeDevice | os.ModeCharDevice
	default:
		fmt.Fprintf(env.Stderr, "mknod: invalid device type '%s'\n", typ)
		return 1
	}

	// For b, c, u we need major and minor
	if typ != "p" {
		if len(targets) < 4 {
			fmt.Fprintf(env.Stderr, "mknod: missing major or minor device number\n")
			return 1
		}
		// We ignore major/minor in simulation
	}

	err := env.FS.MkdirAll(filepath.Dir(fullPath), 0755)
	if err != nil {
		fmt.Fprintf(env.Stderr, "mknod: %v\n", err)
		return 1
	}

	f, err := env.FS.OpenFile(fullPath, os.O_CREATE|os.O_EXCL, fileMode)
	if err != nil {
		fmt.Fprintf(env.Stderr, "mknod: %v\n", err)
		return 1
	}
	f.Close()

	return 0
}
