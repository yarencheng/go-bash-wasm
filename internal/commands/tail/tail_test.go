package tail

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

func TestTail_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "line1\nline2\nline3\nline4\nline5\nline6\nline7\nline8\nline9\nline10\nline11\n"
	require.NoError(t, afero.WriteFile(fs, "/large.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	tr := New()

	// Test default tail (10 lines)
	env.Stdout.(*bytes.Buffer).Reset()
	status := tr.Run(context.Background(), env, []string{"/large.txt"})
	assert.Equal(t, 0, status)
	expected := "line2\nline3\nline4\nline5\nline6\nline7\nline8\nline9\nline10\nline11\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())

	// Test tail -n 2
	env.Stdout.(*bytes.Buffer).Reset()
	status = tr.Run(context.Background(), env, []string{"-n", "2", "/large.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "line10\nline11\n", env.Stdout.(*bytes.Buffer).String())
}
