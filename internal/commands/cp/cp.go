package cp

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Cp struct{}

func New() *Cp {
	return &Cp{}
}

func (c *Cp) Name() string {
	return "cp"
}

func (c *Cp) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("cp", pflag.ContinueOnError)
	recursive := flags.BoolP("recursive", "r", false, "copy directories recursively")
	recursiveUpper := flags.BoolP("recursive-upper", "R", false, "identical to -r")
	verbose := flags.BoolP("verbose", "v", false, "explain what is being done")
	_ = flags.BoolP("interactive", "i", false, "prompt before overwrite (ignored)")
	_ = flags.BoolP("force", "f", false, "if an existing destination file cannot be opened, remove it and try again (ignored)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "cp: %v\n", err)
		return 1
	}

	posArgs := flags.Args()
	if len(posArgs) < 2 {
		fmt.Fprintf(env.Stderr, "cp: missing file operand\n")
		return 1
	}

	sources := posArgs[:len(posArgs)-1]
	dest := posArgs[len(posArgs)-1]

	doRecursive := *recursive || *recursiveUpper
	exitCode := 0

	destFullPath := dest
	if !filepath.IsAbs(dest) {
		destFullPath = filepath.Join(env.Cwd, dest)
	}

	destInfo, destErr := env.FS.Stat(destFullPath)
	isDestDir := destErr == nil && destInfo.IsDir()

	if len(sources) > 1 && !isDestDir {
		fmt.Fprintf(env.Stderr, "cp: target '%s' is not a directory\n", dest)
		return 1
	}

	for _, src := range sources {
		srcFullPath := src
		if !filepath.IsAbs(src) {
			srcFullPath = filepath.Join(env.Cwd, src)
		}

		srcInfo, err := env.FS.Stat(srcFullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "cp: cannot stat '%s': %v\n", src, err)
			exitCode = 1
			continue
		}

		finalDest := destFullPath
		if isDestDir {
			finalDest = filepath.Join(destFullPath, filepath.Base(srcFullPath))
		}

		if srcInfo.IsDir() {
			if !doRecursive {
				fmt.Fprintf(env.Stderr, "cp: -r not specified; omitting directory '%s'\n", src)
				exitCode = 1
				continue
			}
			err = c.copyDir(env, srcFullPath, finalDest, *verbose)
		} else {
			err = c.copyFile(env, srcFullPath, finalDest, *verbose)
		}

		if err != nil {
			fmt.Fprintf(env.Stderr, "cp: %v\n", err)
			exitCode = 1
		}
	}

	return exitCode
}

func (c *Cp) copyFile(env *commands.Environment, src, dest string, verbose bool) error {
	in, err := env.FS.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := env.FS.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	if verbose {
		fmt.Fprintf(env.Stdout, "'%s' -> '%s'\n", src, dest)
	}
	return nil
}

func (c *Cp) copyDir(env *commands.Environment, src, dest string, verbose bool) error {
	srcInfo, err := env.FS.Stat(src)
	if err != nil {
		return err
	}

	err = env.FS.MkdirAll(dest, srcInfo.Mode())
	if err != nil {
		return err
	}

	entries, err := afero.ReadDir(env.FS, src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			err = c.copyDir(env, srcPath, destPath, verbose)
		} else {
			err = c.copyFile(env, srcPath, destPath, verbose)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
