package testcmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
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
	case "-s":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && info.Size() > 0
	case "-L", "-h":
		path := t.absPath(env, arg)
		if l, ok := env.FS.(afero.Lstater); ok {
			info, _, err := l.LstatIfPossible(path)
			return err == nil && (info.Mode()&os.ModeSymlink != 0)
		}
		info, err := env.FS.Stat(path)
		return err == nil && (info.Mode()&os.ModeSymlink != 0)
	case "-r", "-w":
		// TODO: Implement proper permission checks using env.Uid/Gid/Groups
		path := t.absPath(env, arg)
		_, err := env.FS.Stat(path)
		return err == nil
	case "-x":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && (info.Mode()&0111 != 0)
	case "-S":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && (info.Mode()&os.ModeSocket != 0)
	case "-p":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && (info.Mode()&os.ModeNamedPipe != 0)
	case "-b":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && (info.Mode()&os.ModeDevice != 0) && (info.Mode()&os.ModeCharDevice == 0)
	case "-c":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && (info.Mode()&os.ModeCharDevice != 0)
	case "-u":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && (info.Mode()&os.ModeSetuid != 0)
	case "-g":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && (info.Mode()&os.ModeSetgid != 0)
	case "-k":
		path := t.absPath(env, arg)
		info, err := env.FS.Stat(path)
		return err == nil && (info.Mode()&os.ModeSticky != 0)
	case "-v":
		_, ok := env.EnvVars[arg]
		return ok
	case "-t":
		// FD defaults to 0 if not specified (though -t requires an arg in test)
		fd := t.toInt(arg)
		return fd >= 0 && fd <= 2 // In our simulator, standard FDs are always "terminals"
	case "-o":
		val, ok := env.Shopts[arg]
		return ok && val
	case "-O":
		path := t.absPath(env, arg)
		_, err := env.FS.Stat(path)
		// Simulation: User owns all files in the sandbox
		return err == nil
	case "-G":
		path := t.absPath(env, arg)
		_, err := env.FS.Stat(path)
		// Simulation: User's group owns all files in the sandbox
		return err == nil
	case "-N":
		path := t.absPath(env, arg)
		_, err := env.FS.Stat(path)
		if err != nil {
			return false
		}
		// Simulation: Approximate newer than atime
		// afero doesn't track atime well, but we can assume if Mtime > Atime (simulated)
		// Typical shell behavior is Mtime > Atime.
		return true // Mock behavior
	case "-R":
		// Mock nameref check
		return false
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
	case "-nt":
		i1, err1 := env.FS.Stat(t.absPath(env, left))
		i2, err2 := env.FS.Stat(t.absPath(env, right))
		if err1 != nil {
			return false
		}
		if err2 != nil {
			return true
		}
		return i1.ModTime().After(i2.ModTime())
	case "-ot":
		i1, err1 := env.FS.Stat(t.absPath(env, left))
		i2, err2 := env.FS.Stat(t.absPath(env, right))
		if err1 != nil {
			return false
		}
		if err2 != nil {
			return false
		}
		return i1.ModTime().Before(i2.ModTime())
	case "-ef":
		p1 := t.absPath(env, left)
		p2 := t.absPath(env, right)
		if p1 == p2 {
			return true
		}
		i1, err1 := env.FS.Stat(p1)
		i2, err2 := env.FS.Stat(p2)
		if err1 != nil || err2 != nil {
			return false
		}
		return os.SameFile(i1, i2)
	case "<":
		return left < right
	case ">":
		return left > right
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
