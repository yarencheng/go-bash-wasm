package comm

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

func TestComm_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/file1.txt", []byte("a\nb\nc\n"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/file2.txt", []byte("b\nc\nd\n"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stdout: &bytes.Buffer{},
		Stderr: io.Discard,
	}

	c := New()
	status := c.Run(context.Background(), env, []string{"/file1.txt", "/file2.txt"})
	assert.Equal(t, 0, status)
	// Expected:
	// a
	// 	b
	// 	c
	// 		d (wait, 'd' is only in file 2)
	// Output should have 3 columns: 1: only in f1, 2: only in f2, 3: in both
	// a (col 1)
	//     b (col 3)
	//     c (col 3)
	//   d (col 2)
	assert.Contains(t, env.Stdout.(*bytes.Buffer).String(), "a\n")
}
