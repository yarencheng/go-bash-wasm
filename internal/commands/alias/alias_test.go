package alias

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestAlias_Run(t *testing.T) {
	env := &commands.Environment{
		Aliases: make(map[string]string),
	}

	a := New()
	status := a.Run(context.Background(), env, []string{"ll=ls -l"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "ls -l", env.Aliases["ll"])
}
