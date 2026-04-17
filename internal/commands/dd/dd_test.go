package dd

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDD_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/input.txt", []byte("0123456789ABCDEF"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdin:  io.NopCloser(strings.NewReader("")),
		Stdout: io.Discard,
		Stderr: io.Discard,
	}

	d := New("dd")

	// Test basic if/of
	status := d.Run(context.Background(), env, []string{"if=/input.txt", "of=/output.txt"})
	assert.Equal(t, 0, status)
	content, _ := afero.ReadFile(fs, "/output.txt")
	assert.Equal(t, "0123456789ABCDEF", string(content))

	// Test bs and count
	status = d.Run(context.Background(), env, []string{"if=/input.txt", "of=/output2.txt", "bs=4", "count=2"})
	assert.Equal(t, 0, status)
	content, _ = afero.ReadFile(fs, "/output2.txt")
	assert.Equal(t, "01234567", string(content))

	// Test skip
	status = d.Run(context.Background(), env, []string{"if=/input.txt", "of=/output3.txt", "bs=4", "skip=1", "count=1"})
	assert.Equal(t, 0, status)
	content, _ = afero.ReadFile(fs, "/output3.txt")
	assert.Equal(t, "4567", string(content))

	// Test seek
	require.NoError(t, afero.WriteFile(fs, "/output4.txt", []byte("XXXXXXXXXXXX"), 0644))
	status = d.Run(context.Background(), env, []string{"if=/input.txt", "of=/output4.txt", "bs=4", "seek=1", "count=1", "conv=notrunc"})
	assert.Equal(t, 0, status)
	content, _ = afero.ReadFile(fs, "/output4.txt")
	assert.Equal(t, "XXXX0123XXXX", string(content))
}
