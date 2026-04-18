package enable

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockCmd struct{ name string }
func (m *mockCmd) Name() string { return m.name }
func (m *mockCmd) Run(ctx context.Context, env *commands.Environment, args []string) int { return 0 }

func TestEnable(t *testing.T) {
	registry := commands.New()
	registry.Register(&mockCmd{"test1"})
	registry.Register(&mockCmd{"test2"})

	e := New()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	env := &commands.Environment{
		Stdout:   stdout,
		Stderr:   stderr,
		Registry: registry,
	}

	t.Run("list enabled", func(t *testing.T) {
		stdout.Reset()
		code := e.Run(context.Background(), env, []string{})
		assert.Equal(t, 0, code)
		assert.Contains(t, stdout.String(), "enable test1")
		assert.Contains(t, stdout.String(), "enable test2")
	})

	t.Run("disable command", func(t *testing.T) {
		stdout.Reset()
		code := e.Run(context.Background(), env, []string{"-n", "test1"})
		assert.Equal(t, 0, code)
		assert.False(t, registry.IsEnabled("test1"))
	})

	t.Run("list disabled", func(t *testing.T) {
		stdout.Reset()
		code := e.Run(context.Background(), env, []string{"-n"})
		assert.Equal(t, 0, code)
		assert.Contains(t, stdout.String(), "enable -n test1")
		assert.NotContains(t, stdout.String(), "enable test2")
	})

	t.Run("enable command", func(t *testing.T) {
		stdout.Reset()
		code := e.Run(context.Background(), env, []string{"test1"})
		assert.Equal(t, 0, code)
		assert.True(t, registry.IsEnabled("test1"))
	})

	t.Run("not a builtin", func(t *testing.T) {
		stderr.Reset()
		code := e.Run(context.Background(), env, []string{"invalid"})
		assert.Equal(t, 1, code)
		assert.Contains(t, stderr.String(), "not a shell builtin")
	})
}
