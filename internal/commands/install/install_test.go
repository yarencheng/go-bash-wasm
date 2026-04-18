package install

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestInstall_Directory(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	i := New()
	status := i.Run(context.Background(), env, []string{"-d", "/new/dir"})
	assert.Equal(t, 0, status)

	info, err := fs.Stat("/new/dir")
	assert.NoError(t, err)
	assert.True(t, info.IsDir())
}

func TestInstall_File(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/src.txt", []byte("hello"), 0644)
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	i := New()
	status := i.Run(context.Background(), env, []string{"/src.txt", "/dst.txt"})
	assert.Equal(t, 0, status)

	content, _ := afero.ReadFile(fs, "/dst.txt")
	assert.Equal(t, "hello", string(content))
}
