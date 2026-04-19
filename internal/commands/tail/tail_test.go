package tail

import (
	"bytes"
	"context"
	"io"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTail_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		input          string
		files          map[string]string
		expectedStatus int
		containsOutput []string
		contextTimeout bool
	}{
		{
			name:           "stdin last 2 lines",
			args:           []string{"-n", "2"},
			input:          "l1\nl2\nl3\nl4\n",
			expectedStatus: 0,
			containsOutput: []string{"l3\nl4\n"},
		},
		{
			name: "file last 5 bytes",
			args: []string{"-c", "5", "file1"},
			files: map[string]string{
				"file1": "hello world\n",
			},
			expectedStatus: 0,
			containsOutput: []string{"orld\n"},
		},
		{
			name: "multiple files with headers",
			args: []string{"-n", "1", "f1", "f2"},
			files: map[string]string{
				"f1": "line1\n",
				"f2": "line2\n",
			},
			expectedStatus: 0,
			containsOutput: []string{"==> f1 <==", "line1\n", "==> f2 <==", "line2\n"},
		},
		{
			name:           "missing file",
			args:           []string{"nonexistent"},
			expectedStatus: 1,
		},
		{
			name:           "zero terminated stdin",
			args:           []string{"-z", "-n", "2"},
			input:          "l1\x00l2\x00l3\x00l4\x00",
			expectedStatus: 0,
			containsOutput: []string{"l3\x00l4\x00"},
		},
		{
			name: "multiple files with headers and quiet",
			args: []string{"-n", "1", "-q", "f1", "f2"},
			files: map[string]string{
				"f1": "line1\n",
				"f2": "line2\n",
			},
			expectedStatus: 0,
			containsOutput: []string{"line1\n", "line2\n"},
		},
		{
			name: "verbose header single file",
			args: []string{"-n", "1", "-v", "f1"},
			files: map[string]string{
				"f1": "line1\n",
			},
			expectedStatus: 0,
			containsOutput: []string{"==> f1 <==", "line1\n"},
		},
		{
			name:           "help flag",
			args:           []string{"--help"},
			expectedStatus: 0,
			containsOutput: []string{"Usage: tail"},
		},
		{
			name:           "version flag",
			args:           []string{"--version"},
			expectedStatus: 0,
			containsOutput: []string{"tail (go-bash-wasm)"},
		},
		{
			name: "multiple files last 5 bytes",
			args: []string{"-c", "5", "f1", "f2"},
			files: map[string]string{
				"f1": "hello world\n",
				"f2": "goodbye world\n",
			},
			expectedStatus: 0,
			containsOutput: []string{"orld\n", "orld\n", "==> f1 <==", "==> f2 <=="},
		},
		{
			name: "follow flag",
			args: []string{"-f", "f1"},
			files: map[string]string{
				"f1": "line1\n",
			},
			expectedStatus: 0,
			contextTimeout: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			for name, content := range tt.files {
				path := name
				if !filepath.IsAbs(path) {
					path = "/" + path
				}
				_ = afero.WriteFile(fs, path, []byte(content), 0644)
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

			tail := New()
			ctx := context.Background()
			if tt.contextTimeout {
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				cancel() // Cancel immediately
			}
			status := tail.Run(ctx, env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			for _, expected := range tt.containsOutput {
				assert.Contains(t, stdout.String(), expected)
			}
		})
	}
}

func TestTail_Metadata(t *testing.T) {
	tail := New()
	assert.Equal(t, "tail", tail.Name())
}
