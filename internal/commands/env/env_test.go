package envcmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestEnv_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		EnvVars: map[string]string{
			"FOO": "BAR",
		},
	}

	e := New()
	status := e.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "FOO=BAR\n")
}
