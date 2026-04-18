package nl

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

func TestNl_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "line 1\nline 2\n\nline 4\n"
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

	// Default nl numbers non-empty lines of the body with 6-character width and tab separator.
	// Bash nl default is more like -bt (number all non-empty lines)
	expected := "     1\tline 1\n     2\tline 2\n      \n     3\tline 4\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}

func TestNl_NumberAll(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "line 1\n\nline 3\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-ba", "/test.txt"})
	assert.Equal(t, 0, status)

	expected := "     1\tline 1\n     2\t\n     3\tline 3\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}

func TestNl_Sections(t *testing.T) {
	fs := afero.NewMemMapFs()
	// Default delim is \:, so header is \:\:\:, body is \:\:, footer is \:
	content := "\\:\\:\\:\nheader\n\\:\\:\nbody\n\\:\nfooter\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	// -ha -ba -fa -> number all lines in all sections
	status := cmd.Run(context.Background(), env, []string{"-ha", "-ba", "-fa", "/test.txt"})
	assert.Equal(t, 0, status)

	// Default is starting from 1 for each section
	output := env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "\n     1\theader\n")
	assert.Contains(t, output, "\n     1\tbody\n")
	assert.Contains(t, output, "\n     1\tfooter\n")
}

func TestNl_Increments(t *testing.T) {
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
	status := cmd.Run(context.Background(), env, []string{"-v", "10", "-i", "5", "/test.txt"})
	assert.Equal(t, 0, status)

	output := env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "    10\ta\n")
	assert.Contains(t, output, "    15\tb\n")
	assert.Contains(t, output, "    20\tc\n")
}
