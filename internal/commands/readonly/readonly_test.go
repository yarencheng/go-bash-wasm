package readonly

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestReadonly_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		envVars        map[string]string
		expectedStatus int
		contains       string
		expectedEnv    map[string]string
	}{
		{
			name:           "list readonly",
			args:           []string{},
			envVars:        map[string]string{"foo": "bar"},
			expectedStatus: 0,
			contains:       "readonly foo=\"bar\"",
		},
		{
			name:           "set and list",
			args:           []string{"-p", "x=y"},
			envVars:        map[string]string{"foo": "bar"},
			expectedStatus: 0,
			contains:       "readonly x=\"y\"",
			expectedEnv:    map[string]string{"foo": "bar", "x": "y"},
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
			for k, v := range tt.envVars {
				env.EnvVars[k] = v
			}

			r := New()
			status := r.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.contains != "" {
				assert.Contains(t, stdout.String(), tt.contains)
			}
			if tt.expectedEnv != nil {
				for k, v := range tt.expectedEnv {
					assert.Equal(t, v, env.EnvVars[k])
				}
			}
		})
	}
}

func TestReadonly_Metadata(t *testing.T) {
	r := New()
	assert.Equal(t, "readonly", r.Name())
}
