package chmod

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestChmod_Numeric(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("test"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: &bytes.Buffer{},
	}

	c := New()
	status := c.Run(context.Background(), env, []string{"755", "/test.txt"})
	assert.Equal(t, 0, status)

	info, err := fs.Stat("/test.txt")
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0755), info.Mode().Perm())
}

func TestChmod_Symbolic(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("test"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: &bytes.Buffer{},
	}

	c := New()

	// u+x
	status := c.Run(context.Background(), env, []string{"u+x", "/test.txt"})
	assert.Equal(t, 0, status)
	info, _ := fs.Stat("/test.txt")
	assert.Equal(t, os.FileMode(0744), info.Mode().Perm())

	// g-r
	status = c.Run(context.Background(), env, []string{"g-r", "/test.txt"})
	assert.Equal(t, 0, status)
	info, _ = fs.Stat("/test.txt")
	assert.Equal(t, os.FileMode(0704), info.Mode().Perm())

	// a+w
	status = c.Run(context.Background(), env, []string{"a+w", "/test.txt"})
	assert.Equal(t, 0, status)
	info, _ = fs.Stat("/test.txt")
	assert.Equal(t, os.FileMode(0726), info.Mode().Perm())

	// =r
	status = c.Run(context.Background(), env, []string{"=r", "/test.txt"})
	assert.Equal(t, 0, status)
	info, _ = fs.Stat("/test.txt")
	assert.Equal(t, os.FileMode(0444), info.Mode().Perm())
}

func TestChmod_Recursive(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, fs.MkdirAll("/dir/subdir", 0755))
	require.NoError(t, afero.WriteFile(fs, "/dir/file1", []byte("1"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/dir/subdir/file2", []byte("2"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: &bytes.Buffer{},
	}

	c := New()
	status := c.Run(context.Background(), env, []string{"-R", "777", "/dir"})
	assert.Equal(t, 0, status)

	info, _ := fs.Stat("/dir")
	assert.Equal(t, os.FileMode(0777), info.Mode().Perm())
	info, _ = fs.Stat("/dir/file1")
	assert.Equal(t, os.FileMode(0777), info.Mode().Perm())
	info, _ = fs.Stat("/dir/subdir/file2")
	assert.Equal(t, os.FileMode(0777), info.Mode().Perm())
}
