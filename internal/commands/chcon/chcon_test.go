package chcon

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestChcon_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		containsStderr string
	}{
		{
			name:           "basic call",
			args:           []string{"user_t", "file.txt"},
			expectedStatus: 1,
			containsStderr: "SELinux not supported",
		},
		{
			name:           "invalid flag",
			args:           []string{"--invalid"},
			expectedStatus: 1,
			containsStderr: "unknown flag",
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
			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestChcon_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "chcon", c.Name())
}
