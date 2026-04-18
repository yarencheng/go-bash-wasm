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
		Stdout:  &stdout,
		EnvVars: make(map[string]string),
	}

	h := New()
	status := h.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "wasm-host\n", stdout.String())
}

func TestHostname_Help(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
	}

	h := New()
	status := h.Run(context.Background(), env, []string{"--help"})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "Usage: hostname")
}

func TestHostname_Version(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
	}

	h := New()
	status := h.Run(context.Background(), env, []string{"--version"})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "hostname (go-bash-wasm)")
}
