package testcmd

import (
	"context"
	"fmt"
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

	res := t.eval(env, args)
	if res {
		return 0
	}
	return 1
}

func (t *Test) eval(env *commands.Environment, args []string) bool {
	if len(args) == 0 {
		return false
	}

	// Handle logical OR (-o)
	for i := 0; i < len(args); i++ {
		if args[i] == "-o" {
			return t.eval(env, args[:i]) || t.eval(env, args[i+1:])
		}
	}

	// Handle logical AND (-a)
	for i := 0; i < len(args); i++ {
		if args[i] == "-a" {
			return t.eval(env, args[:i]) && t.eval(env, args[i+1:])
		}
	}

	// Handle negation (!)
	if args[0] == "!" {
		return !t.eval(env, args[1:])
	}

	// Unary operators
	if len(args) == 2 {
		return t.evalUnary(env, args[0], args[1])
	}

	// Binary operators
	if len(args) == 3 {
		return t.evalBinary(env, args[0], args[1], args[2])
	}

	// Default: true if non-empty string
	if len(args) == 1 {
		return args[0] != ""
	}

	return false
}

func (t *Test) evalUnary(env *commands.Environment, op, arg string) bool {
	switch op {
	case "-e":
		path := t.absPath(env, arg)
		_, err := env.FS.Stat(path)
		return err == nil
	case "-f":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && !info.IsDir()
	case "-d":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && info.IsDir()
	case "-z":
		return arg == ""
	case "-n":
		return arg != ""
	}
	return false
}

func (t *Test) evalBinary(env *commands.Environment, left, op, right string) bool {
	switch op {
	case "=", "==":
		return left == right
	case "!=":
		return left != right
	case "-eq":
		return t.toInt(left) == t.toInt(right)
	case "-ne":
		return t.toInt(left) != t.toInt(right)
	case "-lt":
		return t.toInt(left) < t.toInt(right)
	case "-le":
		return t.toInt(left) <= t.toInt(right)
	case "-gt":
		return t.toInt(left) > t.toInt(right)
	case "-ge":
		return t.toInt(left) >= t.toInt(right)
	}
	return false
}

func (t *Test) absPath(env *commands.Environment, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(env.Cwd, path)
}

func (t *Test) toInt(s string) int64 {
	var n int64
	fmt.Sscanf(s, "%d", &n)
	return n
}
