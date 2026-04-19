package dirname

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDirname_Run(t *testing.T) {
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
			expectedOutput: "/usr/bin\n",
		},
		{
			name:           "multiple arguments",
			args:           []string{"/usr/bin/go", "/home/user/test.txt"},
			expectedStatus: 0,
			expectedOutput: "/usr/bin\n/home/user\n",
		},
		{
			name:           "zero termination",
			args:           []string{"-z", "/usr/bin/go"},
			expectedStatus: 0,
			expectedOutput: "/usr/bin\x00",
		},
		{
			name:           "no arguments",
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

			d := New()
			status := d.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			if tt.expectedOutput != "" {
				assert.Equal(t, tt.expectedOutput, stdout.String())
			}
		})
	}
}

func TestDirname_Metadata(t *testing.T) {
	d := New()
	assert.Equal(t, "dirname", d.Name())
}
