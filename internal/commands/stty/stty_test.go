package stty

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestStty_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		contains       string
	}{
		{
			name:           "basic stty",
			args:           []string{},
			expectedStatus: 0,
			contains:       "speed",
		},
		{
			name:           "stty -a",
			args:           []string{"-a"},
			expectedStatus: 0,
			contains:       "baud",
		},
		{
			name:           "stty with file",
			args:           []string{"-F", "/dev/tty"},
			expectedStatus: 0,
		},
		{
			name:           "invalid flag",
			args:           []string{"--invalid"},
			expectedStatus: 1,
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

			s := New()
			status := s.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.contains != "" {
				assert.Contains(t, stdout.String(), tt.contains)
			}
		})
	}
}

func TestStty_Metadata(t *testing.T) {
	s := New()
	assert.Equal(t, "stty", s.Name())
}
