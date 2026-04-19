package colon

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestColon_Run(t *testing.T) {
	c := New()
	env := &commands.Environment{}

	status := c.Run(context.Background(), env, []string{"arg1", "arg2"})
	assert.Equal(t, 0, status)
}

func TestColon_Name(t *testing.T) {
	c := New()
	assert.Equal(t, ":", c.Name())
}
