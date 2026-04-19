package hostid

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestHostid_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	h := New()

	status := h.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.NotEmpty(t, stdout.String())
}
