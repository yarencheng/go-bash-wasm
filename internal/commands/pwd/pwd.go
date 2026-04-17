package pwd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Pwd struct{}

func New() *Pwd {
	return &Pwd{}
}

func (p *Pwd) Name() string {
	return "pwd"
}

func (p *Pwd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("pwd", pflag.ContinueOnError)
	logical := flags.BoolP("logical", "L", true, "use PWD from environment, even if it contains symlinks")
	physical := flags.BoolP("physical", "P", false, "avoid all symlinks")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "pwd: %v\n", err)
		}
		return 2
	}

	path := env.Cwd
	if *physical {
		// Use realpath logic or EvalSymlinks if possible with afero
		// For now, satisfy with EvalSymlinks on the effective path if it's a real FS, 
		// but since we use MemMapFs, we might need to be careful.
		// Actually, let's just use filepath.Clean for now as a baseline, 
		// but if we want true -P we should resolve it.
		if resolved, err := filepath.EvalSymlinks(env.Cwd); err == nil {
			path = resolved
		} else {
			// If EvalSymlinks fails (e.g. on virtual fs), just clean it
			path = filepath.Clean(env.Cwd)
		}
	} else if *logical {
		// Default behavior
		path = env.Cwd
	}

	fmt.Fprintln(env.Stdout, path)
	return 0
}
