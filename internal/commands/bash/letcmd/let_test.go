package letcmd

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestLet_Run(t *testing.T) {
	env := &commands.Environment{
		EnvVars: map[string]string{
			"a": "2",
		},
		Stdout: io.Discard,
		Stderr: io.Discard,
	}

	l := New()

	// Test simple assignment
	status := l.Run(context.Background(), env, []string{"b=5"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "5", env.EnvVars["b"])

	// Test shell variable usage and arithmetic
	// Note: Our current let implementation will be basic.
	// Bash let supports complex arithmetic.
}
