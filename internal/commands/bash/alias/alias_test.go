package alias

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestAlias_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		initialAliases map[string]string
		expectedStatus int
		containsOutput string
		containsStderr string
		checkEnv       func(t *testing.T, env *commands.Environment)
	}{
		{
			name: "define alias",
			args: []string{"ll=ls -l"},
			checkEnv: func(t *testing.T, env *commands.Environment) {
				assert.Equal(t, "ls -l", env.Aliases["ll"])
			},
			expectedStatus: 0,
		},
		{
			name: "print specific alias",
			args: []string{"g"},
			initialAliases: map[string]string{
				"g": "grep",
			},
			expectedStatus: 0,
			containsOutput: "alias g='grep'",
		},
		{
			name: "alias not found",
			args: []string{"missing"},
			expectedStatus: 1,
			containsStderr: "not found",
		},
		{
			name: "list all aliases",
			args: []string{"-p"},
			initialAliases: map[string]string{
				"a": "echo a",
				"b": "echo b",
			},
			expectedStatus: 0,
			containsOutput: "alias a='echo a'\nalias b='echo b'",
		},
		{
			name: "invalid flag",
			args: []string{"--invalid"},
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
				Aliases: make(map[string]string),
			}
			for k, v := range tt.initialAliases {
				env.Aliases[k] = v
			}

			a := New()
			status := a.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.containsOutput != "" {
				assert.Contains(t, stdout.String(), tt.containsOutput)
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
			if tt.checkEnv != nil {
				tt.checkEnv(t, env)
			}
		})
	}
}

func TestAlias_Metadata(t *testing.T) {
	a := New()
	assert.Equal(t, "alias", a.Name())
}
