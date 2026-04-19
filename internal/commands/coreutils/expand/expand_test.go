package expand

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

func TestExpand_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "\tline 1\n  \tline 2\n"
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

	// Default tab size is 8
	expected := "        line 1\n        line 2\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}

func TestExpand_Initial(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "\ta\tb\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	// Only initial tabs
	status := cmd.Run(context.Background(), env, []string{"-i", "/test.txt"})
	assert.Equal(t, 0, status)

	expected := "        a\tb\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}

func TestExpand_TabSize(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "\ta\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-t", "4", "/test.txt"})
	assert.Equal(t, 0, status)

	expected := "    a\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}
