package paste

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

func TestPaste_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/f1.txt", []byte("1\n2\n"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/f2.txt", []byte("a\nb\nc\n"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"/f1.txt", "/f2.txt"})
	assert.Equal(t, 0, status)
	
	expected := "1\ta\n2\tb\n\tc\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}

func TestPaste_Delimiter(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/f1.txt", []byte("1\n2\n"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/f2.txt", []byte("a\nb\n"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"-d", ",", "/f1.txt", "/f2.txt"})
	assert.Equal(t, 0, status)
	
	expected := "1,a\n2,b\n"
	assert.Equal(t, expected, env.Stdout.(*bytes.Buffer).String())
}
