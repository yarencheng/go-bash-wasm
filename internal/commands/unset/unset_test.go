package unset

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUnset_Run(t *testing.T) {
	env := &commands.Environment{
		EnvVars: map[string]string{
			"VAR1": "val1",
			"VAR2": "val2",
		},
		Functions: map[string]string{
			"func1": "echo hello",
		},
		Stderr: io.Discard,
	}

	cmd := New()

	// Test default: remove variable
	status := cmd.Run(context.Background(), env, []string{"VAR1"})
	assert.Equal(t, 0, status)
	assert.NotContains(t, env.EnvVars, "VAR1")

	// Test default: remove function if variable not found
	status = cmd.Run(context.Background(), env, []string{"func1"})
	assert.Equal(t, 0, status)
	assert.NotContains(t, env.Functions, "func1")

	// Test -f (functions only)
	env.EnvVars["VAR2"] = "val2"
	env.Functions["func2"] = "echo world"
	status = cmd.Run(context.Background(), env, []string{"-f", "VAR2", "func2"})
	assert.Equal(t, 0, status)
	assert.Contains(t, env.EnvVars, "VAR2") // Should still be there
	assert.NotContains(t, env.Functions, "func2") // Should be removed

	// Test -v (variables only)
	env.Functions["func3"] = "echo"
	status = cmd.Run(context.Background(), env, []string{"-v", "VAR2", "func3"})
	assert.Equal(t, 0, status)
	assert.NotContains(t, env.EnvVars, "VAR2")
	assert.Contains(t, env.Functions, "func3")
}
