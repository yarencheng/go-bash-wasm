package cksum

import (
	"bytes"
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

func TestCksum_Hashing(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/file1.txt", []byte("hello world"), 0644)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	env := &commands.Environment{
		FS:     fs,
		Stdout: &stdout,
		Stderr: &stderr,
	}

	c := New()

	// Test MD5
	status := c.Run(context.Background(), env, []string{"-a", "md5", "/file1.txt"})
	assert.Equal(t, 0, status)
	// md5("hello world") = 5eb63bbbe01eeed093cb22bb8f5acdc3
	assert.Contains(t, stdout.String(), "5eb63bbbe01eeed093cb22bb8f5acdc3")

	// Test Check
	stdout.Reset()
	checkFile := "5eb63bbbe01eeed093cb22bb8f5acdc3  /file1.txt\n"
	afero.WriteFile(fs, "/check.md5", []byte(checkFile), 0644)
	status = c.Run(context.Background(), env, []string{"-c", "/check.md5"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/file1.txt: OK\n", stdout.String())
}
