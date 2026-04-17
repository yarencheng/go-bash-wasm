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
