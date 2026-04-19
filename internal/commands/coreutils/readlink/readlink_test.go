package readlink

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockSymlinkFs struct {
	afero.Fs
	readlink func(name string) (string, error)
}

func (m *mockSymlinkFs) ReadlinkIfPossible(name string) (string, error) {
	return m.readlink(name)
}

func (m *mockSymlinkFs) SymlinkIfPossible(oldname, newname string) error {
	return nil
}

func (m *mockSymlinkFs) LstatIfPossible(name string) (os.FileInfo, bool, error) {
	if l, ok := m.Fs.(afero.Lstater); ok {
		return l.LstatIfPossible(name)
	}
	return nil, false, nil
}

func TestReadlink_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		setup          func(fs afero.Fs) afero.Fs
		expectedStatus int
		expectedOutput string
		containsStderr string
	}{
		{
			name: "read valid symlink",
			args: []string{"link"},
			setup: func(fs afero.Fs) afero.Fs {
				return &mockSymlinkFs{
					Fs: fs,
					readlink: func(name string) (string, error) {
						if name == "/link" {
							return "/target", nil
						}
						return name, nil
					},
				}
			},
			expectedStatus: 0,
			expectedOutput: "/target\n",
		},
		{
			name: "read not a symlink",
			args: []string{"file"},
			setup: func(fs afero.Fs) afero.Fs {
				_ = afero.WriteFile(fs, "/file", []byte("content"), 0644)
				return &mockSymlinkFs{
					Fs: fs,
					readlink: func(name string) (string, error) {
						return name, nil
					},
				}
			},
			expectedStatus: 1,
		},
		{
			name: "no newline",
			args: []string{"-n", "link"},
			setup: func(fs afero.Fs) afero.Fs {
				return &mockSymlinkFs{
					Fs: fs,
					readlink: func(name string) (string, error) {
						if name == "/link" {
							return "/target", nil
						}
						return name, nil
					},
				}
			},
			expectedStatus: 0,
			expectedOutput: "/target",
		},
		{
			name: "canonicalize",
			args: []string{"-f", "/path/../target"},
			setup: func(fs afero.Fs) afero.Fs {
				_ = afero.WriteFile(fs, "/target", []byte("content"), 0644)
				return fs
			},
			expectedStatus: 0,
			expectedOutput: "/target\n",
		},
		{
			name:           "missing target",
			args:           []string{},
			expectedStatus: 1,
		},
		{
			name: "verbose error",
			args: []string{"-v", "nonexistent"},
			setup: func(fs afero.Fs) afero.Fs {
				return fs // No symlink support
			},
			expectedStatus: 1,
			containsStderr: "readlink",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fs afero.Fs = afero.NewMemMapFs()
			if tt.setup != nil {
				fs = tt.setup(fs)
			}

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: stdout,
				Stderr: stderr,
			}

			r := New()
			status := r.Run(context.Background(), env, tt.args)
			if !assert.Equal(t, tt.expectedStatus, status) {
				t.Logf("STDOUT: %s", stdout.String())
				t.Logf("STDERR: %s", stderr.String())
			}

			if tt.expectedOutput != "" {
				assert.Equal(t, tt.expectedOutput, stdout.String())
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestReadlink_Metadata(t *testing.T) {
	r := New()
	assert.Equal(t, "readlink", r.Name())
}
