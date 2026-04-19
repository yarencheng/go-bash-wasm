package pr

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPr_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	content := "line 1\nline 2\n"
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(content), 0644))

	env := &commands.Environment{
		FS:        fs,
		Cwd:       "/",
		Stdout:    &bytes.Buffer{},
		Stderr:    io.Discard,
		StartTime: time.Unix(1000000000, 0), // Constant time for testing
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)

	output := env.Stdout.(*bytes.Buffer).String()
	// Default pr has header with date, filename, and page number.
	// We'll check if it contains the filename and line 1.
	assert.Contains(t, output, "/test.txt")
	assert.Contains(t, output, "line 1")
}
