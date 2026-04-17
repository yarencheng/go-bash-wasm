package exit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestExit_Run(t *testing.T) {
	env := &commands.Environment{}

	e := New()
	status := e.Run(context.Background(), env, []string{"42"})
	assert.Equal(t, 42, status)
	assert.True(t, env.ExitRequested)
	assert.Equal(t, 42, env.ExitCode)
}
