package sumlegacy

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestSumLegacy(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/test", []byte("hello world\n"), 0644)
	
	env := &commands.Environment{
		FS:     fs,
		Stdout: new(bytes.Buffer),
		Stderr: new(bytes.Buffer),
		Cwd:    "/",
	}
	
	s := New()

	t.Run("BSD sum", func(t *testing.T) {
		env.Stdout.(*bytes.Buffer).Reset()
		code := s.Run(context.Background(), env, []string{"/test"})
		assert.Equal(t, 0, code)
		assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "03762     1 /test")
	})

	t.Run("SysV sum", func(t *testing.T) {
		env.Stdout.(*bytes.Buffer).Reset()
		code := s.Run(context.Background(), env, []string{"-s", "/test"})
		assert.Equal(t, 0, code)
		assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "1126 1 /test")
	})
}
