package df

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDf_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	d := New()
	status := d.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "/")
}
