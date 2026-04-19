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

	// Test squeeze
	stdout.Reset()
	env.Stdin = io.NopCloser(bytes.NewBufferString("aaabbbccc"))
	status = trCmd.Run(context.Background(), env, []string{"-s", "abc"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "abc", stdout.String())

	// Test complement
	stdout.Reset()
	env.Stdin = io.NopCloser(bytes.NewBufferString("abc123def"))
	status = trCmd.Run(context.Background(), env, []string{"-c", "a-z", " "})
	assert.Equal(t, 0, status)
	assert.Equal(t, "abc   def", stdout.String())

	// Test delete complement
	stdout.Reset()
	env.Stdin = io.NopCloser(bytes.NewBufferString("abc123def"))
	status = trCmd.Run(context.Background(), env, []string{"-cd", "a-z"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "abcdef", stdout.String())

	// Test error cases
	assert.Equal(t, 1, trCmd.Run(context.Background(), env, []string{"a-z"})) // Missing second set
	assert.Equal(t, 1, trCmd.Run(context.Background(), env, []string{}))    // Missing operands
}

func TestTr_Metadata(t *testing.T) {
	cmd := New()
	assert.Equal(t, "tr", cmd.Name())
}
