package logname

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestLogname_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		User:   "wasm",
	}

	l := New()
	status := l.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "wasm\n", stdout.String())
}
