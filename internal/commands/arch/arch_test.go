package arch

import (
	"bytes"
	"context"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestArch_Run(t *testing.T) {
	stdout := &bytes.Buffer{}
	env := &commands.Environment{
		Stdout: stdout,
	}
	a := New()
	status := a.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, runtime.GOARCH+"\n", stdout.String())
}

func TestArch_Metadata(t *testing.T) {
	a := New()
	assert.Equal(t, "arch", a.Name())
}
