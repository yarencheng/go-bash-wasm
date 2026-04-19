package vdir

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestVdir_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/file.txt", []byte("content"), 0644)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: stdout,
		Stderr: stderr,
		Stdin:  io.NopCloser(bytes.NewReader(nil)),
	}

	v := New()
	status := v.Run(context.Background(), env, []string{"file.txt"})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "file.txt")
	// Since vdir is ls -l, it should have multiple columns
	assert.Contains(t, stdout.String(), "root") // owner
}

func TestVdir_Metadata(t *testing.T) {
	v := New()
	assert.Equal(t, "vdir", v.Name())
}
