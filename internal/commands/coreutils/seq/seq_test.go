package seq

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestSeq_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	s := New()
	status := s.Run(context.Background(), env, []string{"3"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "1\n2\n3\n", stdout.String())

	stdout.Reset()
	status = s.Run(context.Background(), env, []string{"5", "8"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "5\n6\n7\n8\n", stdout.String())
}
