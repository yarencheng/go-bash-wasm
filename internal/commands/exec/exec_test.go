package exec

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/echo"
)

func TestExec_Run(t *testing.T) {
	var stdout bytes.Buffer
	registry := commands.New()
	registry.Register(echo.New())

	env := &commands.Environment{
		Stdout:   &stdout,
		Registry: registry,
	}

	e := New()
	status := e.Run(context.Background(), env, []string{"echo", "hello"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello\n", stdout.String())
	assert.True(t, env.ExitRequested)
}
