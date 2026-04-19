package tty

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTTY_Basic(t *testing.T) {
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdout: out,
	}

	tr := New()
	status := tr.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Contains(t, out.String(), "/dev/tty")
}
