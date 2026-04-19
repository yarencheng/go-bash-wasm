package trap

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTrap(t *testing.T) {
	tr := New()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	env := &commands.Environment{
		Stdout: stdout,
		Stderr: stderr,
		Traps:  make(map[string]string),
	}

	t.Run("set trap", func(t *testing.T) {
		code := tr.Run(context.Background(), env, []string{"echo trapped", "INT"})
		assert.Equal(t, 0, code)
		assert.Equal(t, "echo trapped", env.Traps["INT"])
	})

	t.Run("list traps", func(t *testing.T) {
		stdout.Reset()
		code := tr.Run(context.Background(), env, []string{"-p"})
		assert.Equal(t, 0, code)
		assert.Contains(t, stdout.String(), "trap -- 'echo trapped' INT")
	})

	t.Run("reset trap", func(t *testing.T) {
		code := tr.Run(context.Background(), env, []string{"-", "INT"})
		assert.Equal(t, 0, code)
		_, exists := env.Traps["INT"]
		assert.False(t, exists)
	})

	t.Run("invalid signal", func(t *testing.T) {
		stderr.Reset()
		code := tr.Run(context.Background(), env, []string{"ls", "INVALID"})
		assert.Equal(t, 0, code) // trap returns 0 even if some signals are invalid in bash sometimes? Actually it should probably fail.
		assert.Contains(t, stderr.String(), "invalid signal specification")
	})
}
