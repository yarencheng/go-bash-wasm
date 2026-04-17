package unalias

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUnalias_Run(t *testing.T) {
	env := &commands.Environment{
		Aliases: map[string]string{
			"ll": "ls -l",
			"la": "ls -a",
		},
		Stderr: io.Discard,
	}

	cmd := New()

	// Test remove one
	status := cmd.Run(context.Background(), env, []string{"ll"})
	assert.Equal(t, 0, status)
	_, ok := env.Aliases["ll"]
	assert.False(t, ok)
	assert.True(t, len(env.Aliases) == 1)

	// Test -a (remove all)
	status = cmd.Run(context.Background(), env, []string{"-a"})
	assert.Equal(t, 0, status)
	assert.Empty(t, env.Aliases)

	// Test not found
	status = cmd.Run(context.Background(), env, []string{"nonexistent"})
	assert.Equal(t, 1, status)
}
