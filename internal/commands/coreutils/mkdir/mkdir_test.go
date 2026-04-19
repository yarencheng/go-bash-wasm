package mkdir

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

	// Test verbose
	stdout := &bytes.Buffer{}
	env.Stdout = stdout
	status = m.Run(context.Background(), env, []string{"-v", "vdir"})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "created directory 'vdir'")

	// Test help
	stdout.Reset()
	status = m.Run(context.Background(), env, []string{"--help"})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "Usage: mkdir")

	// Test version
	stdout.Reset()
	status = m.Run(context.Background(), env, []string{"--version"})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "mkdir (go-bash-wasm)")

	// Test invalid mode
	status = m.Run(context.Background(), env, []string{"-m", "invalid", "errdir"})
	assert.Equal(t, 1, status)

	// Test missing operand
	status = m.Run(context.Background(), env, []string{})
	assert.Equal(t, 1, status)
}

func TestMkdir_Metadata(t *testing.T) {
	m := New()
	assert.Equal(t, "mkdir", m.Name())
}
