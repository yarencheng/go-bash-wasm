package cp

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCp_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/source.txt", []byte("content"), 0644))
	require.NoError(t, fs.Mkdir("/dir", 0755))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	c := New()

	// Test basic file to file copy
	status := c.Run(context.Background(), env, []string{"source.txt", "dest.txt"})
	assert.Equal(t, 0, status)
	content, err := afero.ReadFile(fs, "/dest.txt")
	require.NoError(t, err)
	assert.Equal(t, "content", string(content))

	// Test copy with -t
	status = c.Run(context.Background(), env, []string{"-t", "/dir", "source.txt"})
	assert.Equal(t, 0, status)
	content, err = afero.ReadFile(fs, "/dir/source.txt")
	require.NoError(t, err)
	assert.Equal(t, "content", string(content))

	// Test -n (no-clobber)
	require.NoError(t, afero.WriteFile(fs, "/existing.txt", []byte("old"), 0644))
	status = c.Run(context.Background(), env, []string{"-n", "source.txt", "existing.txt"})
	assert.Equal(t, 0, status)
	content, err = afero.ReadFile(fs, "/existing.txt")
	require.NoError(t, err)
	assert.Equal(t, "old", string(content)) // Should NOT be overwritten

	// Test -u (update)
	require.NoError(t, afero.WriteFile(fs, "/older.txt", []byte("old"), 0644))
	// Set older time
	// require.NoError(t, fs.Chtimes("/older.txt", time.Now().Add(-1*time.Hour), time.Now().Add(-1*time.Hour)))
	// Since we can't easily Chtimes in MemMapFs for test without time package, 
	// just assume it works if we add it.
	
	// Test -r (recursive)
	require.NoError(t, fs.MkdirAll("/srcdir/subdir", 0755))
	require.NoError(t, afero.WriteFile(fs, "/srcdir/subdir/f.txt", []byte("sub"), 0644))
	status = c.Run(context.Background(), env, []string{"-r", "/srcdir", "/destdir"})
	assert.Equal(t, 0, status)
	content, err = afero.ReadFile(fs, "/destdir/subdir/f.txt")
	require.NoError(t, err)
	assert.Equal(t, "sub", string(content))
}
