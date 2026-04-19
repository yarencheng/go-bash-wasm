package ptx

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPtx_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		stdin          string
		files          map[string]string
		expectedStatus int
		contains       string
	}{
		{
			name:           "stdin ptx",
			args:           []string{"-"},
			stdin:          "hello world\nfoo bar",
			expectedStatus: 0,
			contains:       "hello world",
		},
		{
			name:           "file ptx",
			args:           []string{"test.txt"},
			files:          map[string]string{"/test.txt": "file content"},
			expectedStatus: 0,
			contains:       "file content",
		},
		{
			name:           "invalid file",
			args:           []string{"notfound.txt"},
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
			fs := afero.NewMemMapFs()
			for path, content := range tt.files {
				_ = afero.WriteFile(fs, path, []byte(content), 0644)
			}

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			stdin := bytes.NewBufferString(tt.stdin)
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: stdout,
				Stderr: stderr,
				Stdin:  io.NopCloser(stdin),
			}
			p := New()
			status := p.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.contains != "" {
				assert.Contains(t, stdout.String(), tt.contains)
			}
		})
	}
}

func TestPtx_Metadata(t *testing.T) {
	p := New()
	assert.Equal(t, "ptx", p.Name())
}
