package shred

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestShred_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		files          map[string]string
		expectedStatus int
		check          func(t *testing.T, fs afero.Fs)
	}{
		{
			name:           "missing file",
			args:           []string{},
			expectedStatus: 1,
		},
		{
			name:           "shred non-existent file",
			args:           []string{"notfound"},
			expectedStatus: 1,
		},
		{
			name:  "basic shred",
			args:  []string{"test.txt"},
			files: map[string]string{"/test.txt": "some content"},
			check: func(t *testing.T, fs afero.Fs) {
				data, err := afero.ReadFile(fs, "/test.txt")
				require.NoError(t, err)
				assert.NotEqual(t, "some content", string(data))
				assert.Equal(t, len("some content"), len(data))
			},
		},
		{
			name:  "shred and remove",
			args:  []string{"-u", "remove.txt"},
			files: map[string]string{"/remove.txt": "delete me"},
			check: func(t *testing.T, fs afero.Fs) {
				exists, _ := afero.Exists(fs, "/remove.txt")
				assert.False(t, exists)
			},
		},
		{
			name:  "shred with zero and verbose",
			args:  []string{"-z", "-v", "zero.txt"},
			files: map[string]string{"/zero.txt": "zero me"},
			check: func(t *testing.T, fs afero.Fs) {
				data, err := afero.ReadFile(fs, "/zero.txt")
				require.NoError(t, err)
				for _, b := range data {
					assert.Equal(t, byte(0), b)
				}
			},
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
				Stdin:  io.NopCloser(bytes.NewReader(nil)),
			}

			s := New()
			status := s.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.check != nil {
				tt.check(t, fs)
			}
		})
	}
}

func TestShred_Metadata(t *testing.T) {
	s := New()
	assert.Equal(t, "shred", s.Name())
}
