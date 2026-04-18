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
	s := &Shell{
		Registry: registry,
		Env:      env,
	}
	env.Executor = s
	return s
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

		s.Execute(context.Background(), line)

		if s.Env.ExitRequested {
			break
		}
	}
	return nil
}

// Execute parses and runs a command line, supporting sequential commands and expansion.
func (s *Shell) Execute(ctx context.Context, line string) int {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "#") {
		return 0
	}

	s.Env.History = append(s.Env.History, line)

	// Support sequential commands separated by ;
	commandsList := strings.Split(line, ";")
	lastExitCode := 0

	for _, cmdLine := range commandsList {
		cmdLine = strings.TrimSpace(cmdLine)
		if cmdLine == "" {
			continue
		}

		// Expand variables
		cmdLine = s.expand(cmdLine)

		// Support pipeline negation !
		negate := false
		if strings.HasPrefix(cmdLine, "!") {
			negate = true
			cmdLine = strings.TrimSpace(cmdLine[1:])
		}

		// Parse simple command
		args := strings.Fields(cmdLine)
		if len(args) == 0 {
			continue
		}

		cmdName := args[0]
		cmdArgs := args[1:]

		exitCode := 0
		if cmd, ok := s.Registry.Get(cmdName); ok {
			exitCode = cmd.Run(ctx, s.Env, cmdArgs)
		} else {
			fmt.Fprintf(s.Env.Stderr, "%s: command not found\n", cmdName)
			exitCode = 127
		}

		if negate {
			if exitCode == 0 {
				lastExitCode = 1
			} else {
				lastExitCode = 0
			}
		} else {
			lastExitCode = exitCode
		}

		if s.Env.ExitRequested {
			break
		}
	}

	return lastExitCode
}

func (s *Shell) expand(line string) string {
	// Simple variable expansion: $VAR or ${VAR} or ${VAR:-default}
	
	result := line
	
	// Handle ${VAR} and ${VAR:-default}
	for strings.Contains(result, "${") {
		start := strings.Index(result, "${")
		end := strings.Index(result[start:], "}")
		if end == -1 {
			break
		}
		end += start
		expr := result[start+2 : end]
		
		val := s.resolveVariable(expr)
		result = result[:start] + val + result[end+1:]
	}
	
	// Simple $VAR expansion (non-greedy)
	// This is a rough approximation
	return result
}

func (s *Shell) resolveVariable(expr string) string {
	name := expr
	def := ""
	hasDefault := false

	if strings.Contains(expr, ":-") {
		parts := strings.SplitN(expr, ":-", 2)
		name = parts[0]
		def = parts[1]
		hasDefault = true
	}

	val, ok := s.Env.EnvVars[name]
	if !ok || (hasDefault && val == "") {
		// Handle dynamic variables
		switch name {
		case "RANDOM":
			return fmt.Sprintf("%d", time.Now().UnixNano()%32768)
		case "SECONDS":
			return fmt.Sprintf("%d", int(time.Since(s.Env.StartTime).Seconds()))
		case "EPOCHSECONDS":
			return fmt.Sprintf("%d", time.Now().Unix())
		case "UID":
			return fmt.Sprintf("%d", s.Env.Uid)
		case "GID":
			return fmt.Sprintf("%d", s.Env.Gid)
		case "EUID":
			return fmt.Sprintf("%d", s.Env.Uid)
		case "HOSTNAME":
			return s.Env.EnvVars["HOSTNAME"]
		}
		
		if hasDefault {
			return def
		}
		return ""
	}
	return val
}
