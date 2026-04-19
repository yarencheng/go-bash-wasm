package cat

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCat_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		files          map[string]string
		stdin          string
		expectedStatus int
		containsOutput string
		containsStderr string
	}{
		{
			name:           "basic cat",
			args:           []string{"test.txt"},
			files:          map[string]string{"/test.txt": "hello world"},
			expectedStatus: 0,
			containsOutput: "hello world",
		},
		{
			name:           "cat stdin",
			args:           []string{"-"},
			stdin:          "from stdin",
			expectedStatus: 0,
			containsOutput: "from stdin",
		},
		{
			name:           "cat missing file",
			args:           []string{"missing.txt"},
			expectedStatus: 1,
			containsStderr: "file does not exist",
		},
		{
			name:           "number lines",
			args:           []string{"-n", "test.txt"},
			files:          map[string]string{"/test.txt": "line1\nline2"},
			expectedStatus: 0,
			containsOutput: "     1\tline1\n     2\tline2",
		},
		{
			name:           "number nonblank",
			args:           []string{"-b", "test.txt"},
			files:          map[string]string{"/test.txt": "line1\n\nline2"},
			expectedStatus: 0,
			containsOutput: "     1\tline1\n\n     2\tline2",
		},
		{
			name:           "squeeze blank",
			args:           []string{"-s", "test.txt"},
			files:          map[string]string{"/test.txt": "line1\n\n\nline2"},
			expectedStatus: 0,
			containsOutput: "line1\n\nline2",
		},
		{
			name:           "show ends",
			args:           []string{"-E", "test.txt"},
			files:          map[string]string{"/test.txt": "line1\nline2"},
			expectedStatus: 0,
			containsOutput: "line1$\nline2",
		},
		{
			name:           "show tabs",
			args:           []string{"-T", "test.txt"},
			files:          map[string]string{"/test.txt": "a\tb"},
			expectedStatus: 0,
			containsOutput: "a^Ib",
		},
		{
			name:           "show all",
			args:           []string{"-A", "test.txt"},
			files:          map[string]string{"/test.txt": "a\t\x01\n"},
			expectedStatus: 0,
			containsOutput: "a^I^A$\n",
		},
		{
			name:           "version",
			args:           []string{"--version"},
			expectedStatus: 0,
			containsOutput: "Version",
		},
		{
			name:           "help",
			args:           []string{"--help"},
			expectedStatus: 0,
			containsOutput: "Usage: cat",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			for path, content := range tt.files {
				_ = afero.WriteFile(fs, path, []byte(content), 0644)
			}

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: stdout,
				Stderr: stderr,
				Stdin:  io.NopCloser(strings.NewReader(tt.stdin)),
			}

			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.containsOutput != "" {
				assert.Contains(t, strings.ToLower(stdout.String()), strings.ToLower(tt.containsOutput))
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestCat_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "cat", c.Name())
}
