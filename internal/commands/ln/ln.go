package ln

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Ln struct{}

func New() *Ln {
	return &Ln{}
}

func (l *Ln) Name() string {
	return "ln"
}

func (l *Ln) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("ln", pflag.ContinueOnError)
	symbolic := flags.BoolP("symbolic", "s", false, "make symbolic links instead of hard links")
	force := flags.BoolP("force", "f", false, "remove existing destination files")
	verbose := flags.BoolP("verbose", "v", false, "print name of each linked file")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "ln: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) < 2 {
		fmt.Fprintf(env.Stderr, "ln: missing file operand\n")
		return 1
	}

	oldName := targets[0]
	if !filepath.IsAbs(oldName) {
		oldName = filepath.Join(env.Cwd, oldName)
	}

	newName := targets[1]
	if !filepath.IsAbs(newName) {
		newName = filepath.Join(env.Cwd, newName)
	}

	if *force {
		_ = env.FS.Remove(newName)
	}

	if *symbolic {
		if sl, ok := env.FS.(interface {
			SymlinkIfPossible(oldname, newname string) error
		}); ok {
			if err := sl.SymlinkIfPossible(oldName, newName); err != nil {
				fmt.Fprintf(env.Stderr, "ln: %v\n", err)
				return 1
			}
			if *verbose {
				fmt.Fprintf(env.Stdout, "'%s' -> '%s'\n", newName, oldName)
			}
			return 0
		}
		fmt.Fprintf(env.Stderr, "ln: symbolic links not supported by filesystem\n")
		return 1
	}

	// Hard link
	if linker, ok := env.FS.(interface {
		Link(oldname, newname string) error
	}); ok {
		if err := linker.Link(oldName, newName); err != nil {
			fmt.Fprintf(env.Stderr, "ln: %v\n", err)
			return 1
		}
		if *verbose {
			fmt.Fprintf(env.Stdout, "'%s' => '%s'\n", newName, oldName)
		}
		return 0
	}
	fmt.Fprintf(env.Stderr, "ln: hard links not supported by filesystem\n")
	return 1
}
