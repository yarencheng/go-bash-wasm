package breakcmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestBreak_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		expectedBreak  int
		containsStderr string
	}{
		{
			name:           "default break",
			args:           []string{},
			expectedStatus: 0,
			expectedBreak:  1,
		},
		{
			name:           "break 2",
			args:           []string{"2"},
			expectedStatus: 0,
			expectedBreak:  2,
		},
		{
			name:           "invalid numeric",
			args:           []string{"abc"},
			expectedStatus: 1,
			containsStderr: "numeric argument required",
		},
		{
			name:           "invalid count",
			args:           []string{"0"},
			expectedStatus: 1,
			containsStderr: "must be greater than or equal to 1",
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
			if tt.expectedStatus == 0 {
				assert.Equal(t, tt.expectedBreak, env.BreakRequested)
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestBreak_Metadata(t *testing.T) {
	b := New()
	assert.Equal(t, "break", b.Name())
}
