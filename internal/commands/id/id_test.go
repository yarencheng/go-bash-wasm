package id

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestId_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		User:   "wasm",
		Uid:    1000,
		Gid:    1000,
		Groups: []int{1000},
	}

	i := New()
	status := i.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "uid=1000(wasm)")
	assert.Contains(t, stdout.String(), "gid=1000(wasm)")
}
