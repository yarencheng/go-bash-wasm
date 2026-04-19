package basename

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestBasename_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		expectedOutput string
	}{
		{
			name:           "basic",
			args:           []string{"/usr/bin/test"},
			expectedStatus: 0,
			expectedOutput: "test\n",
		},
		{
			name:           "with suffix",
			args:           []string{"/usr/bin/test.txt", ".txt"},
			expectedStatus: 0,
			expectedOutput: "test\n",
		},
		{
			name:           "with suffix flag",
			args:           []string{"-s", ".txt", "/usr/bin/test.txt"},
			expectedStatus: 0,
			expectedOutput: "test\n",
		},
		{
			name:           "multiple",
			args:           []string{"-a", "/usr/bin/test1", "/usr/bin/test2"},
			expectedStatus: 0,
			expectedOutput: "test1\ntest2\n",
		},
		{
			name:           "zero",
			args:           []string{"-z", "test"},
			expectedStatus: 0,
			expectedOutput: "test\x00",
		},
		{
			name:           "missing operand",
			args:           []string{},
			expectedStatus: 1,
		},
		{
			name:           "invalid flag",
			args:           []string{"--invalid"},
			expectedStatus: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout: stdout,
				Stderr: stderr,
			}
			b := New()
			status := b.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.expectedStatus == 0 {
				assert.Equal(t, tt.expectedOutput, stdout.String())
			}
		})
	}
}

func TestBasename_Metadata(t *testing.T) {
	b := New()
	assert.Equal(t, "basename", b.Name())
}
