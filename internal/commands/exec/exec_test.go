package exec

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/echo"
)

func TestExec_Name(t *testing.T) {
	e := New()
	assert.Equal(t, "exec", e.Name())
}

func TestExec_Run(t *testing.T) {
	registry := commands.New()
	registry.Register(echo.New())

	tests := []struct {
		name       string
		args       []string
		wantStatus int
		wantOut    string
		wantErr    string
		wantExit   bool
	}{
		{
			name:       "exec echo",
			args:       []string{"echo", "hi"},
			wantStatus: 0,
			wantOut:    "hi\n",
			wantExit:   true,
		},
		{
			name:       "command not found",
			args:       []string{"nosuchcmd"},
			wantStatus: 127,
			wantErr:    "command not found",
			wantExit:   true,
		},
		{
			name:       "no arguments",
			args:       []string{},
			wantStatus: 0,
			wantExit:   false,
		},
		{
			name:       "help flag",
			args:       []string{"--help"},
			wantStatus: 0,
			wantOut:    "Usage: exec",
		},
		{
			name:       "version flag",
			args:       []string{"--version"},
			wantStatus: 0,
			wantOut:    "exec",
		},
		{
			name:       "invalid flag",
			args:       []string{"--invalid"},
			wantStatus: 1,
			wantErr:    "unknown flag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout:   out,
				Stderr:   errout,
				Registry: registry,
			}

			e := New()
			status := e.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantOut != "" {
				assert.Contains(t, out.String(), tt.wantOut)
			}
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}
			assert.Equal(t, tt.wantExit, env.ExitRequested)
		})
	}
}
