package returncmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestReturn(t *testing.T) {
	env := &commands.Environment{
		ExitCode: 42,
	}
	r := New()

	t.Run("default return", func(t *testing.T) {
		env.ReturnRequested = false
		env.ReturnCode = 0
		code := r.Run(context.Background(), env, []string{})
		assert.Equal(t, 42, code)
		assert.True(t, env.ReturnRequested)
		assert.Equal(t, 42, env.ReturnCode)
	})

	t.Run("return with code", func(t *testing.T) {
		env.ReturnRequested = false
		env.ReturnCode = 0
		code := r.Run(context.Background(), env, []string{"7"})
		assert.Equal(t, 7, code)
		assert.True(t, env.ReturnRequested)
		assert.Equal(t, 7, env.ReturnCode)
	})

	t.Run("invalid code", func(t *testing.T) {
		env.ReturnRequested = false
		env.ReturnCode = 0
		env.ExitCode = 10
		code := r.Run(context.Background(), env, []string{"abc"})
		assert.Equal(t, 10, code)
		assert.True(t, env.ReturnRequested)
		assert.Equal(t, 10, env.ReturnCode)
	})
}
