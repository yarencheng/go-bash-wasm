package base64cmd

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

func TestBase64_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	b := New()

	// Test encoding
	status := b.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "aGVsbG8=\n", env.Stdout.(*bytes.Buffer).String())

	// Test decoding
	env.Stdout = &bytes.Buffer{}
	require.NoError(t, afero.WriteFile(fs, "/encoded.txt", []byte("aGVsbG8="), 0644))
	status = b.Run(context.Background(), env, []string{"-d", "/encoded.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello", env.Stdout.(*bytes.Buffer).String())
}
