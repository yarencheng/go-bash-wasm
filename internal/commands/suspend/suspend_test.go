package suspend

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestSuspend_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
	}{
		{
			name:           "basic suspend",
			args:           []string{},
			expectedStatus: 1,
		},
		{
			name:           "suspend forced",
			args:           []string{"-f"},
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
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				Stderr: stderr,
			}

			s := New()
			status := s.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
		})
	}
}

func TestSuspend_Metadata(t *testing.T) {
	s := New()
	assert.Equal(t, "suspend", s.Name())
}
