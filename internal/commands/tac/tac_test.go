package tac

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

func TestTac_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "line 1\nline 2\nline 3\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)

	expected := "line 3\nline 2\nline 1\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}
