package find

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestFind_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, fs.Mkdir("/dir1", 0755))
	require.NoError(t, afero.WriteFile(fs, "/dir1/file1.txt", []byte("data"), 0644))
	require.NoError(t, fs.Mkdir("/dir1/subdir", 0755))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	f := New()

	// Test basic find
	status := f.Run(context.Background(), env, []string{"/dir1"})
	assert.Equal(t, 0, status)
	output := env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "/dir1\n")
	assert.Contains(t, output, "/dir1/file1.txt\n")
	assert.Contains(t, output, "/dir1/subdir\n")

	// Test name filter
	env.Stdout = &bytes.Buffer{}
	status = f.Run(context.Background(), env, []string{"/dir1", "-name", "*.txt"})
	assert.Equal(t, 0, status)
	output = env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "/dir1/file1.txt\n")
	assert.NotContains(t, output, "/dir1/subdir\n")
}
