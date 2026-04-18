package du

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

func TestDu_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello"), 0644))
	require.NoError(t, fs.Mkdir("/dir", 0755))
	require.NoError(t, afero.WriteFile(fs, "/dir/file.txt", []byte("world"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	d := New()
	status := d.Run(context.Background(), env, []string{"/dir"})
	assert.Equal(t, 0, status)
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "/dir")
}

func TestDu_Flags(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/file1.txt", make([]byte, 2000), 0644))
	require.NoError(t, afero.WriteFile(fs, "/file2.txt", make([]byte, 5000), 0644))
	
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	d := New()
	
	t.Run("Threshold", func(t *testing.T) {
		var out bytes.Buffer
		env.Stdout = &out
		status := d.Run(context.Background(), env, []string{"-t", "3000", "/file1.txt", "/file2.txt"})
		assert.Equal(t, 0, status)
		assert.NotContains(t, out.String(), "file1.txt")
		assert.Contains(t, out.String(), "file2.txt")
	})

	t.Run("SI", func(t *testing.T) {
		var out bytes.Buffer
		env.Stdout = &out
		status := d.Run(context.Background(), env, []string{"--si", "/file2.txt"})
		assert.Equal(t, 0, status)
		// 5000 bytes in SI is 5.0k (1000 division)
		assert.Contains(t, out.String(), "5.0K")
	})

	t.Run("Inodes", func(t *testing.T) {
		var out bytes.Buffer
		env.Stdout = &out
		status := d.Run(context.Background(), env, []string{"--inodes", "/file1.txt"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "1\t/file1.txt")
	})
}
