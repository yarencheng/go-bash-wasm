package boolcmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTrue_Run(t *testing.T) {
	c := NewTrue()
	env := &commands.Environment{}
	
	status := c.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
}

func TestFalse_Run(t *testing.T) {
	c := NewFalse()
	env := &commands.Environment{}
	
	status := c.Run(context.Background(), env, nil)
	assert.Equal(t, 1, status)
}
