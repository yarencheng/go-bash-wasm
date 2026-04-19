package whoami

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestWhoami_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		User:   "testuser",
	}

	w := New()
	status := w.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "testuser\n", stdout.String())
}
