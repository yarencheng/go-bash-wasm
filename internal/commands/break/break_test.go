package breakcmd

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestBreak_Run(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedBreak int
		expectedCode  int
	}{
		{
			name:          "basic break",
			args:          []string{},
			expectedBreak: 1,
			expectedCode:  0,
		},
		{
			name:          "break 2",
			args:          []string{"2"},
			expectedBreak: 2,
			expectedCode:  0,
		},
		{
			name:          "break invalid",
			args:          []string{"invalid"},
			expectedBreak: 0,
			expectedCode:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := &commands.Environment{
				Stdout: io.Discard,
				Stderr: io.Discard,
			}
			b := New()
			status := b.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedCode, status)
			assert.Equal(t, tt.expectedBreak, env.BreakRequested)
		})
	}
}
