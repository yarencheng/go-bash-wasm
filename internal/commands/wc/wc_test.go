package wc

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

func TestWc_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "hello world\nthis is a test\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	w := New()

	// Test default wc
	env.Stdout.(*bytes.Buffer).Reset()
	status := w.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)
	// Output format: lines words characters filename
	// lines: 2, words: 6, bytes: 27
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "2")
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "6")
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "27")
}
