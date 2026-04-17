package printf

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPrintf_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	p := New()
	status := p.Run(context.Background(), env, []string{"hello %s", "world"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello world", stdout.String())
}
