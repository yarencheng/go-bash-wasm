package fmt

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

func TestFmt_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "This is a long line that should be wrapped to a smaller width to see if it works correctly.\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	// Wrap to 20 chars
	status := cmd.Run(context.Background(), env, []string{"-w", "20", "/test.txt"})
	assert.Equal(t, 0, status)
	
	// Check if any line exceeds 20 chars
	lines := bytes.Split(env.Stdout.(*bytes.Buffer).Bytes(), []byte("\n"))
	for _, line := range lines {
		if len(line) > 0 {
			assert.LessOrEqual(t, len(line), 20)
		}
	}
}

func TestFmt_SplitOnly(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "Short line.\nAnother short line.\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	// Split only should not join short lines
	status := cmd.Run(context.Background(), env, []string{"-s", "/test.txt"})
	assert.Equal(t, 0, status)
	
	assert.Equal(t, content, env.Stdout.(*bytes.Buffer).String())
}
