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

	// Test unique
	env.Stdout = &bytes.Buffer{}
	require.NoError(t, afero.WriteFile(fs, "/dup.txt", []byte("a\na\nb\n"), 0644))
	status = s.Run(context.Background(), env, []string{"-u", "/dup.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "a\nb\n", env.Stdout.(*bytes.Buffer).String())

	// Test ignore-case
	env.Stdout = &bytes.Buffer{}
	require.NoError(t, afero.WriteFile(fs, "/case.txt", []byte("B\na\n"), 0644))
	status = s.Run(context.Background(), env, []string{"-f", "/case.txt"})
	assert.Equal(t, 0, status)
	// Without -f, B < a. With -f, a and B are compared case-insensitively.
	// Bash sort -f: a comes before B typically? Depends on locale. 
	// Usually "a" and "B" compared as "A" vs "B".
	assert.Equal(t, "a\nB\n", env.Stdout.(*bytes.Buffer).String())

	// Test numeric
	env.Stdout = &bytes.Buffer{}
	require.NoError(t, afero.WriteFile(fs, "/num.txt", []byte("10\n2\n1\n"), 0644))
	status = s.Run(context.Background(), env, []string{"-n", "/num.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "1\n2\n10\n", env.Stdout.(*bytes.Buffer).String())

	// Test check
	status = s.Run(context.Background(), env, []string{"-c", "/test.txt"}) // /test.txt is c, a, b
	assert.Equal(t, 1, status) // Should fail check

	// Test output file
	status = s.Run(context.Background(), env, []string{"-o", "/out.txt", "/test.txt"})
	assert.Equal(t, 0, status)
	content, _ := afero.ReadFile(fs, "/out.txt")
	assert.Equal(t, "a\nb\nc\n", string(content))
}
