package umask

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUmask_Run(t *testing.T) {
	stdout := &bytes.Buffer{}
	env := &commands.Environment{
		Umask:  022,
		Stdout: stdout,
		Stderr: io.Discard,
	}

	cmd := New()

	// Test display
	status := cmd.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "0022\n", stdout.String())
	stdout.Reset()

	// Test -S (symbolic)
	status = cmd.Run(context.Background(), env, []string{"-S"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "u=rwx,g=rx,o=rx\n", stdout.String())
	stdout.Reset()

	// Test set
	status = cmd.Run(context.Background(), env, []string{"077"})
	assert.Equal(t, 0, status)
	assert.Equal(t, uint32(077), env.Umask)

	// Test -p
	status = cmd.Run(context.Background(), env, []string{"-p"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "umask 0077\n", stdout.String())
	stdout.Reset()
}
