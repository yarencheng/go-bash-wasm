package commands

import (
	"context"
	"io"

	"github.com/spf13/afero"
)

// Environment defines the execution environment for a command.
type Environment struct {
	FS     afero.Fs
	Stdin  io.ReadCloser
	Stdout io.Writer
	Stderr io.Writer
	Cwd    string
}

// Command is the interface that all shell commands must implement.
type Command interface {
	// Name returns the name of the command (e.g., "ls").
	Name() string
	// Run executes the command with the given arguments.
	// Returns the exit status code.
	Run(ctx context.Context, env *Environment, args []string) int
}
