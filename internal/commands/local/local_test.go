package local

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestLocal_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		checkEnv       func(t *testing.T, env *commands.Environment)
	}{
		{
			name:           "declare and assign",
			args:           []string{"foo=bar"},
			expectedStatus: 0,
			checkEnv: func(t *testing.T, env *commands.Environment) {
				assert.Equal(t, "bar", env.EnvVars["foo"])
			},
		},
		{
			name:           "declare multiple",
			args:           []string{"x=1", "y=2"},
			expectedStatus: 0,
			checkEnv: func(t *testing.T, env *commands.Environment) {
				assert.Equal(t, "1", env.EnvVars["x"])
				assert.Equal(t, "2", env.EnvVars["y"])
			},
		},
		{
			name:           "declare without value",
			args:           []string{"z"},
			expectedStatus: 0,
			checkEnv: func(t *testing.T, env *commands.Environment) {
				assert.Equal(t, "", env.EnvVars["z"])
			},
		},
		{
			name:           "no args",
			args:           []string{},
			expectedStatus: 0,
		},
		{
			name:           "invalid flag",
			args:           []string{"--invalid"},
			expectedStatus: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout:  stdout,
				Stderr:  stderr,
				EnvVars: make(map[string]string),
			}
			l := New()
			status := l.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.checkEnv != nil {
				tt.checkEnv(t, env)
			}
		})
	}
}

func TestLocal_Metadata(t *testing.T) {
	l := New()
	assert.Equal(t, "local", l.Name())
}
