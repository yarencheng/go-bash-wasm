package hostname

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestHostname_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
	}

	h := New()
	status := h.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "wasm-host\n", stdout.String())
}
