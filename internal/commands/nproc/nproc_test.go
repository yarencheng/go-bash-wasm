package nproccmd

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestNproc_Basic(t *testing.T) {
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdout: out,
	}

	n := New()
	status := n.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	
	val := strings.TrimSpace(out.String())
	assert.NotEmpty(t, val)
}
