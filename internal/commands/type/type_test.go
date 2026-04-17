package typecmd

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/boolcmd"
)

func TestType_Run(t *testing.T) {
	registry := commands.New()
	_ = registry.Register(boolcmd.NewTrue()) // builtin "true"

	fs := afero.NewMemMapFs()
	_ = fs.MkdirAll("/usr/bin", 0755)
	_ = afero.WriteFile(fs, "/usr/bin/ls", []byte(""), 0755)

	env := &commands.Environment{
		FS:       fs,
		Registry: registry,
		Aliases:  map[string]string{"ll": "ls -l"},
		EnvVars:  map[string]string{"PATH": "/usr/bin"},
		Stdout:   &bytes.Buffer{},
		Stderr:   io.Discard,
	}

	cmd := New()

	// Test Alias
	status := cmd.Run(context.Background(), env, []string{"ll"})
	assert.Equal(t, 0, status)
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "ll is aliased to `ls -l'")
	env.Stdout.(*bytes.Buffer).Reset()

	// Test Builtin
	status = cmd.Run(context.Background(), env, []string{"true"})
	assert.Equal(t, 0, status)
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "true is a shell builtin")
	env.Stdout.(*bytes.Buffer).Reset()

	// Test File
	status = cmd.Run(context.Background(), env, []string{"ls"})
	assert.Equal(t, 0, status)
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "ls is /usr/bin/ls")
	env.Stdout.(*bytes.Buffer).Reset()

	// Test -t (type only)
	status = cmd.Run(context.Background(), env, []string{"-t", "ll", "true", "ls"})
	assert.Equal(t, 0, status)
	output := env.Stdout.(*bytes.Buffer).String()
	assert.Contains(t, output, "alias\n")
	assert.Contains(t, output, "builtin\n")
	assert.Contains(t, output, "file\n")
	env.Stdout.(*bytes.Buffer).Reset()

	// Test -p (path only)
	status = cmd.Run(context.Background(), env, []string{"-p", "ls"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/usr/bin/ls\n", env.Stdout.(*bytes.Buffer).String())
	env.Stdout.(*bytes.Buffer).Reset()

	// Test not found
	status = cmd.Run(context.Background(), env, []string{"nonexistent"})
	assert.Equal(t, 1, status)
}
