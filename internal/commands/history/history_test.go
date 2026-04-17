package history

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestHistory_Basic(t *testing.T) {
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdout: out,
		History: []string{
			"ls",
			"pwd",
		},
	}

	h := New()
	status := h.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Contains(t, out.String(), "1  ls")
	assert.Contains(t, out.String(), "2  pwd")
}

func TestHistory_Clear(t *testing.T) {
	env := &commands.Environment{
		History: []string{"ls"},
	}

	h := New()
	status := h.Run(context.Background(), env, []string{"-c"})
	assert.Equal(t, 0, status)
	assert.Empty(t, env.History)
}
