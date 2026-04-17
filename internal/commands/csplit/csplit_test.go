package csplit

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCsplit_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "line1\nline2\nline3\nline4\nline5\n"
	require.NoError(t, afero.WriteFile(fs, "/input.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdin:  io.NopCloser(nil),
		Stdout: io.Discard,
		Stderr: io.Discard,
	}

	c := New()

	// Split at line 3
	status := c.Run(context.Background(), env, []string{"/input.txt", "3"})
	assert.Equal(t, 0, status)

	out0, _ := afero.ReadFile(fs, "/xx00")
	assert.Equal(t, "line1\nline2\n", string(out0))

	out1, _ := afero.ReadFile(fs, "/xx01")
	assert.Equal(t, "line3\nline4\nline5\n", string(out1))
}
