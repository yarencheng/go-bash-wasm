package csplit

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"io"
)

func TestCsplit_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		input          string
		expectedStatus int
		checkFiles     map[string]string
		containsOutput string
		containsStderr string
	}{
		{
			name:           "split by line number",
			args:           []string{"-", "3"},
			input:          "line1\nline2\nline3\nline4",
			expectedStatus: 0,
			checkFiles: map[string]string{
				"/xx00": "line1\nline2\n",
				"/xx01": "line3\nline4\n",
			},
		},
		{
			name:           "split by regex",
			args:           []string{"-", "/line3/"},
			input:          "line1\nline2\nline3\nline4",
			expectedStatus: 0,
			checkFiles: map[string]string{
				"/xx00": "line1\nline2\n",
				"/xx01": "line3\nline4\n",
			},
		},
		{
			name:           "custom prefix and digits",
			args:           []string{"-f", "out", "-n", "3", "--", "-", "2"},
			input:          "a\nb\nc",
			expectedStatus: 0,
			checkFiles: map[string]string{
				"/out000": "a\n",
				"/out001": "b\nc\n",
			},
		},
		{
			name:           "quiet mode",
			args:           []string{"-s", "--", "-", "2"},
			input:          "a\nb",
			expectedStatus: 0,
			containsOutput: "", // No sizes printed
		},
		{
			name:           "invalid regex",
			args:           []string{"-", "/[/"},
			input:          "a",
			expectedStatus: 1,
			containsStderr: "invalid pattern",
		},
		{
			name:           "line not found regex",
			args:           []string{"-", "/missing/"},
			input:          "a",
			expectedStatus: 1,
			containsStderr: "line not found",
		},
		{
			name:           "missing operand",
			args:           []string{"-"},
			expectedStatus: 1,
			containsStderr: "missing operand",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			// Create file if it's referenced in args
			for _, arg := range tt.args {
				if arg == "/input.txt" {
					_ = afero.WriteFile(fs, "/input.txt", []byte(tt.input), 0644)
					break
				}
			}

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: stdout,
				Stderr: stderr,
				Stdin:  io.NopCloser(bytes.NewReader([]byte(tt.input))),
			}

			c := New()
			// Make a copy of args to avoid modifying the template
			args := make([]string, len(tt.args))
			copy(args, tt.args)
			status := c.Run(context.Background(), env, args)
			if !assert.Equal(t, tt.expectedStatus, status) {
				t.Logf("STDOUT: %s", stdout.String())
				t.Logf("STDERR: %s", stderr.String())
			}

			for path, expectedContent := range tt.checkFiles {
				content, err := afero.ReadFile(fs, path)
				assert.NoError(t, err)
				assert.Equal(t, expectedContent, string(content))
			}

			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestCsplit_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "csplit", c.Name())
}
