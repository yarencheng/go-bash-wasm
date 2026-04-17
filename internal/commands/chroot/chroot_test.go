package chroot

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pwd"
)

func TestChroot_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = fs.MkdirAll("/newroot/etc", 0755)
	_ = afero.WriteFile(fs, "/newroot/etc/hostname", []byte("chrooted-host"), 0644)
	
	var stdout bytes.Buffer
	registry := commands.New()
	registry.Register(pwd.New())
	
	env := &commands.Environment{
		FS:       fs,
		Stdout:   &stdout,
		Registry: registry,
		Cwd:      "/",
	}

	c := New()
	// chroot /newroot pwd
	status := c.Run(context.Background(), env, []string{"/newroot", "pwd"})
	assert.Equal(t, 0, status)
	// After chroot, pwd should report /
	assert.Equal(t, "/\n", stdout.String())
}
