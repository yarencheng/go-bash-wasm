package realpath

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestRealpath_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, fs.MkdirAll("/data/logs", 0755))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/data",
		Stdout: &bytes.Buffer{},
		Stderr: &bytes.Buffer{},
	}

	r := New()

	// Test basic
	env.Stdout.(*bytes.Buffer).Reset()
	status := r.Run(context.Background(), env, []string{"logs/../logs"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/data/logs\n", env.Stdout.(*bytes.Buffer).String())

	// Test -z
	env.Stdout.(*bytes.Buffer).Reset()
	status = r.Run(context.Background(), env, []string{"-z", "logs"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/data/logs\x00", env.Stdout.(*bytes.Buffer).String())

	// Test -e (existing)
	status = r.Run(context.Background(), env, []string{"-e", "nonexistent"})
	assert.Equal(t, 1, status)

	// Test --relative-to
	env.Stdout.(*bytes.Buffer).Reset()
	status = r.Run(context.Background(), env, []string{"--relative-to=/data", "/data/logs"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "logs\n", env.Stdout.(*bytes.Buffer).String())

	// Test --relative-base
	env.Stdout.(*bytes.Buffer).Reset()
	status = r.Run(context.Background(), env, []string{"--relative-base=/data", "/data/logs", "/etc/passwd"})
	assert.Equal(t, 0, status)
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "logs\n")
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "/etc/passwd\n")
}
