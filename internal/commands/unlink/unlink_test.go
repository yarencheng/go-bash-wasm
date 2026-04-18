package unlink

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUnlink_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("test"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: &bytes.Buffer{},
	}

	u := New()

	// Test basic unlink
	status := u.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)

	exists, _ := afero.Exists(fs, "/test.txt")
	assert.False(t, exists)

	// Test unlink non-existent
	status = u.Run(context.Background(), env, []string{"/nonexistent"})
	assert.Equal(t, 1, status)

	// Test unlink too many args
	status = u.Run(context.Background(), env, []string{"a", "b"})
	assert.Equal(t, 1, status)
}
