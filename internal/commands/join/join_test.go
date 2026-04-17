package joincmd

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

func TestJoin_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/file1.txt", []byte("1 a\n2 b\n"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/file2.txt", []byte("1 x\n2 y\n"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	j := New()
	status := j.Run(context.Background(), env, []string{"/file1.txt", "/file2.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "1 a x\n2 b y\n", env.Stdout.(*bytes.Buffer).String())
}
