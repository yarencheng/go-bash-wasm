package command

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/echo"
)

func TestCommand_Run(t *testing.T) {
	registry := commands.New()
	registry.Register(echo.New())
	
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdout:   out,
		Registry: registry,
	}

	c := New()
	status := c.Run(context.Background(), env, []string{"echo", "hello"})
	assert.Equal(t, 0, status)
	assert.Contains(t, out.String(), "hello")
}
