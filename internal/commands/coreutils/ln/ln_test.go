package ln

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestLn_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	lnCmd := New()

	// Test symbolic link
	status := lnCmd.Run(context.Background(), env, []string{"-s", "/test.txt", "/link.txt"})

	// MemMapFs might not support symlinks
	if _, ok := fs.(interface {
		SymlinkIfPossible(oldname, newname string) error
	}); ok {
		assert.Equal(t, 0, status)
		_, err := fs.Stat("/link.txt")
		assert.NoError(t, err)
	} else {
		assert.Equal(t, 1, status)
	}
}
