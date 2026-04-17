package touch

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTouch_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	tr := New()

	// Test basic touch (create file)
	status := tr.Run(context.Background(), env, []string{"newfile.txt"})
	assert.Equal(t, 0, status)
	_, err := fs.Stat("/newfile.txt")
	assert.NoError(t, err)

	// Test touch existing file
	status = tr.Run(context.Background(), env, []string{"newfile.txt"})
	assert.Equal(t, 0, status)
}
