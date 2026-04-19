package bind

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestBind_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		contains       string
	}{
		{
			name:           "list functions",
			args:           []string{"-l"},
			expectedStatus: 0,
			contains:       "accept-line",
		},
		{
			name:           "list variables",
			args:           []string{"-v"},
			expectedStatus: 0,
			contains:       "bell-style",
		},
		{
			name:           "print status",
			args:           []string{"-p"},
			expectedStatus: 0,
		},
		{
			name:           "invalid flag",
			args:           []string{"--invalid"},
			expectedStatus: 2,
		},
		{
			name:           "no args",
			args:           []string{},
			expectedStatus: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout: stdout,
				Stderr: stderr,
			}
			b := New()
			status := b.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.contains != "" {
				assert.Contains(t, stdout.String(), tt.contains)
			}
		})
	}
}

func TestBind_Metadata(t *testing.T) {
	b := New()
	assert.Equal(t, "bind", b.Name())
}
