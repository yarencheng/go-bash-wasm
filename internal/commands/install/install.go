package install

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Install struct{}

func New() *Install {
	return &Install{}
}

func (i *Install) Name() string {
	return "install"
}

func (i *Install) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("install", pflag.ContinueOnError)
	directory := flags.BoolP("directory", "d", false, "treat all arguments as directory names")
	mode := flags.StringP("mode", "m", "", "set permission mode (as in chmod)")
	owner := flags.StringP("owner", "o", "", "set ownership (super-user only)")
	group := flags.StringP("group", "g", "", "set group ownership")
	verbose := flags.BoolP("verbose", "v", false, "explain what is being done")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "install: %v\n", err)
		return 1
	}

	// Suppress unused variable errors for flags we don't handle yet
	_ = mode
	_ = owner
	_ = group

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "install: missing file operand\n")
		return 1
	}

	if *directory {
		for _, target := range targets {
			fullPath := target
			if !filepath.IsAbs(target) {
				fullPath = filepath.Join(env.Cwd, target)
			}
			err := env.FS.MkdirAll(fullPath, 0755)
			if err != nil {
				fmt.Fprintf(env.Stderr, "install: cannot create directory '%s': %v\n", target, err)
				return 1
			}
			if *verbose {
				fmt.Fprintf(env.Stdout, "install: creating directory '%s'\n", target)
			}
		}
		return 0
	}

	if len(targets) < 2 {
		fmt.Fprintf(env.Stderr, "install: missing destination file operand after '%s'\n", targets[0])
		return 1
	}

	// Basic implementation: copy files from targets[0...n-1] to targets[n-1]
	dest := targets[len(targets)-1]
	sources := targets[:len(targets)-1]

	fullDestPath := dest
	if !filepath.IsAbs(dest) {
		fullDestPath = filepath.Join(env.Cwd, dest)
	}

	destInfo, err := env.FS.Stat(fullDestPath)
	destIsDir := err == nil && destInfo.IsDir()

	if len(sources) > 1 && !destIsDir {
		fmt.Fprintf(env.Stderr, "install: target '%s' is not a directory\n", dest)
		return 1
	}

	for _, src := range sources {
		fullSrcPath := src
		if !filepath.IsAbs(src) {
			fullSrcPath = filepath.Join(env.Cwd, src)
		}

		srcFile, err := env.FS.Open(fullSrcPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "install: cannot stat '%s': %v\n", src, err)
			return 1
		}
		defer srcFile.Close()

		finalDest := fullDestPath
		if destIsDir {
			finalDest = filepath.Join(fullDestPath, filepath.Base(src))
		}

		dstFile, err := env.FS.OpenFile(finalDest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprintf(env.Stderr, "install: cannot open '%s' for writing: %v\n", finalDest, err)
			return 1
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			fmt.Fprintf(env.Stderr, "install: error copying '%s' to '%s': %v\n", src, finalDest, err)
			return 1
		}

		if *verbose {
			fmt.Fprintf(env.Stdout, "'%s' -> '%s'\n", src, finalDest)
		}
	}

	return 0
}
