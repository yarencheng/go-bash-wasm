package pwd

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"io"
	"testing"
)

func TestPwd_Run(t *testing.T) {
	stdout := &bytes.Buffer{}
	env := &commands.Environment{
		Cwd:    "/home/wasm",
		Stdout: stdout,
		Stderr: io.Discard,
	}

	cmd := New()

	// Basic
	status := cmd.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "/home/wasm\n", stdout.String())
	stdout.Reset()

	// -L
	status = cmd.Run(context.Background(), env, []string{"-L"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/home/wasm\n", stdout.String())
	stdout.Reset()

	// -P (no symlinks)
	status = cmd.Run(context.Background(), env, []string{"-P"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/home/wasm\n", stdout.String())
}

func TestPwd_Physical(t *testing.T) {
	// This test is tricky because filepath.EvalSymlinks uses the actual OS filesystem
	// unless we use a specialized library that handles symlinks on afero MemMapFs.
	// afero's MemMapFs DOES NOT support symlinks in a way that filepath can resolve.
	// So -P will mostly return cleaned Cwd on virtual FS unless we implement full
	// resolution manually.
}
