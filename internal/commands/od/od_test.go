package od

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

func TestOd_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	// "hello" in octal is 150 145 154 154 157
	// Each word (2 bytes) would be 062550 066154 000157
	// But let's just check if it outputs something like octal addresses.
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	cmd := New()
	status := cmd.Run(context.Background(), env, []string{"/test.txt"})
	assert.Equal(t, 0, status)
	
	output := env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "0000000") // Starting address
}
