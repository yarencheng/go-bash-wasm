package cut

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

func TestCut_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("a,b,c\n1,2,3\n"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	c := New()
	status := c.Run(context.Background(), env, []string{"-d", ",", "-f", "2", "/test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "b\n2\n", env.Stdout.(*bytes.Buffer).String())
}
