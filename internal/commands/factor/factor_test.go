package factor

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestFactor_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	f := New()

	tests := []struct {
		arg      string
		expected string
	}{
		{"12", "12: 2 2 3\n"},
		{"13", "13: 13\n"},
		{"1", "1:\n"},
	}

	for _, tt := range tests {
		stdout.Reset()
		status := f.Run(context.Background(), env, []string{tt.arg})
		assert.Equal(t, 0, status)
		assert.Equal(t, tt.expected, stdout.String())
	}
}
