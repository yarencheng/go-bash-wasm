package stat

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

func TestStat_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("data"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	s := New()

	// Test basic stat
	status := s.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)
	output := env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "File: /test.txt")
	assert.Contains(t, output, "Size: 4")
}
