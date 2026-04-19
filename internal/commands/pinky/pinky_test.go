package pinky

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPinky_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		contains       string
	}{
		{
			name:           "basic pinky",
			args:           []string{},
			expectedStatus: 0,
			contains:       "Login",
		},
		{
			name:           "pinky with flags",
			args:           []string{"-l"},
			expectedStatus: 0,
			contains:       "Where",
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
			p := New()
			status := p.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.contains != "" {
				assert.Contains(t, stdout.String(), tt.contains)
			}
		})
	}
}

func TestPinky_Metadata(t *testing.T) {
	p := New()
	assert.Equal(t, "pinky", p.Name())
}
