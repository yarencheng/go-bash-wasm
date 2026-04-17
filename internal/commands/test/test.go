package testcmd

import (
	"context"
	"path/filepath"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Test struct {
	name string
}

func New(name string) *Test {
	return &Test{name: name}
}

func (t *Test) Name() string {
	return t.name
}

func (t *Test) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if t.name == "[" {
		if len(args) == 0 || args[len(args)-1] != "]" {
			return 2
		}
		args = args[:len(args)-1]
	}

	if len(args) == 0 {
		return 1
	}

	switch args[0] {
	case "-e":
		if len(args) < 2 {
			return 1
		}
		path := args[1]
		if !filepath.IsAbs(path) {
			path = filepath.Join(env.Cwd, path)
		}
		_, err := env.FS.Stat(path)
		if err == nil {
			return 0
		}
		return 1
	case "-f":
		if len(args) < 2 {
			return 1
		}
		path := args[1]
		if !filepath.IsAbs(path) {
			path = filepath.Join(env.Cwd, path)
		}
		info, err := env.FS.Stat(path)
		if err == nil && !info.IsDir() {
			return 0
		}
		return 1
	case "-d":
		if len(args) < 2 {
			return 1
		}
		path := args[1]
		if !filepath.IsAbs(path) {
			path = filepath.Join(env.Cwd, path)
		}
		info, err := env.FS.Stat(path)
		if err == nil && info.IsDir() {
			return 0
		}
		return 1
	case "-z":
		if len(args) < 2 {
			return 0
		}
		if args[1] == "" {
			return 0
		}
		return 1
	case "-n":
		if len(args) < 2 {
			return 1
		}
		if args[1] != "" {
			return 0
		}
		return 1
	default:
		if len(args) == 1 {
			if args[0] != "" {
				return 0
			}
			return 1
		}
	}

	return 1
}
