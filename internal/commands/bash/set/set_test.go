package set

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestSet_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		envVars        map[string]string
		expectedStatus int
		contains       string
		checkEnv       func(t *testing.T, env *commands.Environment)
	}{
		{
			name:           "list variables",
			args:           []string{},
			envVars:        map[string]string{"foo": "bar"},
			expectedStatus: 0,
			contains:       "foo=bar",
		},
		{
			name:           "set positional parameters",
			args:           []string{"--", "a", "b", "c"},
			expectedStatus: 0,
			checkEnv: func(t *testing.T, env *commands.Environment) {
				assert.Equal(t, []string{"a", "b", "c"}, env.PositionalArgs)
			},
		},
		{
			name:           "invalid flag",
			args:           []string{"--invalid"},
			expectedStatus: 2,
		},
		{
			name:           "help",
			args:           []string{"--help"},
			expectedStatus: 0,
			contains:       "Usage: set",
		},
		{
			name:           "version",
			args:           []string{"--version"},
			expectedStatus: 0,
			contains:       "version",
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
			for k, v := range tt.envVars {
				env.EnvVars[k] = v
			}

			s := New()
			status := s.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.contains != "" {
				assert.Contains(t, strings.ToLower(stdout.String()), strings.ToLower(tt.contains))
			}
			if tt.checkEnv != nil {
				tt.checkEnv(t, env)
			}
		})
	}
}

func TestSet_Metadata(t *testing.T) {
	s := New()
	assert.Equal(t, "set", s.Name())
}
