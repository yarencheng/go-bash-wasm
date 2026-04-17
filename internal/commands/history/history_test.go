package history

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

func TestHistory_Run(t *testing.T) {
	stdout := &bytes.Buffer{}
	env := &commands.Environment{
		FS:      afero.NewMemMapFs(),
		History: []string{"ls", "cd /tmp", "echo hello"},
		EnvVars: map[string]string{"HISTFILE": "/.history"},
		Stdout:  stdout,
		Stderr:  io.Discard,
	}

	cmd := New()

	// Test display
	status := cmd.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "1  ls")
	assert.Contains(t, stdout.String(), "2  cd /tmp")
	assert.Contains(t, stdout.String(), "3  echo hello")
	stdout.Reset()

	// Test -d (delete)
	status = cmd.Run(context.Background(), env, []string{"-d", "2"})
	assert.Equal(t, 0, status)
	assert.Equal(t, []string{"ls", "echo hello"}, env.History)

	// Test -w (write)
	status = cmd.Run(context.Background(), env, []string{"-w"})
	assert.Equal(t, 0, status)
	data, err := afero.ReadFile(env.FS, "/.history")
	require.NoError(t, err)
	assert.Equal(t, "ls\necho hello\n", string(data))

	// Test -c (clear)
	status = cmd.Run(context.Background(), env, []string{"-c"})
	assert.Equal(t, 0, status)
	assert.Empty(t, env.History)

	// Test -r (read)
	status = cmd.Run(context.Background(), env, []string{"-r"})
	assert.Equal(t, 0, status)
	assert.Equal(t, []string{"ls", "echo hello"}, env.History)
}
