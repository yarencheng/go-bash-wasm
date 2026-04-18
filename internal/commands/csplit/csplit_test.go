package csplit

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCsplit_Run_Regex(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "header\nline1\n---SEP---\nline2\n---SEP---\nline3\n"
	require.NoError(t, afero.WriteFile(fs, "/input.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdin:  io.NopCloser(nil),
		Stdout: io.Discard,
		Stderr: io.Discard,
	}

	c := New()

	// Split at regex /---SEP---/
	status := c.Run(context.Background(), env, []string{"-f", "out", "-n", "3", "/input.txt", "/---SEP---/", "/---SEP---/"})
	assert.Equal(t, 0, status)

	out0, _ := afero.ReadFile(fs, "/out000")
	assert.Equal(t, "header\nline1\n", string(out0))

	out1, _ := afero.ReadFile(fs, "/out001")
	assert.Equal(t, "---SEP---\nline2\n", string(out1))

	out2, _ := afero.ReadFile(fs, "/out002")
	assert.Equal(t, "---SEP---\nline3\n", string(out2))
}

func TestCsplit_SuppressMatched(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "header\nline1\n---SEP---\nline2\n"
	require.NoError(t, afero.WriteFile(fs, "/input.txt", []byte(content), 0644))

	out := &strings.Builder{}
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdin:  io.NopCloser(nil),
		Stdout: out,
		Stderr: io.Discard,
	}

	c := New()

	// Split at regex /---SEP---/ with suppress
	status := c.Run(context.Background(), env, []string{"--suppress-matched", "/input.txt", "/---SEP---/"})
	assert.Equal(t, 0, status)

	out0, _ := afero.ReadFile(fs, "/xx00")
	assert.Equal(t, "header\nline1\n", string(out0))

	out1, _ := afero.ReadFile(fs, "/xx01")
	assert.Equal(t, "line2\n", string(out1))
}

func TestCsplit_Run_Numeric(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "line1\nline2\nline3\nline4\nline5\n"
	require.NoError(t, afero.WriteFile(fs, "/input.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdin:  io.NopCloser(nil),
		Stdout: io.Discard,
		Stderr: io.Discard,
	}

	c := New()

	// Split at line 3
	status := c.Run(context.Background(), env, []string{"/input.txt", "3"})
	assert.Equal(t, 0, status)

	out0, _ := afero.ReadFile(fs, "/xx00")
	assert.Equal(t, "line1\nline2\n", string(out0))

	out1, _ := afero.ReadFile(fs, "/xx01")
	assert.Equal(t, "line3\nline4\nline5\n", string(out1))
}
