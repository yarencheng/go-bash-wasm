package pwd

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPwd_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Cwd:    "/home/user",
	}
	p := New()
	status := p.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "/home/user\n", stdout.String())
}
