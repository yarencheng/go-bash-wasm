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
	targetDir := flags.StringP("target-directory", "t", "", "copy all SOURCE arguments into DIRECTORY")
	noTargetDir := flags.BoolP("no-target-directory", "T", false, "treat DEST as a normal file")
	verbose := flags.BoolP("verbose", "v", false, "explain what is being done")
	noClobber := flags.BoolP("no-clobber", "n", false, "do not overwrite an existing file")
	update := flags.BoolP("update", "u", false, "copy only when the SOURCE file is newer than the destination file or when the destination file is missing")
	archive := flags.BoolP("archive", "a", false, "copy directories recursively and preserve all attributes")
	preserve := flags.BoolP("preserve", "p", false, "preserve timestamps and mode")
	dereference := flags.BoolP("dereference", "L", false, "always follow symbolic links in SOURCE")
	noDereference := flags.BoolP("no-dereference", "P", false, "never follow symbolic links in SOURCE")
	_ = flags.BoolP("interactive", "i", false, "prompt before overwrite (ignored)")
	_ = flags.BoolP("force", "f", false, "if an existing destination file cannot be opened, remove it and try again (ignored)")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "cp: %v\n", err)
		}
		return 1
	}

	posArgs := flags.Args()

	var sources []string
	var dest string

	if *targetDir != "" {
		sources = posArgs
		dest = *targetDir
	} else {
		if len(posArgs) < 1 {
			fmt.Fprintf(env.Stderr, "cp: missing file operand\n")
			return 1
		}
		if len(posArgs) == 1 {
			fmt.Fprintf(env.Stderr, "cp: missing destination file operand after '%s'\n", posArgs[0])
			return 1
		}
		sources = posArgs[:len(posArgs)-1]
		dest = posArgs[len(posArgs)-1]
	}

	doRecursive := *recursive || *recursiveUpper || *archive
	doPreserve := *preserve || *archive
	// Simplistic symlink handling: default to following links unless -P or -a is set
	followLinks := *dereference || (!*noDereference && !*archive)

	exitCode := 0

	destFullPath := dest
	if !filepath.IsAbs(dest) {
		destFullPath = filepath.Join(env.Cwd, dest)
	}

	destInfo, destErr := env.FS.Stat(destFullPath)
	isDestDir := destErr == nil && destInfo.IsDir()

	if *noTargetDir && isDestDir && len(sources) > 0 {
		// If -T is specified, dest cannot be a directory unless we're copying ONE thing into it?
		// No, -T means treat it as a file. If it already IS a directory, it's usually an error for -T.
	}

	if len(sources) > 1 && !isDestDir && *targetDir == "" {
		fmt.Fprintf(env.Stderr, "cp: target '%s' is not a directory\n", dest)
		return 1
	}

	for _, src := range sources {
		srcFullPath := src
		if !filepath.IsAbs(src) {
			srcFullPath = filepath.Join(env.Cwd, src)
		}

		var srcInfo os.FileInfo
		var err error
		if followLinks {
			srcInfo, err = env.FS.Stat(srcFullPath)
		} else {
			srcInfo, err = lstat(env.FS, srcFullPath)
		}

		if err != nil {
			fmt.Fprintf(env.Stderr, "cp: cannot stat '%s': %v\n", src, err)
			exitCode = 1
			continue
		}

		finalDest := destFullPath
		if isDestDir && !*noTargetDir {
			finalDest = filepath.Join(destFullPath, filepath.Base(srcFullPath))
		}

		if srcInfo.IsDir() {
			if !doRecursive {
				fmt.Fprintf(env.Stderr, "cp: -r not specified; omitting directory '%s'\n", src)
				exitCode = 1
				continue
			}
			err = c.copyDir(env, srcFullPath, finalDest, *verbose, *noClobber, *update, doPreserve, followLinks)
		} else {
			err = c.copyFile(env, srcFullPath, finalDest, *verbose, *noClobber, *update, doPreserve, followLinks)
		}

		if err != nil {
			fmt.Fprintf(env.Stderr, "cp: %v\n", err)
			exitCode = 1
		}
	}

	return exitCode
}

func (c *Cp) copyFile(env *commands.Environment, src, dest string, verbose, noClobber, update, preserve, followLinks bool) error {
	var srcInfo os.FileInfo
	var err error
	if followLinks {
		srcInfo, err = env.FS.Stat(src)
	} else {
		srcInfo, err = lstat(env.FS, src)
	}
	if err != nil {
		return err
	}

	destInfo, err := env.FS.Stat(dest)
	if err == nil {
		if noClobber {
			return nil
		}
		if update && !srcInfo.ModTime().After(destInfo.ModTime()) {
			return nil
		}
	}

	// Handle symlink copy if not following
	if !followLinks && (srcInfo.Mode()&os.ModeSymlink != 0) {
		// Afero limited symlink support
		// For now, we skip or try to readlink
		if linker, ok := env.FS.(afero.Symlinker); ok {
			target, err := linker.ReadlinkIfPossible(src)
			if err == nil {
				return linker.SymlinkIfPossible(target, dest)
			}
		}
	}

	in, err := env.FS.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := env.FS.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, srcInfo.Mode())
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	if preserve {
		_ = env.FS.Chtimes(dest, srcInfo.ModTime(), srcInfo.ModTime())
		_ = env.FS.Chmod(dest, srcInfo.Mode())
	}

	if verbose {
		fmt.Fprintf(env.Stdout, "'%s' -> '%s'\n", src, dest)
	}
	return nil
}

func (c *Cp) copyDir(env *commands.Environment, src, dest string, verbose, noClobber, update, preserve, followLinks bool) error {
	var srcInfo os.FileInfo
	var err error
	if followLinks {
		srcInfo, err = env.FS.Stat(src)
	} else {
		srcInfo, err = lstat(env.FS, src)
	}
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
			err = c.copyDir(env, srcPath, destPath, verbose, noClobber, update, preserve, followLinks)
		} else {
			err = c.copyFile(env, srcPath, destPath, verbose, noClobber, update, preserve, followLinks)
		}
		if err != nil {
			return err
		}
	}

	if preserve {
		_ = env.FS.Chtimes(dest, srcInfo.ModTime(), srcInfo.ModTime())
	}

	return nil
}

func lstat(fs afero.Fs, path string) (os.FileInfo, error) {
	if lstater, ok := fs.(afero.Lstater); ok {
		fi, _, err := lstater.LstatIfPossible(path)
		return fi, err
	}
	return fs.Stat(path)
}
