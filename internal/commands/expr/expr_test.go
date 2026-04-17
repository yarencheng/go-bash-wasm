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
}
