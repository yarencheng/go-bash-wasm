package mktemp

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestMktemp_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	out := &strings.Builder{}
	env := &commands.Environment{
		FS:     fs,
		Stdout: out,
		Cwd:    "/",
	}

	m := New()
	status := m.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)

	path := strings.TrimSpace(out.String())
	assert.Contains(t, path, "tmp.")

	info, err := fs.Stat(path)
	assert.NoError(t, err)
	assert.False(t, info.IsDir())
}

func TestMktemp_Directory(t *testing.T) {
	fs := afero.NewMemMapFs()
	out := &strings.Builder{}
	env := &commands.Environment{
		FS:     fs,
		Stdout: out,
		Cwd:    "/",
	}

	m := New()
	status := m.Run(context.Background(), env, []string{"-d"})
	assert.Equal(t, 0, status)

	path := strings.TrimSpace(out.String())
	info, err := fs.Stat(path)
	assert.NoError(t, err)
	assert.True(t, info.IsDir())
}

func TestMktemp_Suffix(t *testing.T) {
	fs := afero.NewMemMapFs()
	out := &strings.Builder{}
	env := &commands.Environment{
		FS:     fs,
		Stdout: out,
		Cwd:    "/",
	}

	m := New()
	// mktemp test.XXXXXX
	status := m.Run(context.Background(), env, []string{"test.XXXXXX"})
	assert.Equal(t, 0, status)

	path := strings.TrimSpace(out.String())
	assert.True(t, strings.HasPrefix(filepath.Base(path), "test."))
}
