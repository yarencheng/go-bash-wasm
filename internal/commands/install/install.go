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
	_ = flags.BoolP("create-leading", "D", false, "create all leading components of DEST except the last (ignored)")
	_ = flags.BoolP("preserve-timestamps", "p", false, "apply access/modification times of SOURCE files to corresponding destination files (ignored)")
	_ = flags.BoolP("strip", "s", false, "strip symbol tables (ignored)")
	_ = flags.StringP("suffix", "S", "", "override the usual backup suffix (ignored)")
	_ = flags.StringP("target-directory", "t", "", "copy all SOURCE arguments into DIRECTORY (ignored)")
	_ = flags.BoolP("no-target-directory", "T", false, "treat DEST as a normal file (ignored)")
	_ = flags.BoolP("compare", "C", false, "compare content of source and destination files (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "install: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: install [OPTION]... [-T] SOURCE DEST\n")
		fmt.Fprintf(env.Stdout, "  or:  install [OPTION]... SOURCE... DIRECTORY\n")
		fmt.Fprintf(env.Stdout, "  or:  install [OPTION]... -t DIRECTORY SOURCE...\n")
		fmt.Fprintf(env.Stdout, "  or:  install [OPTION]... -d DIRECTORY...\n")
		fmt.Fprintf(env.Stdout, "Copy SOURCE to DEST or multiple SOURCE(s) to the existing DIRECTORY, while setting permission modes and owner/group.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "install")
		return 0
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
