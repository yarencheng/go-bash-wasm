package mkdir

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestMkdir_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	m := New()

	// Test basic mkdir
	status := m.Run(context.Background(), env, []string{"testdir"})
	assert.Equal(t, 0, status)
	info, err := fs.Stat("/testdir")
	require.NoError(t, err)
	assert.True(t, info.IsDir())

	// Test -p prefix
	status = m.Run(context.Background(), env, []string{"-p", "parent/child"})
	assert.Equal(t, 0, status)
	info, err = fs.Stat("/parent/child")
	require.NoError(t, err)
	assert.True(t, info.IsDir())

	// Test -m flag
	status = m.Run(context.Background(), env, []string{"-m", "0700", "modedir"})
	assert.Equal(t, 0, status)
	info, err = fs.Stat("/modedir")
	require.NoError(t, err)
	assert.Equal(t, uint32(0700), uint32(info.Mode().Perm()))
}
