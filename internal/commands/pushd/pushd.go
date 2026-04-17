package pushd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Pushd struct{}

func New() *Pushd {
	return &Pushd{}
}

func (p *Pushd) Name() string {
	return "pushd"
}

func (p *Pushd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		// pushd without args swaps top two elements?
		// but standard bash behavior for pushd with no args is to swap cwd and stack[0]
		if len(env.DirStack) == 0 {
			fmt.Fprintln(env.Stderr, "pushd: no other directory")
			return 1
		}
		newDir := env.DirStack[len(env.DirStack)-1]
		oldDir := env.Cwd
		env.Cwd = newDir
		env.DirStack[len(env.DirStack)-1] = oldDir
		fmt.Printf("%s %v\n", env.Cwd, env.DirStack) // should use dirs format
		return 0
	}

	dir := args[0]
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(env.Cwd, dir)
	}

	// Verify dir exists
	info, err := env.FS.Stat(dir)
	if err != nil || !info.IsDir() {
		fmt.Fprintf(env.Stderr, "pushd: %s: No such directory\n", dir)
		return 1
	}

	env.DirStack = append(env.DirStack, env.Cwd)
	env.Cwd = dir
	
	// Print stack (dirs style)
	stack := append([]string{env.Cwd}, env.DirStack...)
	for i, d := range stack {
		fmt.Fprint(env.Stdout, d)
		if i < len(stack)-1 {
			fmt.Fprint(env.Stdout, " ")
		}
	}
	fmt.Fprintln(env.Stdout, "")

	return 0
}
