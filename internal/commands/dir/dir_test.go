package dir

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDir_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/file1.txt", []byte(""), 0644)
	_ = afero.WriteFile(fs, "/file2.txt", []byte(""), 0644)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: stdout,
		Stderr: stderr,
		Stdin:  io.NopCloser(bytes.NewReader(nil)),
	}

	d := New()
	status := d.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "file1.txt")
	assert.Contains(t, stdout.String(), "file2.txt")
}

func TestDir_Metadata(t *testing.T) {
	d := New()
	assert.Equal(t, "dir", d.Name())
}
