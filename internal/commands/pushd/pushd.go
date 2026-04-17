package pushd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
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
	flags := pflag.NewFlagSet("pushd", pflag.ContinueOnError)
	noChdir := flags.BoolP("no-chdir", "n", false, "suppress the normal change of directory")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "pushd: %v\n", err)
		}
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		if len(env.DirStack) == 0 {
			if env.Stderr != nil {
				fmt.Fprintln(env.Stderr, "pushd: no other directory")
			}
			return 1
		}
		newDir := env.DirStack[len(env.DirStack)-1]
		oldDir := env.Cwd
		env.Cwd = newDir
		env.DirStack[len(env.DirStack)-1] = oldDir
	} else {
		dir := targets[0]
		if !filepath.IsAbs(dir) {
			dir = filepath.Join(env.Cwd, dir)
		}

		// Verify dir exists
		info, err := env.FS.Stat(dir)
		if err != nil || !info.IsDir() {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "pushd: %s: No such directory\n", dir)
			}
			return 1
		}

		if *noChdir {
			// Add to stack but don't change CWD
			// Bash adds it to the second position (env.DirStack[0])
			env.DirStack = append([]string{dir}, env.DirStack...)
		} else {
			env.DirStack = append([]string{env.Cwd}, env.DirStack...)
			env.Cwd = dir
		}
	}

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
