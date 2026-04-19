package fold

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

func TestFold_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "12345678901234567890\n" // 20 chars
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	// Fold to 5 chars
	status := cmd.Run(context.Background(), env, []string{"-w", "5", "/test.txt"})
	assert.Equal(t, 0, status)

	expected := "12345\n67890\n12345\n67890\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}

func TestFold_Spaces(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "123 456 789\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	// Fold to 5 chars, breaking at spaces
	status := cmd.Run(context.Background(), env, []string{"-w", "5", "-s", "/test.txt"})
	assert.Equal(t, 0, status)

	expected := "123 \n456 \n789\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}
