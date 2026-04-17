package clear

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestClear_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	c := New()
	status := c.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "\033[H\033[2J")
}
