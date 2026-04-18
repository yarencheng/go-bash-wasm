package chroot

import (
	"context"
	"fmt"

	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Chroot struct{}

func New() *Chroot {
	return &Chroot{}
}

func (c *Chroot) Name() string {
	return "chroot"
}

func (c *Chroot) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) < 1 {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "chroot: missing operand\n")
		}
		return 1
	}

	newRoot := args[0]
	cmdName := "/bin/sh"
	cmdArgs := []string{}

	if len(args) > 1 {
		cmdName = args[1]
		cmdArgs = args[2:]
	}

	// Verify newRoot exists and is a directory
	info, err := env.FS.Stat(newRoot)
	if err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "chroot: cannot change root directory to '%s': %v\n", newRoot, err)
		}
		return 1
	}
	if !info.IsDir() {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "chroot: cannot change root directory to '%s': Not a directory\n", newRoot)
		}
		return 1
	}

	// Wrap the filesystem
	originalFS := env.FS
	env.FS = afero.NewBasePathFs(originalFS, newRoot)

	// Reset CWD to / in the new FS
	originalCwd := env.Cwd
	env.Cwd = "/"

	defer func() {
		env.FS = originalFS
		env.Cwd = originalCwd
	}()

	if cmd, ok := env.Registry.Get(cmdName); ok {
		return cmd.Run(ctx, env, cmdArgs)
	}

	if env.Stderr != nil {
		fmt.Fprintf(env.Stderr, "chroot: %s: command not found\n", cmdName)
	}
	return 127
}
