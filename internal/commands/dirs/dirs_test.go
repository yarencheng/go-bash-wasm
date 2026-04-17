package dirs

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDirs_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Cwd:      "/home/wasm",
		DirStack: []string{"/etc", "/var"},
		Stdout:   &stdout,
		Stderr:   io.Discard,
	}

	d := New()

	// Test basic listing
	status := d.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/home/wasm /etc /var\n", stdout.String())

	// Test -c (clear)
	stdout.Reset()
	status = d.Run(context.Background(), env, []string{"-c"})
	assert.Equal(t, 0, status)
	assert.Equal(t, 0, len(env.DirStack))

	// Test -p (one per line)
	env.DirStack = []string{"/etc", "/var"}
	stdout.Reset()
	status = d.Run(context.Background(), env, []string{"-p"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "/home/wasm\n/etc\n/var\n", stdout.String())

	// Test -v (verbose)
	stdout.Reset()
	status = d.Run(context.Background(), env, []string{"-v"})
	assert.Equal(t, 0, status)
	assert.Equal(t, " 0  /home/wasm\n 1  /etc\n 2  /var\n", stdout.String())
}
