package readlink

import (
	"context"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestReadlink_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	linker, ok := fs.(afero.Symlinker)
	if !ok {
		t.Skip("FS does not support symlinks")
	}
	_ = afero.WriteFile(fs, "/target.txt", []byte("hello"), 0644)
	err := linker.SymlinkIfPossible("/target.txt", "/link.txt")
	if err != nil {
		t.Skip("Symlink failed on this FS")
	}

	out := &strings.Builder{}
	env := &commands.Environment{
		FS:     fs,
		Stdout: out,
		Cwd:    "/",
	}

	r := New()
	status := r.Run(context.Background(), env, []string{"/link.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/target.txt\n", out.String())
}

func TestReadlink_NoNewline(t *testing.T) {
	fs := afero.NewMemMapFs()
	linker, ok := fs.(afero.Symlinker)
	if !ok {
		t.Skip("FS does not support symlinks")
	}
	err := linker.SymlinkIfPossible("/target", "/link")
	if err != nil {
		t.Skip("Symlink failed on this FS")
	}

	out := &strings.Builder{}
	env := &commands.Environment{
		FS:     fs,
		Stdout: out,
		Cwd:    "/",
	}

	r := New()
	status := r.Run(context.Background(), env, []string{"-n", "/link"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/target", out.String())
}
