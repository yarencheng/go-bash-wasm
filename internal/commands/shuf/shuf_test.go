package shuf

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

func TestShuf_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "a\nb\nc\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)

	output := env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "a")
	assert.Contains(t, output, "b")
	assert.Contains(t, output, "c")
}

func TestShuf_Echo(t *testing.T) {
	env := &commands.Environment{
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-e", "a", "b", "c"})
	assert.Equal(t, 0, status)

	output := env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "a")
	assert.Contains(t, output, "b")
	assert.Contains(t, output, "c")
}

func TestShuf_Range(t *testing.T) {
	env := &commands.Environment{
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-i", "1-3"})
	assert.Equal(t, 0, status)

	output := env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "1")
	assert.Contains(t, output, "2")
	assert.Contains(t, output, "3")
}

func TestShuf_Zero(t *testing.T) {
	env := &commands.Environment{
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-z", "-e", "a", "b"})
	assert.Equal(t, 0, status)

	output := env.Stdout.(*bytes.Buffer).String()
	assert.True(t, output == "a\x00b\x00" || output == "b\x00a\x00")
}
