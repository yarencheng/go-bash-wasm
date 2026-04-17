package cat

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCat_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("line1\nline2"), 0644))

	var stdout bytes.Buffer
	env := &commands.Environment{
		FS:     fs,
		Stdout: &stdout,
		Cwd:    "/",
	}

	c := New()
	status := c.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "line1\nline2", stdout.String())
}
