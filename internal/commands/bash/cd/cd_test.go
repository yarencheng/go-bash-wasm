package cd

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCd_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		initialCwd     string
		initialEnv     map[string]string
		dirs           []string
		expectedStatus int
		expectedCwd    string
		containsOutput string
		containsStderr string
	}{
		{
			name:           "change directory",
			args:           []string{"/home/user"},
			initialCwd:     "/",
			dirs:           []string{"/home/user"},
			expectedStatus: 0,
			expectedCwd:    "/home/user",
		},
		{
			name:           "cd to HOME",
			args:           []string{},
			initialCwd:     "/tmp",
			initialEnv:     map[string]string{"HOME": "/home/user"},
			dirs:           []string{"/home/user"},
			expectedStatus: 0,
			expectedCwd:    "/home/user",
		},
		{
			name:           "cd to OLDPWD",
			args:           []string{"-"},
			initialCwd:     "/home/user",
			initialEnv:     map[string]string{"OLDPWD": "/etc"},
			dirs:           []string{"/etc"},
			expectedStatus: 0,
			expectedCwd:    "/etc",
			containsOutput: "/etc",
		},
		{
			name:           "cd to parent",
			args:           []string{".."},
			initialCwd:     "/home/user",
			dirs:           []string{"/home"},
			expectedStatus: 0,
			expectedCwd:    "/home",
		},
		{
			name:           "relative cd",
			args:           []string{"bin"},
			initialCwd:     "/usr",
			dirs:           []string{"/usr/bin"},
			expectedStatus: 0,
			expectedCwd:    "/usr/bin",
		},
		{
			name:           "cd path",
			args:           []string{"sub"},
			initialCwd:     "/home",
			initialEnv:     map[string]string{"CDPATH": "/other"},
			dirs:           []string{"/other/sub"},
			expectedStatus: 0,
			expectedCwd:    "/other/sub",
			containsOutput: "/other/sub",
		},
		{
			name:           "not a directory",
			args:           []string{"/test.txt"},
			initialCwd:     "/",
			expectedStatus: 1,
			containsStderr: "Not a directory",
		},
		{
			name:           "missing target",
			args:           []string{"/missing"},
			initialCwd:     "/",
			expectedStatus: 1,
			containsStderr: "No such file or directory",
		},
		{
			name:           "HOME not set",
			args:           []string{},
			initialCwd:     "/",
			expectedStatus: 1,
			containsStderr: "HOME not set",
		},
		{
			name:           "invalid flag",
			args:           []string{"--invalid"},
			expectedStatus: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			for _, d := range tt.dirs {
				_ = fs.MkdirAll(d, 0755)
			}
			// For "not a directory" test
			if tt.name == "not a directory" {
				_ = afero.WriteFile(fs, "/test.txt", []byte(""), 0644)
			}

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				FS:      fs,
				Cwd:     tt.initialCwd,
				Stdout:  stdout,
				Stderr:  stderr,
				EnvVars: make(map[string]string),
			}
			for k, v := range tt.initialEnv {
				env.EnvVars[k] = v
			}

			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.expectedStatus == 0 {
				assert.Equal(t, tt.expectedCwd, env.Cwd)
				assert.Equal(t, tt.expectedCwd, env.EnvVars["PWD"])
				if tt.initialCwd != "" {
					assert.Equal(t, tt.initialCwd, env.EnvVars["OLDPWD"])
				}
			}
			if tt.containsOutput != "" {
				assert.Contains(t, stdout.String(), tt.containsOutput)
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestCd_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "cd", c.Name())
}
