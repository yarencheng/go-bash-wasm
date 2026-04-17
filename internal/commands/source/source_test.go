package source

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockExecutor struct {
	env *commands.Environment
}

func (m *mockExecutor) Execute(ctx context.Context, line string) int {
	line = strings.TrimSpace(line)
	if line == "" {
		return 0
	}
	args := strings.Fields(line)
	if args[0] == "setvar" {
		m.env.EnvVars[args[1]] = args[2]
		return 0
	}
	if args[0] == "exit" {
		m.env.ExitRequested = true
		return 0
	}
	return 1
}

func TestSource_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/script.sh", []byte("setvar FOO bar\nsetvar BAZ qux\n"), 0644)

	env := &commands.Environment{
		FS:      fs,
		Cwd:     "/",
		EnvVars: make(map[string]string),
		Stderr:  io.Discard,
	}
	executor := &mockExecutor{env: env}
	env.Executor = executor

	s := New()
	status := s.Run(context.Background(), env, []string{"script.sh"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "bar", env.EnvVars["FOO"])
	assert.Equal(t, "qux", env.EnvVars["BAZ"])
}

func TestSource_Exit(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/script.sh", []byte("setvar FOO bar\nexit\nsetvar BAZ qux\n"), 0644)

	env := &commands.Environment{
		FS:      fs,
		Cwd:     "/",
		EnvVars: make(map[string]string),
		Stderr:  io.Discard,
	}
	executor := &mockExecutor{env: env}
	env.Executor = executor

	s := New()
	status := s.Run(context.Background(), env, []string{"script.sh"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "bar", env.EnvVars["FOO"])
	assert.Empty(t, env.EnvVars["BAZ"])
	assert.True(t, env.ExitRequested)
}
