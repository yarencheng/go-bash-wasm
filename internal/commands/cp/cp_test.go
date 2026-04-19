package cp

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

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
		{
			name:           "help flag",
			args:           []string{"--help"},
			expectedStatus: 0,
			containsOutput: "Usage:",
		},
		{
			name:           "version flag",
			args:           []string{"--version"},
			expectedStatus: 0,
			containsOutput: "cp",
		},
		{
			name:           "archive flag",
			args:           []string{"-a", "srcdir", "destdir"},
			dirs:           []string{"/srcdir"},
			files:          map[string]string{"/srcdir/f": "f"},
			expectedStatus: 0,
			checkFiles:     map[string]string{"/destdir/f": "f"},
		},
		{
			name:           "preserve flag",
			args:           []string{"-p", "src.txt", "dest.txt"},
			files:          map[string]string{"/src.txt": "hello"},
			expectedStatus: 0,
			checkFiles:     map[string]string{"/dest.txt": "hello"},
		},
		{
			name:           "update flag skip",
			args:           []string{"-u", "src.txt", "dest.txt"},
			files:          map[string]string{"/src.txt": "old", "/dest.txt": "new"},
			expectedStatus: 0,
			// Since we can't easily set Chtimes here without manual intervention, 
			// I'll just check it runs. 
			// Use a real test for update in a separate function if needed.
		},
		{
			name:           "omitting directory",
			args:           []string{"srcdir", "destdir"},
			dirs:           []string{"/srcdir"},
			expectedStatus: 1,
			containsStderr: "omitting directory",
		},
		{
			name:           "empty source error",
			args:           []string{},
			expectedStatus: 1,
			containsStderr: "missing file operand",
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

func TestCp_Update(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/newer", []byte("newer"), 0644)
	_ = afero.WriteFile(fs, "/older", []byte("older"), 0644)

	now := time.Now()
	_ = fs.Chtimes("/older", now.Add(-time.Hour), now.Add(-time.Hour))
	_ = fs.Chtimes("/newer", now, now)

	env := &commands.Environment{FS: fs, Cwd: "/", Stdout: io.Discard, Stderr: io.Discard}
	c := New()

	// cp -u older newer -> should not overwrite newer
	status := c.Run(context.Background(), env, []string{"-u", "/older", "/newer"})
	assert.Equal(t, 0, status)
	content, _ := afero.ReadFile(fs, "/newer")
	assert.Equal(t, "newer", string(content))

	// cp -u newer older -> should overwrite older
	status = c.Run(context.Background(), env, []string{"-u", "/newer", "/older"})
	assert.Equal(t, 0, status)
	content, _ = afero.ReadFile(fs, "/older")
	assert.Equal(t, "newer", string(content))
}

func TestCp_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "cp", c.Name())
}
