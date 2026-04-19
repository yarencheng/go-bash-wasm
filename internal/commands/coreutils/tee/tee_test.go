package tee

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTee_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	var stdout bytes.Buffer
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdin:  io.NopCloser(bytes.NewBufferString("hello")),
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	te := New()
	status := te.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello", stdout.String())

	content, _ := afero.ReadFile(fs, "/test.txt")
	assert.Equal(t, "hello", string(content))
}
