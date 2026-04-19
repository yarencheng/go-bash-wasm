package continuecmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestContinue_Run(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		expectedStatus   int
		expectedContinue int
		containsStderr   string
	}{
		{
			name:             "default continue",
			args:             []string{},
			expectedStatus:   0,
			expectedContinue: 1,
		},
		{
			name:             "continue 2",
			args:             []string{"2"},
			expectedStatus:   0,
			expectedContinue: 2,
		},
		{
			name:             "invalid numeric",
			args:             []string{"abc"},
			expectedStatus:   1,
			containsStderr:   "numeric argument required",
		},
		{
			name:             "invalid count",
			args:             []string{"0"},
			expectedStatus:   1,
			containsStderr:   "must be greater than or equal to 1",
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
			if tt.expectedStatus == 0 {
				assert.Equal(t, tt.expectedContinue, env.ContinueRequested)
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestContinue_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "continue", c.Name())
}
