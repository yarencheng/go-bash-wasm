package history

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestHistory_Name(t *testing.T) {
	h := New()
	assert.Equal(t, "history", h.Name())
}

func TestHistory_Run(t *testing.T) {
	envVars := map[string]string{"HISTFILE": "/.bash_history"}

	tests := []struct {
		name          string
		args          []string
		initialHist   []string
		wantStatus    int
		wantHist      []string
		wantOut       string
		wantErr       string
		preWrite      string
		noHistFileVar bool
	}{
		{
			name:        "display history",
			args:        []string{},
			initialHist: []string{"ls", "pwd"},
			wantStatus:  0,
			wantOut:     "1  ls\n    2  pwd\n",
		},
		{
			name:        "display limited history",
			args:        []string{"1"},
			initialHist: []string{"ls", "pwd"},
			wantStatus:  0,
			wantOut:     "2  pwd\n",
		},
		{
			name:        "clear history",
			args:        []string{"-c"},
			initialHist: []string{"ls", "pwd"},
			wantStatus:  0,
			wantHist:    nil,
		},
		{
			name:        "delete entry",
			args:        []string{"-d", "1"},
			initialHist: []string{"ls", "pwd"},
			wantStatus:  0,
			wantHist:    []string{"pwd"},
		},
		{
			name:        "delete out of range",
			args:        []string{"-d", "99"},
			initialHist: []string{"ls"},
			wantStatus:  1,
			wantErr:     "position out of range",
		},
		{
			name:        "write history",
			args:        []string{"-w"},
			initialHist: []string{"ls", "pwd"},
			wantStatus:  0,
		},
		{
			name:        "append history",
			args:        []string{"-a"},
			initialHist: []string{"echo hi"},
			wantStatus:  0,
		},
		{
			name:        "read history",
			args:        []string{"-r"},
			preWrite:    "ls\npwd\n",
			wantStatus:  0,
			wantHist:    []string{"ls", "pwd"},
		},
		{
			name:        "read history fail",
			args:        []string{"-r"},
			wantStatus:  1,
			wantErr:     "file does not exist",
		},
		{
			name:        "print expand",
			args:        []string{"-p", "!!", "echo hi"},
			wantStatus:  0,
			wantOut:     "!!\necho hi\n",
		},
		{
			name:        "store entry",
			args:        []string{"-s", "echo", "hello"},
			wantStatus:  0,
			wantHist:    []string{"echo hello"},
		},
		{
			name:        "invalid flag",
			args:        []string{"-X"},
			wantStatus:  1,
			wantErr:     "unknown shorthand flag",
		},
		{
			name:        "default histfile",
			args:        []string{"-w"},
			initialHist: []string{"ls"},
			noHistFileVar: true,
			wantStatus:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			actualEnvVars := make(map[string]string)
			if !tt.noHistFileVar {
				for k, v := range envVars {
					actualEnvVars[k] = v
				}
			}
			env := &commands.Environment{
				FS:      afero.NewMemMapFs(), // New FS for each test
				History: tt.initialHist,
				EnvVars: actualEnvVars,
				Stdout:  out,
				Stderr:  errout,
			}

			if tt.preWrite != "" {
				histFile := "/.bash_history"
				_ = afero.WriteFile(env.FS, histFile, []byte(tt.preWrite), 0644)
			}

			h := New()
			status := h.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantHist != nil {
				assert.Equal(t, tt.wantHist, env.History)
			}
			if tt.wantOut != "" {
				assert.Contains(t, out.String(), tt.wantOut)
			}
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}

			if tt.name == "write history" {
				data, _ := afero.ReadFile(env.FS, "/.bash_history")
				assert.Contains(t, string(data), "ls\npwd\n")
			}
			if tt.name == "default histfile" {
				data, _ := afero.ReadFile(env.FS, "/home/wasm/.bash_history")
				assert.Contains(t, string(data), "ls\n")
			}
		})
	}
}
