package shell

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"mvdan.cc/sh/v3/syntax"
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

// Execute parses and runs a command line, supporting sequential commands, expansion, and pipelines.
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

	// Use mvdan.cc/sh/syntax to parse the command line
	reader := strings.NewReader(line)
	parser := syntax.NewParser()
	f, err := parser.Parse(reader, "")
	if err != nil {
		fmt.Fprintf(s.Env.Stderr, "bash: syntax error: %v\n", err)
		return 2
	}

	lastExitCode := 0
	for _, stmt := range f.Stmts {
		lastExitCode = s.executeStmt(ctx, s.Env, stmt)
		if s.Env.ExitRequested {
			break
		}
	}

	return lastExitCode
}

func (s *Shell) executeStmt(ctx context.Context, env *commands.Environment, stmt *syntax.Stmt) int {
	if stmt == nil {
		return 0
	}

	// Handle background execution &
	if stmt.Background {
		// Basic stub for background execution in simulator
		jobID := len(env.Jobs) + 1
		job := &commands.Job{
			ID:      jobID,
			Command: s.stmtToString(stmt),
			Status:  "Running",
		}
		env.Jobs = append(env.Jobs, job)
		fmt.Fprintf(env.Stdout, "[%d] %d\n", job.ID, job.PID)
	}

	// Handle redirections
	oldStdout := env.Stdout
	oldStderr := env.Stderr
	oldStdin := env.Stdin
	var closers []io.Closer

	for _, redir := range stmt.Redirs {
		rExitCode := s.applyRedirection(env, redir, &closers)
		if rExitCode != 0 {
			s.cleanupClosers(closers)
			return rExitCode
		}
	}

	exitCode := 0
	if stmt.Cmd != nil {
		exitCode = s.executeCmd(ctx, env, stmt.Cmd)
	}

	// Cleanup redirections
	s.cleanupClosers(closers)
	env.Stdout = oldStdout
	env.Stderr = oldStderr
	env.Stdin = oldStdin

	if stmt.Negated {
		if exitCode == 0 {
			exitCode = 1
		} else {
			exitCode = 0
		}
	}

	return exitCode
}

func (s *Shell) executeCmd(ctx context.Context, env *commands.Environment, cmd syntax.Command) int {
	switch c := cmd.(type) {
	case *syntax.CallExpr:
		return s.executeCallExpr(ctx, env, c)
	case *syntax.BinaryCmd:
		return s.executeBinaryCmd(ctx, env, c)
	case *syntax.Block:
		return s.executeBlock(ctx, env, c)
	case *syntax.Subshell:
		return s.executeSubshell(ctx, env, c)
	case *syntax.IfClause:
		return s.executeIfClause(ctx, env, c)
	case *syntax.ForClause:
		return s.executeForClause(ctx, env, c)
	case *syntax.WhileClause:
		return s.executeWhileClause(ctx, env, c)
	case *syntax.CaseClause:
		return s.executeCaseClause(ctx, env, c)
	case *syntax.ArithmCmd:
		return s.executeArithmCmd(ctx, env, c)
	case *syntax.TestClause:
		return s.executeTestClause(ctx, env, c)
	case *syntax.DeclClause:
		// declaration like 'local' or 'declare'
		args := []*syntax.Word{{Parts: []syntax.WordPart{&syntax.Lit{Value: c.Variant.Value}}}}
		for _, a := range c.Args {
			if a.Name != nil {
				val := a.Name.Value + "="
				if a.Value != nil {
					val += s.wordToString(env, a.Value)
				}
				args = append(args, &syntax.Word{Parts: []syntax.WordPart{&syntax.Lit{Value: val}}})
			}
		}
		return s.executeCallExpr(ctx, env, &syntax.CallExpr{Args: args})
	case *syntax.TimeClause:
		return s.executeStmt(ctx, env, c.Stmt)
	case *syntax.FuncDecl:
		env.Functions[c.Name.Value] = s.stmtToString(c.Body)
		return 0
	default:
		fmt.Fprintf(env.Stderr, "bash: unsupported command type: %T\n", cmd)
		return 1
	}
}

func (s *Shell) executeStmts(ctx context.Context, env *commands.Environment, stmts []*syntax.Stmt) int {
	lastExitCode := 0
	for _, stmt := range stmts {
		lastExitCode = s.executeStmt(ctx, env, stmt)
	}
	return lastExitCode
}

func (s *Shell) executeCallExpr(ctx context.Context, env *commands.Environment, c *syntax.CallExpr) int {
	if len(c.Args) == 0 {
		// Assignment only, e.g., VAR=val
		for _, assign := range c.Assigns {
			name := assign.Name.Value
			val := s.wordToString(env, assign.Value)
			env.EnvVars[name] = val
		}
		return 0
	}

	// Expand arguments
	var args []string
	for _, arg := range c.Args {
		expanded := s.expandWord(env, arg)
		args = append(args, expanded...)
	}

	if len(args) == 0 {
		return 0
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	if cmd, ok := s.Registry.Get(cmdName); ok {
		return cmd.Run(ctx, env, cmdArgs)
	}

	fmt.Fprintf(env.Stderr, "%s: command not found\n", cmdName)
	return 127
}

func (s *Shell) executeBinaryCmd(ctx context.Context, env *commands.Environment, c *syntax.BinaryCmd) int {
	switch c.Op {
	case syntax.Pipe, syntax.PipeAll:
		return s.executePipeline(ctx, env, c)
	case syntax.AndStmt:
		exitCode := s.executeStmt(ctx, env, c.X)
		if exitCode == 0 {
			return s.executeStmt(ctx, env, c.Y)
		}
		return exitCode
	case syntax.OrStmt:
		exitCode := s.executeStmt(ctx, env, c.X)
		if exitCode != 0 {
			return s.executeStmt(ctx, env, c.Y)
		}
		return exitCode
	default:
		fmt.Fprintf(env.Stderr, "bash: unsupported binary operator: %s\n", c.Op.String())
		return 1
	}
}

func (s *Shell) executePipeline(ctx context.Context, env *commands.Environment, c *syntax.BinaryCmd) int {
	// Create a pipe
	pr, pw := io.Pipe()

	// Clone environment for the left side (subshell behavior)
	leftEnv := s.cloneEnv(env)
	leftEnv.Stdout = pw
	if c.Op == syntax.PipeAll {
		leftEnv.Stderr = pw
	}

	// Channel to wait for the left side to finish
	leftDone := make(chan int, 1)

	// Execute left side in a goroutine
	go func() {
		defer pw.Close()
		exitCode := s.executeStmt(ctx, leftEnv, c.X)
		leftDone <- exitCode
	}()

	// Execute right side
	oldStdin := env.Stdin
	env.Stdin = pr
	defer func() {
		env.Stdin = oldStdin
		pr.Close()
	}()

	exitCode := s.executeStmt(ctx, env, c.Y)

	// Wait for left side
	<-leftDone

	return exitCode
}

func (s *Shell) executeBlock(ctx context.Context, env *commands.Environment, b *syntax.Block) int {
	lastExitCode := 0
	for _, stmt := range b.Stmts {
		lastExitCode = s.executeStmt(ctx, env, stmt)
		if env.ExitRequested || env.ReturnRequested || env.BreakRequested > 0 || env.ContinueRequested > 0 {
			break
		}
	}
	return lastExitCode
}

func (s *Shell) executeSubshell(ctx context.Context, env *commands.Environment, sub *syntax.Subshell) int {
	subEnv := s.cloneEnv(env)
	lastExitCode := 0
	for _, stmt := range sub.Stmts {
		lastExitCode = s.executeStmt(ctx, subEnv, stmt)
		if subEnv.ExitRequested {
			break
		}
	}
	return lastExitCode
}

func (s *Shell) applyRedirection(env *commands.Environment, redir *syntax.Redirect, closers *[]io.Closer) int {
	filename := s.wordToString(env, redir.Word)
	path := s.resolvePath(filename)

	var fd int
	if redir.N != nil {
		fmt.Sscanf(redir.N.Value, "%d", &fd)
	} else {
		// Default fds if not specified
		switch redir.Op {
		case syntax.RdrOut, syntax.AppOut, syntax.RdrAll, syntax.AppAll:
			fd = 1
		case syntax.RdrIn:
			fd = 0
		}
	}

	switch redir.Op {
	case syntax.RdrOut: // >
		f, err := env.FS.Create(path)
		if err != nil {
			fmt.Fprintf(env.Stderr, "bash: %s: %v\n", filename, err)
			return 1
		}
		if fd == 1 {
			env.Stdout = f
		} else if fd == 2 {
			env.Stderr = f
		}
		*closers = append(*closers, f)
	case syntax.AppOut: // >>
		f, err := env.FS.OpenFile(path, 0x1|0x40|0x8, 0644) // O_WRONLY|O_CREATE|O_APPEND
		if err != nil {
			fmt.Fprintf(env.Stderr, "bash: %s: %v\n", filename, err)
			return 1
		}
		if fd == 1 {
			env.Stdout = f
		} else if fd == 2 {
			env.Stderr = f
		}
		*closers = append(*closers, f)
	case syntax.RdrIn: // <
		f, err := env.FS.Open(path)
		if err != nil {
			fmt.Fprintf(env.Stderr, "bash: %s: %v\n", filename, err)
			return 1
		}
		env.Stdin = f
		*closers = append(*closers, f)
	case syntax.RdrAll: // &>
		f, err := env.FS.Create(path)
		if err != nil {
			fmt.Fprintf(env.Stderr, "bash: %s: %v\n", filename, err)
			return 1
		}
		env.Stdout = f
		env.Stderr = f
		*closers = append(*closers, f)
	case syntax.AppAll: // &>>
		f, err := env.FS.OpenFile(path, 0x1|0x40|0x8, 0644)
		if err != nil {
			fmt.Fprintf(env.Stderr, "bash: %s: %v\n", filename, err)
			return 1
		}
		env.Stdout = f
		env.Stderr = f
		*closers = append(*closers, f)
	default:
		fmt.Fprintf(env.Stderr, "bash: unsupported redirection: %s\n", redir.Op.String())
		return 1
	}

	return 0
}

func (s *Shell) cleanupClosers(closers []io.Closer) {
	for _, c := range closers {
		_ = c.Close()
	}
}

func (s *Shell) wordToString(env *commands.Environment, w *syntax.Word) string {
	if w == nil {
		return ""
	}
	var sb strings.Builder
	for _, part := range w.Parts {
		switch p := part.(type) {
		case *syntax.Lit:
			sb.WriteString(p.Value)
		case *syntax.SglQuoted:
			sb.WriteString(p.Value)
		case *syntax.DblQuoted:
			for _, qp := range p.Parts {
				switch q := qp.(type) {
				case *syntax.Lit:
					sb.WriteString(q.Value)
				case *syntax.ParamExp:
					sb.WriteString(s.resolveParamExp(env, q))
				}
			}
		case *syntax.ParamExp:
			sb.WriteString(s.resolveParamExp(env, p))
		}
	}
	return sb.String()
}

func (s *Shell) expandWord(env *commands.Environment, w *syntax.Word) []string {
	// 1. Parameter expansion
	val := s.wordToString(env, w)

	// 2. Globbing
	if strings.ContainsAny(val, "*?") {
		matches, err := filepath.Glob(s.resolvePath(val))
		if err == nil && len(matches) > 0 {
			return matches
		}
	}

	return []string{val}
}

func (s *Shell) resolveParamExp(env *commands.Environment, p *syntax.ParamExp) string {
	name := p.Param.Value
	val, ok := env.EnvVars[name]
	if !ok {
		// Dynamic variables
		switch name {
		case "RANDOM":
			return fmt.Sprintf("%d", time.Now().UnixNano()%32768)
		case "SECONDS":
			return fmt.Sprintf("%d", int(time.Since(env.StartTime).Seconds()))
		case "UID", "EUID":
			return fmt.Sprintf("%d", env.Uid)
		case "GID":
			return fmt.Sprintf("%d", env.Gid)
		}
		return ""
	}
	return val
}

func (s *Shell) stmtToString(stmt *syntax.Stmt) string {
	return "cmd" // Placeholder
}

func (s *Shell) cloneEnv(env *commands.Environment) *commands.Environment {
	newEnv := *env // Shallow copy mostly

	// Deep copy maps and slices
	newEnv.EnvVars = make(map[string]string)
	for k, v := range env.EnvVars {
		newEnv.EnvVars[k] = v
	}

	newEnv.Aliases = make(map[string]string)
	for k, v := range env.Aliases {
		newEnv.Aliases[k] = v
	}

	newEnv.Arrays = make(map[string][]string)
	for k, v := range env.Arrays {
		newEnv.Arrays[k] = make([]string, len(v))
		copy(newEnv.Arrays[k], v)
	}

	newEnv.Shopts = make(map[string]bool)
	for k, v := range env.Shopts {
		newEnv.Shopts[k] = v
	}

	newEnv.Traps = make(map[string]string)
	for k, v := range env.Traps {
		newEnv.Traps[k] = v
	}

	newEnv.Functions = make(map[string]string)
	for k, v := range env.Functions {
		newEnv.Functions[k] = v
	}

	newEnv.Hash = make(map[string]string)
	for k, v := range env.Hash {
		newEnv.Hash[k] = v
	}

	return &newEnv
}

func (s *Shell) executeIfClause(ctx context.Context, env *commands.Environment, ifClause *syntax.IfClause) int {
	exitCode := s.executeStmts(ctx, env, ifClause.Cond)
	if exitCode == 0 {
		return s.executeStmts(ctx, env, ifClause.Then)
	}
	if ifClause.Else != nil {
		return s.executeCmd(ctx, env, ifClause.Else)
	}
	return 0
}

func (s *Shell) executeForClause(ctx context.Context, env *commands.Environment, forClause *syntax.ForClause) int {
	if forClause.Select {
		// Basic select loop
		var items []string
		switch f := forClause.Loop.(type) {
		case *syntax.WordIter:
			for _, w := range f.Items {
				items = append(items, s.expandWord(env, w)...)
			}
		}

		if len(items) == 0 {
			return 0
		}

		ps3 := env.EnvVars["PS3"]
		if ps3 == "" {
			ps3 = "#? "
		}

		for {
			for i, item := range items {
				fmt.Fprintf(env.Stdout, "%d) %s\n", i+1, item)
			}
			fmt.Fprint(env.Stdout, ps3)
			reader := bufio.NewReader(env.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			input = strings.TrimSpace(input)
			if input == "" {
				continue
			}
			idx, err := strconv.Atoi(input)
			if err != nil || idx < 1 || idx > len(items) {
				env.EnvVars[forClause.Loop.(*syntax.WordIter).Name.Value] = ""
			} else {
				env.EnvVars[forClause.Loop.(*syntax.WordIter).Name.Value] = items[idx-1]
			}
			s.executeStmts(ctx, env, forClause.Do)
		}
		return 0
	}

	var items []string
	switch f := forClause.Loop.(type) {
	case *syntax.WordIter:
		for _, w := range f.Items {
			items = append(items, s.expandWord(env, w)...)
		}
		var name string
		if f.Name != nil {
			name = f.Name.Value
		}
		lastExitCode := 0
		for _, item := range items {
			if name != "" {
				env.EnvVars[name] = item
			}
			lastExitCode = s.executeStmts(ctx, env, forClause.Do)
		}
		return lastExitCode
	case *syntax.CStyleLoop:
		// Arithmetic for loop not yet fully supported
		return 0
	}
	return 0
}

func (s *Shell) executeWhileClause(ctx context.Context, env *commands.Environment, whileClause *syntax.WhileClause) int {
	lastExitCode := 0
	for {
		exitCode := s.executeStmts(ctx, env, whileClause.Cond)
		if exitCode != 0 {
			break
		}
		lastExitCode = s.executeStmts(ctx, env, whileClause.Do)
	}
	return lastExitCode
}

func (s *Shell) executeCaseClause(ctx context.Context, env *commands.Environment, caseClause *syntax.CaseClause) int {
	val := s.wordToString(env, caseClause.Word)
	for _, item := range caseClause.Items {
		for _, pattern := range item.Patterns {
			pat := s.wordToString(env, pattern)
			if matchPattern(pat, val) {
				return s.executeStmts(ctx, env, item.Stmts)
			}
		}
	}
	return 0
}

func (s *Shell) executeArithmCmd(ctx context.Context, env *commands.Environment, arithm *syntax.ArithmCmd) int {
	// Simple stub for arithmetic command (( expression ))
	return 0
}

func (s *Shell) executeTestClause(ctx context.Context, env *commands.Environment, test *syntax.TestClause) int {
	// Simple stub for [[ expression ]]
	return 0
}

func (s *Shell) resolvePath(p string) string {
	if filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(s.Env.Cwd, p)
}

func matchPattern(pat, val string) bool {
	matched, _ := filepath.Match(pat, val)
	return matched
}
