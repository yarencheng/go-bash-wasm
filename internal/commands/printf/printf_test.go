package printf

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPrintf_Complex(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	p := New()
	
	// Test %b
	stdout.Reset()
	status := p.Run(context.Background(), env, []string{"%b", "hello\\nworld"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello\nworld", stdout.String())

	// Test %q
	stdout.Reset()
	status = p.Run(context.Background(), env, []string{"%q", "hello world"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "'hello world'", stdout.String())

	// Test reuse format
	stdout.Reset()
	status = p.Run(context.Background(), env, []string{"<%s>", "a", "b", "c"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "<a><b><c>", stdout.String())
}
