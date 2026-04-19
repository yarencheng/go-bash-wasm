package printenv

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPrintenv_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		envVars        map[string]string
		expectedStatus int
		containsOutput []string
	}{
		{
			name: "print all",
			args: []string{},
			envVars: map[string]string{
				"VAR1": "val1",
				"VAR2": "val2",
			},
			expectedStatus: 0,
			containsOutput: []string{"VAR1=val1", "VAR2=val2"},
		},
		{
			name: "print specific",
			args: []string{"VAR1"},
			envVars: map[string]string{
				"VAR1": "val1",
				"VAR2": "val2",
			},
			expectedStatus: 0,
			containsOutput: []string{"val1"},
		},
		{
			name:           "variable not found",
			args:           []string{"MISSING"},
			envVars:        map[string]string{"VAR1": "val1"},
			expectedStatus: 1,
			containsOutput: []string{},
		},
		{
			name: "multiple specific",
			args: []string{"VAR1", "VAR2"},
			envVars: map[string]string{
				"VAR1": "val1",
				"VAR2": "val2",
			},
			expectedStatus: 0,
			containsOutput: []string{"val1", "val2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout:  stdout,
				EnvVars: tt.envVars,
			}
			p := New()
			status := p.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			for _, expected := range tt.containsOutput {
				assert.Contains(t, stdout.String(), expected)
			}
		})
	}
}

func TestPrintenv_Metadata(t *testing.T) {
	p := New()
	assert.Equal(t, "printenv", p.Name())
}
