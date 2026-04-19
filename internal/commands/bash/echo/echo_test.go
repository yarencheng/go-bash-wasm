package echo

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestEcho_Run(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "basic echo",
			args:     []string{"hello", "world"},
			expected: "hello world\n",
		},
		{
			name:     "echo -n",
			args:     []string{"-n", "hello", "world"},
			expected: "hello world",
		},
		{
			name:     "echo -e",
			args:     []string{"-e", "hello\\nworld"},
			expected: "hello\nworld\n",
		},
		{
			name:     "echo -E",
			args:     []string{"-E", "hello\\nworld"},
			expected: "hello\\nworld\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			env := &commands.Environment{
				Stdout: &stdout,
			}
			e := New()
			status := e.Run(context.Background(), env, tt.args)
			assert.Equal(t, 0, status)
			assert.Equal(t, tt.expected, stdout.String())
		})
	}
}
