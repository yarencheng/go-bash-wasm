package rmdir

import (
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
}
