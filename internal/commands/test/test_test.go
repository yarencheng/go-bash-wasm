package testcmd

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTest_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	tc := New("test")
	
	// Test -e (exists)
	status := tc.Run(context.Background(), env, []string{"-e", "/test.txt"})
	assert.Equal(t, 0, status)

	// Test -d (directory)
	status = tc.Run(context.Background(), env, []string{"-d", "/test.txt"})
	assert.Equal(t, 1, status)

	// Test string
	status = tc.Run(context.Background(), env, []string{"-z", ""})
	assert.Equal(t, 0, status)
}
