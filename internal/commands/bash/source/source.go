package source

import (
	"bufio"
	"context"
	"fmt"
	"path/filepath"
	"strings"

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
	var pathstring string
	i := 0
	for ; i < len(args); i++ {
		if args[i] == "-p" && i+1 < len(args) {
			pathstring = args[i+1]
			i++
		} else if strings.HasPrefix(args[i], "-p") && len(args[i]) > 2 {
			pathstring = args[i][2:]
		} else if args[i] == "--" {
			i++
			break
		} else if strings.HasPrefix(args[i], "-") {
			// ignore unknown flags or handle them if needed
		} else {
			break
		}
	}

	if i >= len(args) {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "%s: filename argument required\n", s.Name())
		}
		return 1
	}

	filename := args[i]
	scriptArgs := args[i+1:]

	var fullPath string
	if filepath.IsAbs(filename) || strings.Contains(filename, "/") {
		fullPath = filename
		if !filepath.IsAbs(fullPath) {
			fullPath = filepath.Join(env.Cwd, fullPath)
		}
	} else if pathstring != "" {
		// Search in pathstring
		paths := filepath.SplitList(pathstring)
		for _, p := range paths {
			if p == "" {
				p = "."
			}
			target := filepath.Join(p, filename)
			if !filepath.IsAbs(target) {
				target = filepath.Join(env.Cwd, target)
			}
			if info, err := env.FS.Stat(target); err == nil && !info.IsDir() {
				fullPath = target
				break
			}
		}
	} else {
		// Search in PATH
		pathVar := env.EnvVars["PATH"]
		paths := filepath.SplitList(pathVar)
		for _, p := range paths {
			if p == "" {
				p = "."
			}
			target := filepath.Join(p, filename)
			if !filepath.IsAbs(target) {
				target = filepath.Join(env.Cwd, target)
			}
			if info, err := env.FS.Stat(target); err == nil && !info.IsDir() {
				fullPath = target
				break
			}
		}
		// If not found in PATH, try CWD
		if fullPath == "" {
			target := filepath.Join(env.Cwd, filename)
			if info, err := env.FS.Stat(target); err == nil && !info.IsDir() {
				fullPath = target
			}
		}
	}

	if fullPath == "" {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "%s: %s: file not found\n", s.Name(), filename)
		}
		return 1
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

	// Save positional parameters
	oldPosArgs := env.PositionalArgs
	if len(scriptArgs) > 0 {
		env.PositionalArgs = scriptArgs
	}

	scanner := bufio.NewScanner(f)
	lastExitCode := 0
	for scanner.Scan() {
		line := scanner.Text()
		lastExitCode = env.Executor.Execute(ctx, line)
		if env.ExitRequested || env.ReturnRequested {
			break
		}
	}

	// Restore positional parameters if they were changed
	if len(scriptArgs) > 0 {
		env.PositionalArgs = oldPosArgs
	}

	if err := scanner.Err(); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "%s: error reading file %s: %v\n", s.Name(), filename, err)
		}
		return 1
	}

	if env.ReturnRequested {
		env.ReturnRequested = false
		return env.ReturnCode
	}

	return lastExitCode
}
