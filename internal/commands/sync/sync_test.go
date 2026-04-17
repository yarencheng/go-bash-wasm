package sync

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestSync_Run(t *testing.T) {
	env := &commands.Environment{
		Stdout: io.Discard,
		Stderr: io.Discard,
	}

	s := New()

	status := s.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
}
