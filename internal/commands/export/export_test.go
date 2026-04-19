package export

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestExport_Name(t *testing.T) {
	e := New()
	assert.Equal(t, "export", e.Name())
}

func TestExport_Run(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		envVars    map[string]string
		wantStatus int
		wantOut    string
		wantErr    string
		checkVars  map[string]string
	}{
		{
			name:       "set variable",
			args:       []string{"FOO=BAR"},
			envVars:    make(map[string]string),
			wantStatus: 0,
			checkVars:  map[string]string{"FOO": "BAR"},
		},
		{
			name:       "set multiple variables",
			args:       []string{"A=1", "B=2"},
			envVars:    make(map[string]string),
			wantStatus: 0,
			checkVars:  map[string]string{"A": "1", "B": "2"},
		},
		{
			name:       "list variables",
			args:       []string{},
			envVars:    map[string]string{"Z": "last", "A": "first"},
			wantStatus: 0,
			wantOut:    "export A=\"first\"\nexport Z=\"last\"\n",
		},
		{
			name:       "help flag",
			args:       []string{"--help"},
			wantStatus: 0,
			wantOut:    "Usage: export",
		},
		{
			name:       "version flag",
			args:       []string{"--version"},
			wantStatus: 0,
			wantOut:    "export",
		},
		{
			name:       "invalid flag",
			args:       []string{"--invalid"},
			wantStatus: 1,
			wantErr:    "unknown flag",
		},
		{
			name:       "ignore no equal",
			args:       []string{"NOEQUAL"},
			envVars:    make(map[string]string),
			wantStatus: 0,
			checkVars:  map[string]string{"NOEQUAL": ""}, // Should NOT be in checkVars if we follow current code
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout:  out,
				Stderr:  errout,
				EnvVars: tt.envVars,
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
			for k, v := range tt.checkVars {
				assert.Equal(t, v, env.EnvVars[k])
			}
			if tt.name == "ignore no equal" {
				_, ok := env.EnvVars["NOEQUAL"]
				assert.False(t, ok)
			}
		})
	}
}
