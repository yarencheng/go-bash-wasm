package split

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

func TestSplit_BasicLines(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "line 1\nline 2\nline 3\nline 4\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	// Default split is 1000 lines. Let's split by 2 lines.
	status := cmd.Run(context.Background(), env, []string{"-l", "2", "/test.txt", "x"})
	assert.Equal(t, 0, status)
	
	f1, err := afero.ReadFile(fs, "/xaa")
	require.NoError(t, err)
	assert.Equal(t, "line 1\nline 2\n", string(f1))

	f2, err := afero.ReadFile(fs, "/xab")
	require.NoError(t, err)
	assert.Equal(t, "line 3\nline 4\n", string(f2))
}

func TestSplit_BasicBytes(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "abcdefghij"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-b", "3", "/test.txt", "y"})
	assert.Equal(t, 0, status)
	
	f1, _ := afero.ReadFile(fs, "/yaa")
	assert.Equal(t, "abc", string(f1))
	f2, _ := afero.ReadFile(fs, "/yab")
	assert.Equal(t, "def", string(f2))
	f3, _ := afero.ReadFile(fs, "/yac")
	assert.Equal(t, "ghi", string(f3))
	f4, _ := afero.ReadFile(fs, "/yad")
	assert.Equal(t, "j", string(f4))
}
