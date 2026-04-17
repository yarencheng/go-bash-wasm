package tr

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTr_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdin:  io.NopCloser(bytes.NewBufferString("hello world")),
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	trCmd := New()

	// Test translation
	status := trCmd.Run(context.Background(), env, []string{"a-z", "A-Z"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "HELLO WORLD", stdout.String())

	// Test deletion
	stdout.Reset()
	env.Stdin = io.NopCloser(bytes.NewBufferString("hello world"))
	status = trCmd.Run(context.Background(), env, []string{"-d", "l"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "heo word", stdout.String())
}
