package unalias

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUnalias_Run(t *testing.T) {
	env := &commands.Environment{
		Aliases: map[string]string{"ll": "ls -l"},
	}

	u := New()
	status := u.Run(context.Background(), env, []string{"ll"})
	assert.Equal(t, 0, status)
	_, exists := env.Aliases["ll"]
	assert.False(t, exists)
}
