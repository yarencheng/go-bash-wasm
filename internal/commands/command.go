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
	VarAttributes     map[string]uint32 // Bitmask: 1=readonly, 2=export, 4=integer, 8=array, 16=assoc
	Registry          *Registry
	Executor          Executor
}

const (
	AttrReadonly  uint32 = 1
	AttrExport    uint32 = 2
	AttrInteger   uint32 = 4
	AttrArray     uint32 = 8
	AttrAssoc     uint32 = 16
	AttrFunction  uint32 = 32
	AttrNameref   uint32 = 64
	AttrLowercase uint32 = 128
	AttrUppercase uint32 = 256
)

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
