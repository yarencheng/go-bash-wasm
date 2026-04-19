package chroot

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockCmd struct {
	name    string
	retCode int
	runArgs []string
	checkFS func(fs afero.Fs) bool
}

func (m *mockCmd) Name() string { return m.name }
func (m *mockCmd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	m.runArgs = args
	if m.checkFS != nil && !m.checkFS(env.FS) {
		return 1
	}
	return m.retCode
}

func TestChroot_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		dirs           []string
		files          map[string]string
		expectedStatus int
		containsStderr string
		checkFS        func(fs afero.Fs) bool
	}{
		{
			name: "basic chroot execution",
			args: []string{"/jail", "hello", "world"},
			dirs: []string{"/jail"},
			checkFS: func(fs afero.Fs) bool {
				// Injail, / should be mapped to original /jail
				_ = afero.WriteFile(fs, "/in_jail.txt", []byte("locked"), 0644)
				return true
			},
			expectedStatus: 0,
		},
		{
			name:           "missing root directory",
			args:           []string{"/missing", "ls"},
			expectedStatus: 1,
			containsStderr: "cannot change root directory",
		},
		{
			name:           "not a directory",
			args:           []string{"/file.txt", "ls"},
			files:          map[string]string{"/file.txt": "content"},
			expectedStatus: 1,
			containsStderr: "Not a directory",
		},
		{
			name:           "command not found",
			args:           []string{"/jail", "unknown_cmd"},
			dirs:           []string{"/jail"},
			expectedStatus: 127,
			containsStderr: "command not found",
		},
		{
			name:           "missing operand",
			args:           []string{},
			expectedStatus: 1,
			containsStderr: "missing operand",
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

			registry := commands.New()
			helloMock := &mockCmd{name: "hello", retCode: 0, checkFS: tt.checkFS}
			_ = registry.Register(helloMock)

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				FS:       fs,
				Cwd:      "/",
				Stdout:   stdout,
				Stderr:   stderr,
				Registry: registry,
			}

			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			if tt.name == "basic chroot execution" {
				// Verify file was written to original /jail/in_jail.txt
				exists, _ := afero.Exists(fs, "/jail/in_jail.txt")
				assert.True(t, exists)
			}

			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestChroot_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "chroot", c.Name())
}
