package complete

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestComplete_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		initialSpecs   map[string]*commands.CompSpec
		expectedStatus int
		containsOutput string
		checkEnv       func(t *testing.T, env *commands.Environment)
	}{
		{
			name:           "add simple spec",
			args:           []string{"-f", "-d", "mycmd"},
			expectedStatus: 0,
			checkEnv: func(t *testing.T, env *commands.Environment) {
				spec, ok := env.Completions["mycmd"]
				assert.True(t, ok)
				assert.Equal(t, uint64((1<<3)|(1<<5)), spec.Actions)
			},
		},
		{
			name: "print specs",
			args: []string{"-p"},
			initialSpecs: map[string]*commands.CompSpec{
				"ls": {Actions: 1 << 5},
			},
			expectedStatus: 0,
			containsOutput: "complete -f ls",
		},
		{
			name: "remove spec",
			args: []string{"-r", "ls"},
			initialSpecs: map[string]*commands.CompSpec{
				"ls": {Actions: 1 << 5},
			},
			expectedStatus: 0,
			checkEnv: func(t *testing.T, env *commands.Environment) {
				_, ok := env.Completions["ls"]
				assert.False(t, ok)
			},
		},
		{
			name: "remove all specs",
			args: []string{"-r"},
			initialSpecs: map[string]*commands.CompSpec{
				"ls":   {Actions: 1 << 5},
				"grep": {Actions: 1 << 5},
			},
			expectedStatus: 0,
			checkEnv: func(t *testing.T, env *commands.Environment) {
				assert.Empty(t, env.Completions)
			},
		},
		{
			name:           "complex spec",
			args:           []string{"-W", "word1 word2", "-P", "pre-", "-S", "-suf", "othercmd"},
			expectedStatus: 0,
			checkEnv: func(t *testing.T, env *commands.Environment) {
				spec := env.Completions["othercmd"]
				assert.Equal(t, "word1 word2", spec.WordList)
				assert.Equal(t, "pre-", spec.Prefix)
				assert.Equal(t, "-suf", spec.Suffix)
			},
		},
		{
			name:           "help",
			args:           []string{"--help"},
			expectedStatus: 0,
			containsOutput: "Usage: complete",
		},
		{
			name:           "version",
			args:           []string{"--version"},
			expectedStatus: 0,
			containsOutput: "Version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout:      stdout,
				Stderr:      stderr,
				Completions: make(map[string]*commands.CompSpec),
			}
			for k, v := range tt.initialSpecs {
				env.Completions[k] = v
			}

			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.containsOutput != "" {
				assert.Contains(t, strings.ToLower(stdout.String()), strings.ToLower(tt.containsOutput))
			}
			if tt.checkEnv != nil {
				tt.checkEnv(t, env)
			}
		})
	}
}

func TestComplete_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "complete", c.Name())
}
