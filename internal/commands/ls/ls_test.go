package ls

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestLs_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	ls := New()

	// Setup mock filesystem
	require.NoError(t, afero.WriteFile(fs, "/file1.txt", []byte("content1"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/file2.txt", []byte("content2"), 0644))
	require.NoError(t, fs.Mkdir("/dir1", 0755))
	require.NoError(t, afero.WriteFile(fs, "/.hidden", []byte("hidden"), 0644))

	t.Run("default listing-alphabetical", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{})
		assert.Equal(t, 0, status)
		// We'll implement basic space separation first.
		assert.Contains(t, stdout.String(), "dir1")
		assert.Contains(t, stdout.String(), "file1.txt")
		assert.Contains(t, stdout.String(), "file2.txt")
		assert.NotContains(t, stdout.String(), ".hidden")
	})

	t.Run("all files -a", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-a"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), ".")
		assert.Contains(t, stdout.String(), "..")
		assert.Contains(t, stdout.String(), ".hidden")
		assert.Contains(t, stdout.String(), "file1.txt")
	})

	t.Run("long listing -l", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-l"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "drwxr-xr-x") // directory1
		assert.Contains(t, stdout.String(), "-rw-r--r--") // file1.txt
		assert.Contains(t, stdout.String(), "8")          // size of file1.txt
	})

	t.Run("indicators -F", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-F"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "dir1/")
		assert.Contains(t, stdout.String(), "file1.txt")
	})

	t.Run("sort by size -S", func(t *testing.T) {
		// file1 and file2 are same size in my setup, let's create a bigger one
		require.NoError(t, afero.WriteFile(fs, "/big.txt", make([]byte, 100), 0644))
		
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-S"})
		assert.Equal(t, 0, status)
		// big.txt should be first
		assert.True(t, strings.HasPrefix(stdout.String(), "big.txt"))
	})

	t.Run("invalid directory", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/nonexistent",
		}

		status := ls.Run(context.Background(), env, []string{})
		assert.NotEqual(t, 0, status)
		assert.Contains(t, stderr.String(), "exist")
	})
}
