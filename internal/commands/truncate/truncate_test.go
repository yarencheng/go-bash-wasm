package truncate

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTruncate_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello world"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-s", "5", "/test.txt"})
	assert.Equal(t, 0, status)

	f, err := afero.ReadFile(fs, "/test.txt")
	require.NoError(t, err)
	assert.Equal(t, 5, len(f))
	assert.Equal(t, "hello", string(f))
}

func TestTruncate_Create(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-s", "10", "/new.txt"})
	assert.Equal(t, 0, status)

	f, err := afero.ReadFile(fs, "/new.txt")
	require.NoError(t, err)
	assert.Equal(t, 10, len(f))
}
