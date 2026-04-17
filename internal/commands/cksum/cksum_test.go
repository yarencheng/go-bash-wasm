package cksum

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCksum_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/test", []byte("hello"), 0644)
	
	out := &strings.Builder{}
	env := &commands.Environment{
		FS:     fs,
		Stdout: out,
	}

	c := New()
	status := c.Run(context.Background(), env, []string{"/test"})
	assert.Equal(t, 0, status)
	assert.Contains(t, out.String(), "5") // length
	assert.Contains(t, out.String(), "test")
}

func TestCksum_Stdin(t *testing.T) {
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdin:  io.NopCloser(strings.NewReader("hello")),
		Stdout: out,
	}

	c := New()
	status := c.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Contains(t, out.String(), "5")
}
