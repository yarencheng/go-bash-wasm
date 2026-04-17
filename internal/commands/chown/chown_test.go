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
}
