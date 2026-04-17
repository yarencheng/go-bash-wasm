package mv

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestMv_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/source.txt", []byte("content"), 0644))
	require.NoError(t, fs.Mkdir("/dir", 0755))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	m := New()

	// Test basic rename
	status := m.Run(context.Background(), env, []string{"source.txt", "dest.txt"})
	assert.Equal(t, 0, status)
	_, err := fs.Stat("/source.txt")
	assert.Error(t, err)
	content, err := afero.ReadFile(fs, "/dest.txt")
	require.NoError(t, err)
	assert.Equal(t, "content", string(content))

	// Test move to directory
	require.NoError(t, afero.WriteFile(fs, "/new.txt", []byte("new"), 0644))
	status = m.Run(context.Background(), env, []string{"new.txt", "dir/"})
	assert.Equal(t, 0, status)
	_, err = fs.Stat("/new.txt")
	assert.Error(t, err)
	content, err = afero.ReadFile(fs, "/dir/new.txt")
	require.NoError(t, err)
	assert.Equal(t, "new", string(content))
}
