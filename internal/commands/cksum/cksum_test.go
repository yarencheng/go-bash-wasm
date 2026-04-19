package cksum

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

func TestCksum_Run(t *testing.T) {
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
			name:           "default crc",
			args:           []string{"test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "4222801193 5 test.txt",
		},
		{
			name:           "md5 hash",
			args:           []string{"-a", "md5", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "5d41402abc4b2a76b9719d911017c592  test.txt",
		},
		{
			name:           "sha256 hash base64",
			args:           []string{"-a", "sha256", "--base64", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=",
		},
		{
			name:           "check mode",
			args:           []string{"-c", "sums.txt"},
			files:          map[string]string{
				"/test.txt": "hello",
				"/sums.txt": "5d41402abc4b2a76b9719d911017c592  test.txt",
			},
			expectedStatus: 0,
			containsOutput: "test.txt: OK",
		},
		{
			name:           "check mode fail",
			args:           []string{"-c", "sums.txt"},
			files:          map[string]string{
				"/test.txt": "wrong",
				"/sums.txt": "5d41402abc4b2a76b9719d911017c592  test.txt",
			},
			expectedStatus: 1,
			containsOutput: "test.txt: FAILED",
		},
		{
			name:           "bsd style tag",
			args:           []string{"-a", "md5", "--tag", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "MD5 (test.txt) = 5d41402abc4b2a76b9719d911017c592",
		},
		{
			name:           "zero terminator",
			args:           []string{"-z", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "4222801193 5 test.txt\x00",
		},
		{
			name:           "unknown algorithm",
			args:           []string{"-a", "unknown", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 1,
			containsStderr: "unknown algorithm",
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
				assert.Contains(t, stdout.String(), tt.containsOutput)
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestCksum_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "cksum", c.Name())
}
