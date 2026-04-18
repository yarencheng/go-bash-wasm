package shell

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockCommand struct {
	name string
	run  func(ctx context.Context, env *commands.Environment, args []string) int
}

func (m *mockCommand) Name() string { return m.name }
func (m *mockCommand) Run(ctx context.Context, env *commands.Environment, args []string) int {
	return m.run(ctx, env, args)
}

func setupTestShell() (*Shell, *commands.Environment, *strings.Builder, *strings.Builder) {
	registry := commands.New()
	stdout := &strings.Builder{}
	stderr := &strings.Builder{}
	env := &commands.Environment{
		FS:        afero.NewMemMapFs(),
		Stdin:     io.NopCloser(strings.NewReader("")),
		Stdout:    stdout,
		Stderr:    stderr,
		Cwd:       "/",
		EnvVars:   make(map[string]string),
		StartTime: time.Now(),
		Registry:  registry,
	}
	s := New(registry, env)
	return s, env, stdout, stderr
}

func TestExecuteBasic(t *testing.T) {
	s, env, stdout, _ := setupTestShell()
	
	env.Registry.Register(&mockCommand{
		name: "echo",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			fmt.Fprint(env.Stdout, strings.Join(args, " "))
			return 0
		},
	})

	exitCode := s.Execute(context.Background(), "echo hello world")
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "hello world", stdout.String())
}

func TestExecutePipeline(t *testing.T) {
	s, env, stdout, _ := setupTestShell()

	env.Registry.Register(&mockCommand{
		name: "echo",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			fmt.Fprint(env.Stdout, strings.Join(args, " "))
			return 0
		},
	})

	env.Registry.Register(&mockCommand{
		name: "cat",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			_, _ = io.Copy(env.Stdout, env.Stdin)
			return 0
		},
	})

	exitCode := s.Execute(context.Background(), "echo hello | cat")
	assert.Equal(t, 0, exitCode)
	// We might have a small race condition or timing issue with goroutines if not careful,
	// but io.Pipe and io.Copy should handle it.
	// Wait a bit if needed or ensure synchronous behavior where possible.
	assert.Equal(t, "hello", stdout.String())
}

func TestExecutePipelineCombined(t *testing.T) {
	s, env, stdout, _ := setupTestShell()

	env.Registry.Register(&mockCommand{
		name: "err_echo",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			fmt.Fprint(env.Stdout, "out")
			fmt.Fprint(env.Stderr, "err")
			return 0
		},
	})

	env.Registry.Register(&mockCommand{
		name: "cat",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			_, _ = io.Copy(env.Stdout, env.Stdin)
			return 0
		},
	})

	// Case 1: normal pipe | (only stdout)
	s.Execute(context.Background(), "err_echo | cat")
	assert.Equal(t, "out", stdout.String())
	stdout.Reset()

	// Case 2: combined pipe |& (stdout + stderr)
	s.Execute(context.Background(), "err_echo |& cat")
	assert.Contains(t, stdout.String(), "out")
	assert.Contains(t, stdout.String(), "err")
}

func TestExecuteLogical(t *testing.T) {
	s, env, _, _ := setupTestShell()

	env.Registry.Register(&mockCommand{
		name: "true",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			return 0
		},
	})
	env.Registry.Register(&mockCommand{
		name: "false",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			return 1
		},
	})

	assert.Equal(t, 0, s.Execute(context.Background(), "true && true"))
	assert.Equal(t, 1, s.Execute(context.Background(), "true && false"))
	assert.Equal(t, 0, s.Execute(context.Background(), "false || true"))
	assert.Equal(t, 1, s.Execute(context.Background(), "false || false"))
}

func TestExecuteRedirection(t *testing.T) {
	s, env, _, _ := setupTestShell()

	env.Registry.Register(&mockCommand{
		name: "echo",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			fmt.Fprint(env.Stdout, strings.Join(args, " "))
			return 0
		},
	})

	exitCode := s.Execute(context.Background(), "echo hello > /test.txt")
	assert.Equal(t, 0, exitCode)

	exists, _ := afero.Exists(env.FS, "/test.txt")
	assert.True(t, exists)

	data, _ := afero.ReadFile(env.FS, "/test.txt")
	assert.Equal(t, "hello", string(data))
}
