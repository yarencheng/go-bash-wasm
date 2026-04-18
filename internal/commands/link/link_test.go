package link

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestLink_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/src.txt", []byte("hello"), 0644)
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	ln := New()
	status := ln.Run(context.Background(), env, []string{"/src.txt", "/dst.txt"})

	// MemMapFs doesn't support Link, so we expect status 1 (error)
	// If it ever supports it, this test will fail and we can update it.
	assert.Equal(t, 1, status)
}
