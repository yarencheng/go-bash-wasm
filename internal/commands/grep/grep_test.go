package grep

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

func TestGrep_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello\nworld\nHELLO\n"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	g := New()

	// Test case-sensitive
	status := g.Run(context.Background(), env, []string{"hello", "/test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello\n", env.Stdout.(*bytes.Buffer).String())

	// Test case-insensitive
	env.Stdout = &bytes.Buffer{}
	status = g.Run(context.Background(), env, []string{"-i", "hello", "/test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello\nHELLO\n", env.Stdout.(*bytes.Buffer).String())

	// Test invert match
	env.Stdout = &bytes.Buffer{}
	status = g.Run(context.Background(), env, []string{"-v", "hello", "/test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "world\nHELLO\n", env.Stdout.(*bytes.Buffer).String())
}
