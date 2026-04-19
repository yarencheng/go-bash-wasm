package chmod

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestChmod_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		files          map[string]os.FileMode
		expectedStatus int
		checkModes     map[string]os.FileMode
		containsStderr string
	}{
		{
			name:           "numeric mode",
			args:           []string{"0755", "test.txt"},
			files:          map[string]os.FileMode{"/test.txt": 0644},
			expectedStatus: 0,
			checkModes:     map[string]os.FileMode{"/test.txt": 0755},
		},
		{
			name:           "symbolic add",
			args:           []string{"u+x", "test.txt"},
			files:          map[string]os.FileMode{"/test.txt": 0644},
			expectedStatus: 0,
			checkModes:     map[string]os.FileMode{"/test.txt": 0744},
		},
		{
			name:           "symbolic subtract",
			args:           []string{"g-w", "test.txt"},
			files:          map[string]os.FileMode{"/test.txt": 0664},
			expectedStatus: 0,
			checkModes:     map[string]os.FileMode{"/test.txt": 0644},
		},
		{
			name:           "symbolic set",
			args:           []string{"a=r", "test.txt"},
			files:          map[string]os.FileMode{"/test.txt": 0777},
			expectedStatus: 0,
			checkModes:     map[string]os.FileMode{"/test.txt": 0444},
		},
		{
			name:           "recursive chmod",
			args:           []string{"-R", "0777", "dir"},
			files:          map[string]os.FileMode{"/dir": 0755, "/dir/file": 0644},
			expectedStatus: 0,
			checkModes:     map[string]os.FileMode{"/dir": 0777, "/dir/file": 0777},
		},
		{
			name:           "reference file",
			args:           []string{"--reference=ref.txt", "target.txt"},
			files:          map[string]os.FileMode{"/ref.txt": 0700, "/target.txt": 0644},
			expectedStatus: 0,
			checkModes:     map[string]os.FileMode{"/target.txt": 0700},
		},
		{
			name:           "missing file",
			args:           []string{"0755", "missing.txt"},
			expectedStatus: 1,
			containsStderr: "cannot access",
		},
		{
			name:           "invalid mode",
			args:           []string{"invalid", "test.txt"},
			files:          map[string]os.FileMode{"/test.txt": 0644},
			expectedStatus: 1,
			containsStderr: "invalid mode",
		},
		{
			name:           "missing operand",
			args:           []string{"0755"},
			expectedStatus: 1,
			containsStderr: "missing operand",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			for path, mode := range tt.files {
				if mode.IsDir() {
					_ = fs.MkdirAll(path, mode.Perm())
				} else {
					_ = afero.WriteFile(fs, path, []byte(""), mode.Perm())
					// Some filesystems might ignore Chmod on create, so enforce it
					_ = fs.Chmod(path, mode.Perm())
				}
			}

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: stdout,
				Stderr: stderr,
			}

			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			for path, expectedMode := range tt.checkModes {
				info, err := fs.Stat(path)
				assert.NoError(t, err)
				assert.Equal(t, expectedMode.Perm(), info.Mode().Perm(), "Mode mismatch for %s", path)
			}

			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestChmod_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "chmod", c.Name())
}
