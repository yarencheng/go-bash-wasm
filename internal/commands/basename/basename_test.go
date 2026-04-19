package basename

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
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
			args:           []string{"/usr/bin/go"},
			expectedStatus: 0,
			expectedOutput: "go\n",
		},
		{
			name:           "with suffix",
			args:           []string{"/usr/bin/go.sh", ".sh"},
			expectedStatus: 0,
			expectedOutput: "go\n",
		},
		{
			name:           "suffix flag",
			args:           []string{"-s", ".sh", "/usr/bin/go.sh"},
			expectedStatus: 0,
			expectedOutput: "go\n",
		},
		{
			name:           "multiple arguments",
			args:           []string{"-a", "/usr/bin/go", "/usr/bin/python"},
			expectedStatus: 0,
			expectedOutput: "go\npython\n",
		},
		{
			name:           "zero termination",
			args:           []string{"-z", "file"},
			expectedStatus: 0,
			expectedOutput: "file\x00",
		},
		{
			name:           "missing operand",
			args:           []string{},
			expectedStatus: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: stdout,
				Stderr: stderr,
			}

			b := New()
			status := b.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			if tt.expectedOutput != "" {
				assert.Equal(t, tt.expectedOutput, stdout.String())
			}
		})
	}
}

func TestBasename_Metadata(t *testing.T) {
	b := New()
	assert.Equal(t, "basename", b.Name())
}
