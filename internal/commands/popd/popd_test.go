package popd

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPopd_Basic(t *testing.T) {
	env := &commands.Environment{
		Stdout:   io.Discard,
		Cwd:      "/dir1",
		DirStack: []string{"/"},
	}

	p := New()
	status := p.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "/", env.Cwd)
	assert.Equal(t, []string{}, env.DirStack)
}

func TestPopd_NoChdir(t *testing.T) {
	env := &commands.Environment{
		Stdout:   io.Discard,
		Cwd:      "/dir1",
		DirStack: []string{"/"},
	}

	p := New()
	// -n should remove / from stack but NOT change CWD
	status := p.Run(context.Background(), env, []string{"-n"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/dir1", env.Cwd)
	assert.Equal(t, []string{}, env.DirStack)
}
