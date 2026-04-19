package rm

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestRm_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/file.txt", []byte("data"), 0644))
	require.NoError(t, fs.Mkdir("/dir", 0755))
	require.NoError(t, afero.WriteFile(fs, "/dir/sub.txt", []byte("sub"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	r := New()

	// Test basic file removal
	status := r.Run(context.Background(), env, []string{"file.txt"})
	assert.Equal(t, 0, status)
	_, err := fs.Stat("/file.txt")
	assert.Error(t, err)

	// Test recursive dir removal failure without -r
	status = r.Run(context.Background(), env, []string{"dir"})
	assert.Equal(t, 1, status)
	_, err = fs.Stat("/dir")
	assert.NoError(t, err)

	// Test recursive dir removal success with -r
	status = r.Run(context.Background(), env, []string{"-r", "dir"})
	assert.Equal(t, 0, status)
	_, err = fs.Stat("/dir")
	assert.Error(t, err)
}
