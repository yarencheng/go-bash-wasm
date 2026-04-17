package uname

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUname_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
	}

	u := New()
	status := u.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "BashWasm")
}
