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
			name:           "sha1 hash",
			args:           []string{"-a", "sha1", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d  test.txt",
		},
		{
			name:           "sha224 hash",
			args:           []string{"-a", "sha224", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "ea09ae9cc6768c50fcee903ed054556e5bfc8347907f12598aa24193  test.txt",
		},
		{
			name:           "sha384 hash",
			args:           []string{"-a", "sha384", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "59e1748777448c69de6b800d7a33bbfb9ff1b463e44354c3553bcdb9c666fa90125a3c79f90397bdf5f6a13de828684f  test.txt",
		},
		{
			name:           "sha512 hash",
			args:           []string{"-a", "sha512", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043  test.txt",
		},
		{
			name:           "status flag",
			args:           []string{"--status", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			// should have no output
		},
		{
			name:           "length flag",
			args:           []string{"-a", "sha256", "--length", "128", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "2cf24dba5fb0a30e26e83b2ac5b9e29e  test.txt",
		},
		{
			name:           "zero flag multiple files",
			args:           []string{"-z", "f1.txt", "f2.txt"},
			files:          map[string]string{
				"/f1.txt": "a",
				"/f2.txt": "b",
			},
			expectedStatus: 0,
			containsOutput: "4240663063 1 f1.txt\x004173770459 1 f2.txt\x00",
		},
		{
			name:           "stdin input crc",
			args:           []string{},
			stdin:          "hello",
			expectedStatus: 0,
			containsOutput: "4222801193 5",
		},
		{
			name:           "stdin input md5",
			args:           []string{"-a", "md5"},
			stdin:          "hello",
			expectedStatus: 0,
			containsOutput: "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name:           "raw output",
			args:           []string{"-a", "md5", "--raw", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			// raw bytes of md5 "hello"
			containsOutput: "\x5d\x41\x40\x2a\xbc\x4b\x2a\x76\xb9\x71\x9d\x91\x10\x17\xc5\x92",
		},
		{
			name:           "untagged output",
			args:           []string{"-a", "md5", "--untagged", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name:           "check mode quiet",
			args:           []string{"-c", "--quiet", "sums.txt"},
			files:          map[string]string{
				"/test.txt": "hello",
				"/sums.txt": "5d41402abc4b2a76b9719d911017c592  test.txt",
			},
			expectedStatus: 0,
			// should not contain "test.txt: OK"
		},
		{
			name:           "check mode status",
			args:           []string{"-c", "--status", "sums.txt"},
			files:          map[string]string{
				"/test.txt": "hello",
				"/sums.txt": "5d41402abc4b2a76b9719d911017c592  test.txt",
			},
			expectedStatus: 0,
			// should not contain any output
		},
		{
			name:           "check mode skip missing",
			args:           []string{"-c", "--ignore-missing", "sums.txt"},
			files:          map[string]string{
				"/sums.txt": "5d41402abc4b2a76b9719d911017c592  missing.txt",
			},
			expectedStatus: 0,
		},
		{
			name:           "check mode warn improper",
			args:           []string{"-c", "--warn", "sums.txt"},
			files:          map[string]string{
				"/sums.txt": "invalid_line",
			},
			expectedStatus: 0,
			containsStderr: "improperly formatted line",
		},
		{
			name:           "check mode strict improper",
			args:           []string{"-c", "--strict", "sums.txt"},
			files:          map[string]string{
				"/sums.txt": "invalid_line",
			},
			expectedStatus: 1,
		},
		{
			name:           "file not found",
			args:           []string{"missing.txt"},
			expectedStatus: 1,
			containsStderr: "file does not exist",
		},

		{
			name:           "bsd style tag stdin",
			args:           []string{"-a", "md5", "--tag"},
			stdin:          "hello",
			expectedStatus: 0,
			containsOutput: "MD5 = 5d41402abc4b2a76b9719d911017c592",
		},
		{
			name:           "unknown algorithm",
			args:           []string{"-a", "unknown", "test.txt"},
			files:          map[string]string{"/test.txt": "hello"},
			expectedStatus: 1,
			containsStderr: "unknown algorithm",
		},
		{
			name:           "bsd style check",
			args:           []string{"-c", "sums.txt"},
			files:          map[string]string{
				"/test.txt": "hello",
				"/sums.txt": "MD5 (test.txt) = 5d41402abc4b2a76b9719d911017c592",
			},
			expectedStatus: 0,
			containsOutput: "test.txt: OK",
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

			if tt.name == "check mode quiet" {
				assert.NotContains(t, stdout.String(), "test.txt: OK")
			}
			if tt.name == "check mode status" {
				assert.Empty(t, stdout.String())
				assert.Empty(t, stderr.String())
			}
		})
	}
}

func TestCksum_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "cksum", c.Name())
}
