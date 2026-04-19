package clear

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestClear_Run(t *testing.T) {
	stdout := &bytes.Buffer{}
	env := &commands.Environment{
		Stdout: stdout,
	}
	c := New()
	status := c.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "\033[H\033[2J", stdout.String())
}

func TestClear_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "clear", c.Name())
}
