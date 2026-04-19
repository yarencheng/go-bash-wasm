package rmdir

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

func TestRmdir_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, fs.Mkdir("/testdir", 0755))
	require.NoError(t, fs.Mkdir("/emptydir", 0755))
	require.NoError(t, afero.WriteFile(fs, "/testdir/file.txt", []byte("data"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	r := New()

	// Test basic rmdir success
	status := r.Run(context.Background(), env, []string{"emptydir"})
	assert.Equal(t, 0, status)
	_, err := fs.Stat("/emptydir")
	assert.Error(t, err)

	// Test rmdir non-empty failure
	status = r.Run(context.Background(), env, []string{"testdir"})
	assert.Equal(t, 1, status)
	_, err = fs.Stat("/testdir")
	assert.NoError(t, err)

	// Test parents flag
	require.NoError(t, fs.MkdirAll("/a/b/c", 0755))
	status = r.Run(context.Background(), env, []string{"-p", "a/b/c"})
	assert.Equal(t, 0, status)
	_, err = fs.Stat("/a/b/c")
	assert.Error(t, err)
	_, err = fs.Stat("/a/b")
	assert.Error(t, err)
	_, err = fs.Stat("/a")
	assert.Error(t, err)

	// Test verbose
	require.NoError(t, fs.Mkdir("/vdir", 0755))
	stdout := &bytes.Buffer{}
	env.Stdout = stdout
	status = r.Run(context.Background(), env, []string{"-v", "vdir"})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "removing directory, 'vdir'")

	// Test not a directory
	require.NoError(t, afero.WriteFile(fs, "/notadir", []byte("data"), 0644))
	status = r.Run(context.Background(), env, []string{"notadir"})
	assert.Equal(t, 1, status)

	// Test non-existent
	status = r.Run(context.Background(), env, []string{"nonexistent"})
	assert.Equal(t, 1, status)

	// Test missing operand
	status = r.Run(context.Background(), env, []string{})
	assert.Equal(t, 1, status)

	// Test invalid flag
	status = r.Run(context.Background(), env, []string{"--invalid"})
	assert.Equal(t, 1, status)
}

func TestRmdir_Metadata(t *testing.T) {
	r := New()
	assert.Equal(t, "rmdir", r.Name())
}
