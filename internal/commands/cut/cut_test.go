package cut

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

func TestCut_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("a,b,c\n1,2,3\n"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: nil,
		Stderr: io.Discard,
	}

	c := New()
	t.Run("cut fields", func(t *testing.T) {
		var stdout bytes.Buffer
		env.Stdout = &stdout
		status := c.Run(context.Background(), env, []string{"-d", ",", "-f", "2", "/test.txt"})
		assert.Equal(t, 0, status)
		assert.Equal(t, "b\n2\n", stdout.String())
	})

	t.Run("complement", func(t *testing.T) {
		var stdout bytes.Buffer
		env.Stdout = &stdout
		status := c.Run(context.Background(), env, []string{"-d", ",", "-f", "2", "--complement", "/test.txt"})
		assert.Equal(t, 0, status)
		assert.Equal(t, "a,c\n1,3\n", stdout.String())
	})

	t.Run("output delimiter", func(t *testing.T) {
		var stdout bytes.Buffer
		env.Stdout = &stdout
		status := c.Run(context.Background(), env, []string{"-d", ",", "-f", "1,3", "--output-delimiter", "|", "/test.txt"})
		assert.Equal(t, 0, status)
		assert.Equal(t, "a|c\n1|3\n", stdout.String())
	})
}
