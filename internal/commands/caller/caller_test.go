package caller

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCaller_Run(t *testing.T) {
	env := &commands.Environment{
		Stdout: io.Discard,
		Stderr: io.Discard,
	}
	c := New()
	status := c.Run(context.Background(), env, nil)
	assert.Equal(t, 1, status)
}

func TestCaller_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "caller", c.Name())
}
