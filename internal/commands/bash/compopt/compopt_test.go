package compopt

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCompopt_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
	}{
		{
			name:           "basic call",
			args:           []string{"-o", "nospace"},
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
			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
		})
	}
}

func TestCompopt_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "compopt", c.Name())
}
