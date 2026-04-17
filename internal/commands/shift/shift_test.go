package shift

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestShift_Run(t *testing.T) {
	env := &commands.Environment{
		PositionalArgs: []string{"a", "b", "c"},
		Stderr: io.Discard,
	}

	s := New()

	// shift 1
	status := s.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, []string{"b", "c"}, env.PositionalArgs)

	// shift 2
	status = s.Run(context.Background(), env, []string{"2"})
	assert.Equal(t, 0, status)
	assert.Equal(t, []string{}, env.PositionalArgs)

	// shift out of range
	env.PositionalArgs = []string{"x"}
	status = s.Run(context.Background(), env, []string{"2"})
	assert.Equal(t, 1, status)
	assert.Equal(t, []string{"x"}, env.PositionalArgs)
}
