package ulimit

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUlimit(t *testing.T) {
	u := New()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	env := &commands.Environment{
		Stdout: stdout,
		Stderr: stderr,
	}

	t.Run("default file size", func(t *testing.T) {
		stdout.Reset()
		code := u.Run(context.Background(), env, []string{})
		assert.Equal(t, 0, code)
		assert.Contains(t, stdout.String(), "unlimited")
	})

	t.Run("list all", func(t *testing.T) {
		stdout.Reset()
		code := u.Run(context.Background(), env, []string{"-a"})
		assert.Equal(t, 0, code)
		assert.Contains(t, stdout.String(), "open files")
		assert.Contains(t, stdout.String(), "stack size")
	})

	t.Run("specific flag", func(t *testing.T) {
		stdout.Reset()
		code := u.Run(context.Background(), env, []string{"-n"})
		assert.Equal(t, 0, code)
		assert.Equal(t, "1024\n", stdout.String())
	})

	t.Run("not supported setting", func(t *testing.T) {
		stderr.Reset()
		code := u.Run(context.Background(), env, []string{"100"})
		assert.Equal(t, 1, code)
		assert.Contains(t, stderr.String(), "setting limit not supported")
	})
}
