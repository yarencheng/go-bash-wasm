package shell

import (
	"context"
	"fmt"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

// LineReader defines the interface for reading lines from the terminal.
type LineReader interface {
	Readline() (string, error)
	Close() error
}

// Shell provides an interactive command-line environment.
type Shell struct {
	Registry *commands.Registry
	Env      *commands.Environment
}

// New creates a new shell with the given registry and environment.
func New(registry *commands.Registry, env *commands.Environment) *Shell {
	return &Shell{
		Registry: registry,
		Env:      env,
	}
}

// RunInteractive starts the REPL loop.
func (s *Shell) RunInteractive() error {
	rl, err := newLineReader(s.Env)
	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			// io.EOF is returned when Ctrl+D is pressed
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Parse simple command
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		cmdName := args[0]
		cmdArgs := args[1:]

		if cmd, ok := s.Registry.Get(cmdName); ok {
			_ = cmd.Run(context.Background(), s.Env, cmdArgs)
		} else {
			fmt.Fprintf(s.Env.Stderr, "%s: command not found\n", cmdName)
		}

		if s.Env.ExitRequested {
			break
		}
	}
	return nil
}
