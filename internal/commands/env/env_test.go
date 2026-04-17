package envcmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestEnv_Run(t *testing.T) {
	t.Run("list variables", func(t *testing.T) {
		var stdout bytes.Buffer
		env := &commands.Environment{
			Stdout: &stdout,
			EnvVars: map[string]string{
				"FOO": "BAR",
			},
		}

		e := New()
		status := e.Run(context.Background(), env, nil)
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "FOO=BAR\n")
	})

	t.Run("ignore environment", func(t *testing.T) {
		var stdout bytes.Buffer
		env := &commands.Environment{
			Stdout: &stdout,
			EnvVars: map[string]string{
				"FOO": "BAR",
			},
		}

		e := New()
		status := e.Run(context.Background(), env, []string{"-i"})
		assert.Equal(t, 0, status)
		assert.Empty(t, stdout.String())
	})

	t.Run("unset variable", func(t *testing.T) {
		var stdout bytes.Buffer
		env := &commands.Environment{
			Stdout: &stdout,
			EnvVars: map[string]string{
				"FOO": "BAR",
				"BAZ": "QUX",
			},
		}

		e := New()
		status := e.Run(context.Background(), env, []string{"-u", "FOO"})
		assert.Equal(t, 0, status)
		assert.NotContains(t, stdout.String(), "FOO=BAR")
		assert.Contains(t, stdout.String(), "BAZ=QUX")
	})

	t.Run("execute command", func(t *testing.T) {
		var stdout bytes.Buffer
		registry := commands.New()
		
		// Mock command
		mockCmd := &mockCommand{name: "hi"}
		registry.Register(mockCmd)

		env := &commands.Environment{
			Stdout:   &stdout,
			Registry: registry,
			EnvVars: map[string]string{
				"FOO": "OLD",
			},
		}

		e := New()
		status := e.Run(context.Background(), env, []string{"FOO=NEW", "hi"})
		assert.Equal(t, 0, status)
		assert.Equal(t, "hi", mockCmd.executedWith[0])
		assert.Equal(t, "NEW", mockCmd.executedEnv.EnvVars["FOO"])
	})

	t.Run("null terminated", func(t *testing.T) {
		var stdout bytes.Buffer
		env := &commands.Environment{
			Stdout: &stdout,
			EnvVars: map[string]string{
				"A": "1",
				"B": "2",
			},
		}

		e := New()
		status := e.Run(context.Background(), env, []string{"-0"})
		assert.Equal(t, 0, status)
		assert.Equal(t, "A=1\x00B=2\x00", stdout.String())
	})

	t.Run("chdir", func(t *testing.T) {
		var stdout bytes.Buffer
		registry := commands.New()
		mockCmd := &mockCommand{name: "pwd"}
		registry.Register(mockCmd)

		env := &commands.Environment{
			Stdout:   &stdout,
			Registry: registry,
			Cwd:      "/home",
		}

		e := New()
		status := e.Run(context.Background(), env, []string{"-C", "/tmp", "pwd"})
		assert.Equal(t, 0, status)
		assert.Equal(t, "/tmp", mockCmd.executedEnv.Cwd)
	})
}

type mockCommand struct {
	name         string
	executedWith []string
	executedEnv  *commands.Environment
}

func (m *mockCommand) Name() string { return m.name }
func (m *mockCommand) Run(ctx context.Context, env *commands.Environment, args []string) int {
	m.executedWith = append([]string{m.name}, args...)
	m.executedEnv = env
	return 0
}
