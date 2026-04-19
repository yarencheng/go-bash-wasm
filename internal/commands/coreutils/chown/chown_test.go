package chown

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestChown_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("test"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: &bytes.Buffer{},
	}

	c := New()

	// Test chown wasm:wasm
	status := c.Run(context.Background(), env, []string{"wasm:wasm", "/test.txt"})
	assert.Equal(t, 0, status)

	// Test chown root
	status = c.Run(context.Background(), env, []string{"root", "/test.txt"})
	assert.Equal(t, 0, status)

	// Test chown :1001
	status = c.Run(context.Background(), env, []string{":1001", "/test.txt"})
	assert.Equal(t, 0, status)

	t.Run("recursive chown", func(t *testing.T) {
		require.NoError(t, fs.Mkdir("/recursive", 0755))
		require.NoError(t, afero.WriteFile(fs, "/recursive/f1", []byte(""), 0644))
		status := c.Run(context.Background(), env, []string{"-R", "1000:1000", "/recursive"})
		assert.Equal(t, 0, status)
	})

	t.Run("verbose output", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		env.Stdout = stdout
		status := c.Run(context.Background(), env, []string{"-v", "root", "/test.txt"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "ownership of '/test.txt' retained or changed")
	})

	t.Run("reference flag", func(t *testing.T) {
		status := c.Run(context.Background(), env, []string{"--reference=/test.txt", "/recursive/f1"})
		assert.Equal(t, 0, status)
	})

	t.Run("numeric IDs and dot syntax", func(t *testing.T) {
		status := c.Run(context.Background(), env, []string{"1001.1001", "/test.txt"})
		assert.Equal(t, 0, status)
	})

	t.Run("errors", func(t *testing.T) {
		stderr := &bytes.Buffer{}
		env.Stderr = stderr
		// Missing operand
		assert.Equal(t, 1, c.Run(context.Background(), env, []string{"root"}))
		assert.Contains(t, stderr.String(), "missing operand")

		// File not found
		stderr.Reset()
		assert.Equal(t, 1, c.Run(context.Background(), env, []string{"root", "/nonexistent"}))
		assert.Contains(t, stderr.String(), "cannot access")
	})
}

func TestChgrp_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("test"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: &bytes.Buffer{},
	}

	c := NewChgrp()
	status := c.Run(context.Background(), env, []string{"wasm", "/test.txt"})
	assert.Equal(t, 0, status)

	t.Run("numeric gid", func(t *testing.T) {
		status := c.Run(context.Background(), env, []string{"1002", "/test.txt"})
		assert.Equal(t, 0, status)
	})
}
