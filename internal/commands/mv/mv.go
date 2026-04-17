package mv

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Mv struct{}

func New() *Mv {
	return &Mv{}
}

func (m *Mv) Name() string {
	return "mv"
}

func (m *Mv) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("mv", pflag.ContinueOnError)
	verbose := flags.BoolP("verbose", "v", false, "explain what is being done")
	_ = flags.BoolP("force", "f", false, "do not prompt before overwriting (ignored)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "mv: %v\n", err)
		return 1
	}

	posArgs := flags.Args()
	if len(posArgs) < 2 {
		fmt.Fprintf(env.Stderr, "mv: missing file operand\n")
		return 1
	}

	sources := posArgs[:len(posArgs)-1]
	dest := posArgs[len(posArgs)-1]

	destFullPath := dest
	if !filepath.IsAbs(dest) {
		destFullPath = filepath.Join(env.Cwd, dest)
	}

	destInfo, destErr := env.FS.Stat(destFullPath)
	isDestDir := destErr == nil && destInfo.IsDir()

	if len(sources) > 1 && !isDestDir {
		fmt.Fprintf(env.Stderr, "mv: target '%s' is not a directory\n", dest)
		return 1
	}

	exitCode := 0
	for _, src := range sources {
		srcFullPath := src
		if !filepath.IsAbs(src) {
			srcFullPath = filepath.Join(env.Cwd, src)
		}

		finalDest := destFullPath
		if isDestDir {
			finalDest = filepath.Join(destFullPath, filepath.Base(srcFullPath))
		}

		err := env.FS.Rename(srcFullPath, finalDest)
		if err != nil {
			fmt.Fprintf(env.Stderr, "mv: cannot move '%s' to '%s': %v\n", src, finalDest, err)
			exitCode = 1
			continue
		}

		if *verbose {
			fmt.Fprintf(env.Stdout, "renamed '%s' -> '%s'\n", src, finalDest)
		}
	}

	return exitCode
}
