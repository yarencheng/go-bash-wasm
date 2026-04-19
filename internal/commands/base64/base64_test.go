package base64cmd

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

func TestBase64_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		stdin          string
		files          map[string]string
		expectedStatus int
		expectedOutput string
	}{
		{
			name:           "encode stdin",
			args:           []string{},
			stdin:          "hello",
			expectedStatus: 0,
			expectedOutput: "aGVsbG8=\n",
		},
		{
			name:           "decode stdin",
			args:           []string{"-d"},
			stdin:          "aGVsbG8=",
			expectedStatus: 0,
			expectedOutput: "hello",
		},
		{
			name:           "encode file",
			args:           []string{"test.txt"},
			files:          map[string]string{"/test.txt": "world"},
			expectedStatus: 0,
			expectedOutput: "d29ybGQ=\n",
		},
		{
			name:           "encode with wrap",
			args:           []string{"-w", "2"},
			stdin:          "hello",
			expectedStatus: 0,
			expectedOutput: "aG\nVs\nbG\n8=\n",
		},
		{
			name:           "decode error",
			args:           []string{"-d"},
			stdin:          "invalid!!!",
			expectedStatus: 1,
		},
		{
			name:           "file not found",
			args:           []string{"missing.txt"},
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
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: stdout,
				Stderr: stderr,
				Stdin:  io.NopCloser(strings.NewReader(tt.stdin)),
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

func TestBase64_Metadata(t *testing.T) {
	b := New()
	assert.Equal(t, "base64", b.Name())
}
