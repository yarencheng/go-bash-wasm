package pushd

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPushd_Basic(t *testing.T) {
	fs := afero.NewMemMapFs()
	fs.MkdirAll("/dir1", 0755)
	
	env := &commands.Environment{
		FS:       fs,
		Stdout:   io.Discard,
		Cwd:      "/",
		DirStack: []string{},
	}

	p := New()
	status := p.Run(context.Background(), env, []string{"/dir1"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/dir1", env.Cwd)
	assert.Equal(t, []string{"/"}, env.DirStack)
}

func TestPushd_NoChdir(t *testing.T) {
	fs := afero.NewMemMapFs()
	fs.MkdirAll("/dir1", 0755)

	env := &commands.Environment{
		FS:       fs,
		Stdout:   io.Discard,
		Cwd:      "/",
		DirStack: []string{},
	}

	p := New()
	// -n should add /dir1 to stack but NOT change CWD
	status := p.Run(context.Background(), env, []string{"-n", "/dir1"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/", env.Cwd)
	// Bash pushd -n /dir1 adds /dir1 to the SECOND position? 
	// Actually, Bash pushd -n dir adds dir to the stack but keeps CWD.
	// The stack becomes: CWD dir old_stack...
	assert.Contains(t, env.DirStack, "/dir1")
}
