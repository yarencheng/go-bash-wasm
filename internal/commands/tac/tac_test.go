package tac

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTac_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		input          string
		files          map[string]string
		expectedStatus int
		expectedOutput string
	}{
		{
			name:           "stdin lines",
			args:           []string{},
			input:          "line1\nline2\nline3\n",
			expectedStatus: 0,
			expectedOutput: "line3\nline2\nline1\n",
		},
		{
			name: "single file",
			args: []string{"file1"},
			files: map[string]string{
				"file1": "a\nb\nc\n",
			},
			expectedStatus: 0,
			expectedOutput: "c\nb\na\n",
		},
		{
			name: "custom separator",
			args: []string{"-s", ":", "file1"},
			files: map[string]string{
				"file1": "a:b:c:",
			},
			expectedStatus: 0,
			expectedOutput: "c:b:a:",
		},
		{
			name: "separator before",
			args: []string{"-b", "-s", ":", "file1"},
			files: map[string]string{
				"file1": "a:b:c",
			},
			expectedStatus: 0,
			expectedOutput: ":c:b:a",
		},
		{
			name:           "missing file",
			args:           []string{"nonexistent"},
			expectedStatus: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			for name, content := range tt.files {
				_ = afero.WriteFile(fs, name, []byte(content), 0644)
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

			tac := New()
			status := tac.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			if tt.expectedOutput != "" {
				assert.Equal(t, tt.expectedOutput, stdout.String())
			}
		})
	}
}

func TestTac_Metadata(t *testing.T) {
	tac := New()
	assert.Equal(t, "tac", tac.Name())
}
