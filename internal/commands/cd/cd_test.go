package cd

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCd_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, fs.Mkdir("/testdir", 0755))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
		Stdout: io.Discard,
	}

	c := New()
	
	// Test basic cd
	status := c.Run(context.Background(), env, []string{"/testdir"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/testdir", env.Cwd)

	// Test relative cd
	status = c.Run(context.Background(), env, []string{".."})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/", env.Cwd)

	// Test cd to non-existent dir
	status = c.Run(context.Background(), env, []string{"/nonexistent"})
	assert.Equal(t, 1, status)
	assert.Equal(t, "/", env.Cwd)
}
