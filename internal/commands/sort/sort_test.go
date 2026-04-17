package sortcmd

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

func TestSort_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("c\na\nb\n"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	s := New()
	status := s.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "a\nb\nc\n", env.Stdout.(*bytes.Buffer).String())

	// Test reverse
	env.Stdout = &bytes.Buffer{}
	status = s.Run(context.Background(), env, []string{"-r", "/test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "c\nb\na\n", env.Stdout.(*bytes.Buffer).String())
}
