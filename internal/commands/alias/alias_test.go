package alias

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestAlias_Run(t *testing.T) {
	env := &commands.Environment{
		Aliases: make(map[string]string),
	}

	a := New()
	status := a.Run(context.Background(), env, []string{"ll=ls -l"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "ls -l", env.Aliases["ll"])
}

func TestAlias_Run_FlagP(t *testing.T) {
	out := &strings.Builder{}
	env := &commands.Environment{
		Aliases: map[string]string{"ll": "ls -l"},
		Stdout:  out,
	}

	a := New()
	status := a.Run(context.Background(), env, []string{"-p"})
	assert.Equal(t, 0, status)
	assert.Contains(t, out.String(), "alias ll='ls -l'\n")
}
