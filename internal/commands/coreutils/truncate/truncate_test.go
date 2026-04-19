package truncate

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTruncate_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello world"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-s", "5", "/test.txt"})
	assert.Equal(t, 0, status)

	f, err := afero.ReadFile(fs, "/test.txt")
	require.NoError(t, err)
	assert.Equal(t, 5, len(f))
	assert.Equal(t, "hello", string(f))
}

func TestTruncate_Create(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-s", "10", "/new.txt"})
	assert.Equal(t, 0, status)

	f, err := afero.ReadFile(fs, "/new.txt")
	require.NoError(t, err)
	assert.Equal(t, 10, len(f))
}

func TestTruncate_Reference(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/ref.txt", []byte("12345678"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-r", "/ref.txt", "/target.txt"})
	assert.Equal(t, 0, status)

	f, err := afero.ReadFile(fs, "/target.txt")
	require.NoError(t, err)
	assert.Equal(t, 8, len(f))
}

func TestTruncate_Relative(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("12345"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	cmd := New()
	// Add 3 bytes
	status := cmd.Run(context.Background(), env, []string{"-s", "+3", "/test.txt"})
	assert.Equal(t, 0, status)
	f, _ := afero.ReadFile(fs, "/test.txt")
	assert.Equal(t, 8, len(f))

	// Subtract 2 bytes
	status = cmd.Run(context.Background(), env, []string{"-s", "-2", "/test.txt"})
	assert.Equal(t, 0, status)
	f, _ = afero.ReadFile(fs, "/test.txt")
	assert.Equal(t, 6, len(f))
	
	// Subtract more than size
	status = cmd.Run(context.Background(), env, []string{"-s", "-10", "/test.txt"})
	assert.Equal(t, 0, status)
	f, _ = afero.ReadFile(fs, "/test.txt")
	assert.Equal(t, 0, len(f))
}

func TestTruncate_Multipliers(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-s", "1K", "/k.txt"})
	assert.Equal(t, 0, status)
	f, _ := afero.ReadFile(fs, "/k.txt")
	assert.Equal(t, 1024, len(f))

	status = cmd.Run(context.Background(), env, []string{"-s", "1M", "/m.txt"})
	assert.Equal(t, 0, status)
	// Afero MemMapFs might be slow/large but 1MB is fine.
	stat, _ := fs.Stat("/m.txt")
	assert.Equal(t, int64(1024*1024), stat.Size())
}

func TestTruncate_NoCreate(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-c", "-s", "10", "/nonexistent.txt"})
	assert.Equal(t, 0, status)
	exists, _ := afero.Exists(fs, "/nonexistent.txt")
	assert.False(t, exists)
}

func TestTruncate_Errors(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	cmd := New()
	// Missing file
	assert.Equal(t, 1, cmd.Run(context.Background(), env, []string{"-s", "10"}))
	// Missing reference
	assert.Equal(t, 1, cmd.Run(context.Background(), env, []string{"-r", "missing", "target"}))
	// Invalid size
	assert.Equal(t, 1, cmd.Run(context.Background(), env, []string{"-s", "invalid", "target"}))
}

func TestTruncate_Metadata(t *testing.T) {
	cmd := New()
	assert.Equal(t, "truncate", cmd.Name())
}
