package commands

import (
	"context"
	"io"
	"time"

	"github.com/spf13/afero"
)

type Job struct {
	ID      int
	PID     int
	Command string
	Status  string // Running, Stopped, Done
}

type CompSpec struct {
	Actions      uint64
	Options      uint64
	GlobPat      string
	WordList     string
	Prefix       string
	Suffix       string
	FunctionName string
	Command      string
	FilterPat    string
}

// Environment defines the execution environment for a command.
type Environment struct {
	FS                afero.Fs
	Stdin             io.ReadCloser
	Stdout            io.Writer
	Stderr            io.Writer
	Cwd               string
	User              string
	Uid               int
	Gid               int
	Umask             uint32
	Groups            []int
	StartTime         time.Time
	ExitRequested     bool
	ExitCode          int
	ReturnRequested   bool
	ReturnCode        int
	BreakRequested    int
	ContinueRequested int
	EnvVars           map[string]string
	PositionalArgs    []string
	Arrays            map[string][]string
	AssocArrays       map[string]map[string]string
	DirStack          []string
	Aliases           map[string]string
	Functions         map[string]string
	Hash              map[string]string
	History           []string
	Jobs              []*Job
	Completions       map[string]*CompSpec
	Shopts            map[string]bool
	Traps             map[string]string
	Registry          *Registry
	Executor          Executor
}

// Executor defines the interface for executing shell commands.
type Executor interface {
	Execute(ctx context.Context, line string) int
}

// Command is the interface that all shell commands must implement.
type Command interface {
	// Name returns the name of the command (e.g., "ls").
	Name() string
	// Run executes the command with the given arguments.
	// Returns the exit status code.
	Run(ctx context.Context, env *Environment, args []string) int
}
