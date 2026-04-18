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

	if s.Env.EnvVars == nil {
		s.Env.EnvVars = make(map[string]string)
	}
	lineno := 0
	fmt.Sscanf(s.Env.EnvVars["LINENO"], "%d", &lineno)
	s.Env.EnvVars["LINENO"] = fmt.Sprintf("%d", lineno+1)

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
	if strings.HasPrefix(expr, "#") {
		name := expr[1:]
		val, _ := s.Env.EnvVars[name]
		return fmt.Sprintf("%d", len(val))
	}

	name := expr
	def := ""
	hasDefault := false
	hasSubstring := false
	hasCaseMod := false
	hasRemoval := false
	hasReplace := false
	replacePattern := ""
	replaceStr := ""
	replaceAll := false
	removalType := ""
	removalPattern := ""
	caseModType := ""
	offset := 0
	length := -1

	// Check for replacement operator
	if strings.Contains(name, "/") {
		parts := strings.Split(name, "/")
		name = parts[0]
		hasReplace = true
		if len(parts) > 1 {
			patternExpr := parts[1]
			if strings.HasPrefix(patternExpr, "/") {
				replaceAll = true
				replacePattern = patternExpr[1:]
			} else {
				replacePattern = patternExpr
			}
		}
		if len(parts) > 2 {
			replaceStr = parts[2]
		}
	}

	// Check for removal operators
	if strings.Contains(name, "##") {
		parts := strings.SplitN(name, "##", 2)
		name = parts[0]
		hasRemoval = true
		removalType = "##"
		removalPattern = parts[1]
	} else if strings.Contains(name, "#") {
		parts := strings.SplitN(name, "#", 2)
		name = parts[0]
		hasRemoval = true
		removalType = "#"
		removalPattern = parts[1]
	} else if strings.Contains(name, "%%") {
		parts := strings.SplitN(name, "%%", 2)
		name = parts[0]
		hasRemoval = true
		removalType = "%%"
		removalPattern = parts[1]
	} else if strings.Contains(name, "%") {
		parts := strings.SplitN(name, "%", 2)
		name = parts[0]
		hasRemoval = true
		removalType = "%"
		removalPattern = parts[1]
	}

	if strings.HasSuffix(name, "^^") {
		name = name[:len(name)-2]
		hasCaseMod = true
		caseModType = "^^"
	} else if strings.HasSuffix(name, "^") {
		name = name[:len(name)-1]
		hasCaseMod = true
		caseModType = "^"
	} else if strings.HasSuffix(name, ",,") {
		name = name[:len(name)-2]
		hasCaseMod = true
		caseModType = ",,"
	} else if strings.HasSuffix(name, ",") {
		name = name[:len(name)-1]
		hasCaseMod = true
		caseModType = ","
	}

	if strings.Contains(expr, ":-") {
		parts := strings.SplitN(expr, ":-", 2)
		name = parts[0]
		def = parts[1]
		hasDefault = true
	} else if strings.Contains(expr, ":+") {
		parts := strings.SplitN(expr, ":+", 2)
		name = parts[0]
		alt := parts[1]
		val, ok := s.Env.EnvVars[name]
		if ok && val != "" {
			return alt
		}
		return ""
	} else if strings.Contains(expr, ":?") {
		parts := strings.SplitN(expr, ":?", 2)
		name = parts[0]
		errMsg := parts[1]
		val, ok := s.Env.EnvVars[name]
		if !ok || val == "" {
			if errMsg == "" {
				errMsg = "parameter null or unset"
			}
			fmt.Fprintf(s.Env.Stderr, "%s: %s\n", name, errMsg)
			return "" // In real bash it might exit, but for simulator we just print error
		}
		return val
	} else if strings.Contains(expr, ":") {
		parts := strings.Split(expr, ":")
		name = parts[0]
		hasSubstring = true
		if len(parts) > 1 {
			fmt.Sscanf(parts[1], "%d", &offset)
		}
		if len(parts) > 2 {
			fmt.Sscanf(parts[2], "%d", &length)
		}
	}

	val, ok := s.Env.EnvVars[name]
	if !ok || (hasDefault && val == "") {
		// Handle dynamic variables
		switch name {
		case "RANDOM":
			val = fmt.Sprintf("%d", time.Now().UnixNano()%32768)
		case "SECONDS":
			val = fmt.Sprintf("%d", int(time.Since(s.Env.StartTime).Seconds()))
		case "EPOCHSECONDS":
			val = fmt.Sprintf("%d", time.Now().Unix())
		case "UID":
			val = fmt.Sprintf("%d", s.Env.Uid)
		case "GID":
			val = fmt.Sprintf("%d", s.Env.Gid)
		case "EUID":
			val = fmt.Sprintf("%d", s.Env.Uid)
		case "LINENO":
			val = s.Env.EnvVars["LINENO"]
		case "HOSTNAME":
			val = s.Env.EnvVars["HOSTNAME"]
		default:
			if hasDefault {
				return def
			}
			return ""
		}
	}

	if hasReplace {
		if replaceAll {
			val = strings.ReplaceAll(val, replacePattern, replaceStr)
		} else {
			val = strings.Replace(val, replacePattern, replaceStr, 1)
		}
	}

	if hasRemoval {
		// Basic prefix/suffix removal (simulated without full glob for now)
		switch removalType {
		case "#", "##": // Prefix removal
			if strings.HasPrefix(val, removalPattern) {
				val = val[len(removalPattern):]
			}
		case "%", "%%": // Suffix removal
			if strings.HasSuffix(val, removalPattern) {
				val = val[:len(val)-len(removalPattern)]
			}
		}
	}

	if hasSubstring {
		runes := []rune(val)
		if offset < 0 {
			offset = len(runes) + offset
		}
		if offset < 0 {
			offset = 0
		}
		if offset > len(runes) {
			return ""
		}
		if length == -1 {
			return string(runes[offset:])
		}
		if length < 0 {
			length = len(runes) + length - offset
		}
		end := offset + length
		if end > len(runes) {
			end = len(runes)
		}
		if end < offset {
			return ""
		}
		return string(runes[offset:end])
	}

	if hasCaseMod {
		runes := []rune(val)
		if len(runes) == 0 {
			return ""
		}
		switch caseModType {
		case "^^":
			return strings.ToUpper(val)
		case "^":
			runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
			return string(runes)
		case ",,":
			return strings.ToLower(val)
		case ",":
			runes[0] = []rune(strings.ToLower(string(runes[0])))[0]
			return string(runes)
		}
	}

	return val
}
