package continuecmd

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestContinue_Run(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		expectedContinue int
		expectedCode     int
	}{
		{
			name:             "basic continue",
			args:             []string{},
			expectedContinue: 1,
			expectedCode:     0,
		},
		{
			name:             "continue 2",
			args:             []string{"2"},
			expectedContinue: 2,
			expectedCode:     0,
		},
		{
			name:             "continue invalid",
			args:             []string{"invalid"},
			expectedContinue: 0,
			expectedCode:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := &commands.Environment{
				Stdout: io.Discard,
				Stderr: io.Discard,
			}
			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedCode, status)
			assert.Equal(t, tt.expectedContinue, env.ContinueRequested)
		})
	}
}
