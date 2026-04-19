package source

import (
	"bufio"
	"context"
	"fmt"
	"path/filepath"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Source struct {
	alias bool
}

func New() *Source {
	return &Source{alias: false}
}

func NewDot() *Source {
	return &Source{alias: true}
}

func (s *Source) Name() string {
	if s.alias {
		return "."
	}
	return "source"
}

func (s *Source) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "%s: filename argument required\n", s.Name())
		}
		return 1
	}

	filename := args[0]
	fullPath := filename
	if !filepath.IsAbs(filename) {
		fullPath = filepath.Join(env.Cwd, filename)
	}

	f, err := env.FS.Open(fullPath)
	if err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "%s: %s: %v\n", s.Name(), filename, err)
		}
		return 1
	}
	defer f.Close()

	if env.Executor == nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "%s: executor not available\n", s.Name())
		}
		return 1
	}

	scanner := bufio.NewScanner(f)
	lastExitCode := 0
	for scanner.Scan() {
		line := scanner.Text()
		lastExitCode = env.Executor.Execute(ctx, line)
		if env.ExitRequested {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "%s: error reading file %s: %v\n", s.Name(), filename, err)
		}
		return 1
	}

	return lastExitCode
}
