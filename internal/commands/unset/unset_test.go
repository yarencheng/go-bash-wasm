package unset

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUnset_Run(t *testing.T) {
	env := &commands.Environment{
		EnvVars: map[string]string{"FOO": "BAR"},
	}

	u := New()
	status := u.Run(context.Background(), env, []string{"FOO"})
	assert.Equal(t, 0, status)
	_, exists := env.EnvVars["FOO"]
	assert.False(t, exists)
}
