package mv

import (
	"context"
	"io"
	"testing"
	"time"

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

func TestMv_NoClobber(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/src", []byte("src"), 0644)
	_ = afero.WriteFile(fs, "/dest", []byte("dest"), 0644)

	env := &commands.Environment{FS: fs, Cwd: "/", Stderr: io.Discard}
	m := New()

	// -n should not overwrite /dest
	status := m.Run(context.Background(), env, []string{"-n", "/src", "/dest"})
	assert.Equal(t, 0, status)
	content, _ := afero.ReadFile(fs, "/dest")
	assert.Equal(t, "dest", string(content))
	_, err := fs.Stat("/src")
	assert.NoError(t, err) // /src should still exist
}

func TestMv_Update(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/newer", []byte("newer"), 0644)
	_ = afero.WriteFile(fs, "/older", []byte("older"), 0644)

	// Artificially make /older older than /newer
	now := time.Now()
	_ = fs.Chtimes("/older", now.Add(-time.Hour), now.Add(-time.Hour))
	_ = fs.Chtimes("/newer", now, now)

	env := &commands.Environment{FS: fs, Cwd: "/", Stderr: io.Discard}
	m := New()

	// mv -u older newer -> should not overwrite newer
	status := m.Run(context.Background(), env, []string{"-u", "/older", "/newer"})
	assert.Equal(t, 0, status)
	content, _ := afero.ReadFile(fs, "/newer")
	assert.Equal(t, "newer", string(content))

	// mv -u newer older -> should overwrite older
	status = m.Run(context.Background(), env, []string{"-u", "/newer", "/older"})
	assert.Equal(t, 0, status)
	content, _ = afero.ReadFile(fs, "/older")
	assert.Equal(t, "newer", string(content))
}
