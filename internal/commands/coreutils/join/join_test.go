package joincmd

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestJoin_Flags(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/f1.txt", []byte("A 1\nB 2\n"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/f2.txt", []byte("a X\nb Y\nc Z\n"), 0644))

	t.Run("IgnoreCase", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{FS: fs, Cwd: "/", Stdout: &out, Stderr: io.Discard}
		status := New().Run(context.Background(), env, []string{"-i", "/f1.txt", "/f2.txt"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "A 1 X")
		assert.Contains(t, out.String(), "B 2 Y")
	})

	t.Run("Unpairable1", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{FS: fs, Cwd: "/", Stdout: &out, Stderr: io.Discard}
		status := New().Run(context.Background(), env, []string{"-a1", "/f1.txt", "/f2.txt"})
		assert.Equal(t, 0, status)
		// A and B don't match exactly (case)
		assert.Contains(t, out.String(), "A 1")
		assert.Contains(t, out.String(), "B 2")
	})

	t.Run("Unpairable2", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{FS: fs, Cwd: "/", Stdout: &out, Stderr: io.Discard}
		status := New().Run(context.Background(), env, []string{"-a2", "/f1.txt", "/f2.txt"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "a X")
		assert.Contains(t, out.String(), "c Z")
	})

	t.Run("OutputFormat", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{FS: fs, Cwd: "/", Stdout: &out, Stderr: io.Discard}
		// -o 0 1.2 2.2 -> key, field 2 of file 1, field 2 of file 2
		status := New().Run(context.Background(), env, []string{"-i", "-o", "0 1.2 2.2", "/f1.txt", "/f2.txt"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "A 1 X")
		assert.Contains(t, out.String(), "B 2 Y")
	})

	t.Run("EmptyField", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{FS: fs, Cwd: "/", Stdout: &out, Stderr: io.Discard}
		// -a2 -e EMPTY -o 0 1.2 2.2
		status := New().Run(context.Background(), env, []string{"-i", "-a2", "-e", "EMPTY", "-o", "0 1.2 2.2", "/f1.txt", "/f2.txt"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "c EMPTY Z")
	})
}
