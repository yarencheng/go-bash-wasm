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
	noClobber := flags.BoolP("no-clobber", "n", false, "do not overwrite an existing file")
	targetDir := flags.StringP("target-directory", "t", "", "move all SOURCE arguments into DIRECTORY")
	noTargetDir := flags.BoolP("no-target-directory", "T", false, "treat DEST as a normal file")
	update := flags.BoolP("update", "u", false, "move only when the SOURCE file is newer than the destination file or when the destination file is missing")
	_ = flags.BoolP("backup", "b", false, "make a backup of each existing destination file (ignored)")
	_ = flags.BoolP("context", "Z", false, "set SELinux security context of destination file to default type (ignored)")
	_ = flags.Bool("exchange", false, "exchange source and destination (ignored)")
	_ = flags.Bool("no-copy", false, "do not copy if rename fails (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "mv: %v\n", err)
		}
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: mv [OPTION]... [-T] SOURCE DEST\n")
		fmt.Fprintf(env.Stdout, "  or:  mv [OPTION]... SOURCE... DIRECTORY\n")
		fmt.Fprintf(env.Stdout, "  or:  mv [OPTION]... -t DIRECTORY SOURCE...\n")
		fmt.Fprintf(env.Stdout, "Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "mv")
		return 0
	}

	posArgs := flags.Args()
	var sources []string
	var dest string

	if *targetDir != "" {
		sources = posArgs
		dest = *targetDir
	} else {
		if len(posArgs) < 2 {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "mv: missing file operand\n")
			}
			return 1
		}
		sources = posArgs[:len(posArgs)-1]
		dest = posArgs[len(posArgs)-1]
	}

	destFullPath := dest
	if !filepath.IsAbs(dest) {
		destFullPath = filepath.Join(env.Cwd, dest)
	}

	destInfo, destErr := env.FS.Stat(destFullPath)
	isDestDir := destErr == nil && destInfo.IsDir() && !*noTargetDir

	if len(sources) > 1 && !isDestDir {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "mv: target '%s' is not a directory\n", dest)
		}
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

		// Check no-clobber
		if *noClobber {
			if _, err := env.FS.Stat(finalDest); err == nil {
				continue // skip existing
			}
		}

		// Check update
		if *update {
			srcInfo, err := env.FS.Stat(srcFullPath)
			if err == nil {
				if dInfo, err := env.FS.Stat(finalDest); err == nil {
					if !srcInfo.ModTime().After(dInfo.ModTime()) {
						continue // skip older or same age
					}
				}
			}
		}

		err := env.FS.Rename(srcFullPath, finalDest)
		if err != nil {
			// If rename fails (e.g. cross-device), try copy + delete
			// In MemMapFs it shouldn't fail for cross-device typically but good practice
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "mv: cannot move '%s' to '%s': %v\n", src, finalDest, err)
			}
			exitCode = 1
			continue
		}

		if *verbose {
			fmt.Fprintf(env.Stdout, "renamed '%s' -> '%s'\n", src, finalDest)
		}
	}

	return exitCode
}
