package who

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestWho_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Stderr: io.Discard,
		User:   "user",
	}

	w := New()
	status := w.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "user")
}
