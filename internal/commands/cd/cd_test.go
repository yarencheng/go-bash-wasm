package cd

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCd_Features(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, fs.Mkdir("/home", 0755))
	require.NoError(t, fs.Mkdir("/home/user", 0755))
	require.NoError(t, fs.Mkdir("/work", 0755))
	require.NoError(t, fs.Mkdir("/work/proj", 0755))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
		Stdout: io.Discard,
		EnvVars: map[string]string{
			"HOME":   "/home/user",
			"CDPATH": ".:/work",
		},
	}

	c := New()

	// 1. Test cd (no args) -> $HOME
	status := c.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/home/user", env.Cwd)

	// 2. Test cd - -> $OLDPWD
	env.EnvVars["OLDPWD"] = "/work"
	status = c.Run(context.Background(), env, []string{"-"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/work", env.Cwd)
	assert.Equal(t, "/home/user", env.EnvVars["OLDPWD"])

	// 3. Test CDPATH
	env.Cwd = "/"
	status = c.Run(context.Background(), env, []string{"proj"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/work/proj", env.Cwd)
}
