package cd

import (
	"context"
	"fmt"
	"path"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Cd struct{}

func New() *Cd {
	return &Cd{}
}

func (c *Cd) Name() string {
	return "cd"
}

func (c *Cd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	target := "/"
	if len(args) > 0 {
		target = args[0]
	}

	newPath := target
	if !path.IsAbs(target) {
		newPath = path.Join(env.Cwd, target)
	}
	newPath = path.Clean(newPath)

	info, err := env.FS.Stat(newPath)
	if err != nil {
		fmt.Fprintf(env.Stderr, "cd: %s: No such file or directory\n", target)
		return 1
	}

	if !info.IsDir() {
		fmt.Fprintf(env.Stderr, "cd: %s: Not a directory\n", target)
		return 1
	}

	env.Cwd = newPath
	return 0
}
