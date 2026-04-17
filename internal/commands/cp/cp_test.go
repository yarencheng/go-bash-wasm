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

	// Test copy to directory
	status = c.Run(context.Background(), env, []string{"source.txt", "dir/"})
	assert.Equal(t, 0, status)
	content, err = afero.ReadFile(fs, "/dir/source.txt")
	require.NoError(t, err)
	assert.Equal(t, "content", string(content))
}
