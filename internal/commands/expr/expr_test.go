package expr

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestExpr_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	e := New()

	// Test arithmetic
	status := e.Run(context.Background(), env, []string{"1", "+", "2"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "3\n", stdout.String())

	// Test comparison
	stdout.Reset()
	status = e.Run(context.Background(), env, []string{"5", ">", "3"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "1\n", stdout.String())

	// Test length
	stdout.Reset()
	status = e.Run(context.Background(), env, []string{"length", "hello"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "5\n", stdout.String())

	// Test index
	stdout.Reset()
	status = e.Run(context.Background(), env, []string{"index", "hello", "e"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "2\n", stdout.String())

	// Test substr
	stdout.Reset()
	status = e.Run(context.Background(), env, []string{"substr", "hello", "2", "3"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "ell\n", stdout.String())

	// Test logical |
	stdout.Reset()
	status = e.Run(context.Background(), env, []string{"0", "|", "5"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "5\n", stdout.String())

	// Test logical &
	stdout.Reset()
	status = e.Run(context.Background(), env, []string{"5", "&", "3"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "5\n", stdout.String())

	// Test help
	stdout.Reset()
	status = e.Run(context.Background(), env, []string{"--help"})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "Usage: expr EXPRESSION")

	// Test version
	stdout.Reset()
	status = e.Run(context.Background(), env, []string{"--version"})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "expr (GNU coreutils) 9.10")
}
