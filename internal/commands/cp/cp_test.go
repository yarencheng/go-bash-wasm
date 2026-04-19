package cp

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCp_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		files          map[string]string
		dirs           []string
		expectedStatus int
		checkFiles     map[string]string
		containsOutput string
		containsStderr string
	}{
		{
			name:           "copy file",
			args:           []string{"src.txt", "dest.txt"},
			files:          map[string]string{"/src.txt": "hello"},
			expectedStatus: 0,
			checkFiles:     map[string]string{"/dest.txt": "hello"},
		},
		{
			name:           "copy to directory",
			args:           []string{"src.txt", "dir/"},
			files:          map[string]string{"/src.txt": "hello"},
			dirs:           []string{"/dir"},
			expectedStatus: 0,
			checkFiles:     map[string]string{"/dir/src.txt": "hello"},
		},
		{
			name:           "recursive copy",
			args:           []string{"-r", "srcdir", "destdir"},
			dirs:           []string{"/srcdir", "/srcdir/sub"},
			files:          map[string]string{"/srcdir/a.txt": "a", "/srcdir/sub/b.txt": "b"},
			expectedStatus: 0,
			checkFiles:     map[string]string{"/destdir/a.txt": "a", "/destdir/sub/b.txt": "b"},
		},
		{
			name:           "verbose output",
			args:           []string{"-v", "src.txt", "dest.txt"},
			files:          map[string]string{"/src.txt": "hello"},
			expectedStatus: 0,
			containsOutput: "'/src.txt' -> '/dest.txt'",
		},
		{
			name:           "target directory flag",
			args:           []string{"-t", "dir", "a.txt", "b.txt"},
			dirs:           []string{"/dir"},
			files:          map[string]string{"/a.txt": "a", "/b.txt": "b"},
			expectedStatus: 0,
			checkFiles:     map[string]string{"/dir/a.txt": "a", "/dir/b.txt": "b"},
		},
		{
			name:           "no clobber",
			args:           []string{"-n", "src.txt", "dest.txt"},
			files:          map[string]string{"/src.txt": "new", "/dest.txt": "old"},
			expectedStatus: 0,
			checkFiles:     map[string]string{"/dest.txt": "old"},
		},
		{
			name:           "missing src",
			args:           []string{"missing.txt", "dest.txt"},
			expectedStatus: 1,
			containsStderr: "cannot stat",
		},
		{
			name:           "missing operand",
			args:           []string{"src.txt"},
			files:          map[string]string{"/src.txt": "hello"},
			expectedStatus: 1,
			containsStderr: "missing destination",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			for _, d := range tt.dirs {
				_ = fs.MkdirAll(d, 0755)
			}
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
			}

			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			for path, expectedContent := range tt.checkFiles {
				exists, _ := afero.Exists(fs, path)
				assert.True(t, exists, "File %s should exist", path)
				content, _ := afero.ReadFile(fs, path)
				assert.Equal(t, expectedContent, string(content))
			}

			if tt.containsOutput != "" {
				assert.Contains(t, stdout.String(), tt.containsOutput)
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestCp_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "cp", c.Name())
}
