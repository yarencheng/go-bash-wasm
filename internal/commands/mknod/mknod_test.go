package mknod

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestMknod_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		check          func(t *testing.T, fs afero.Fs)
	}{
		{
			name:           "missing operand",
			args:           []string{"pipe"},
			expectedStatus: 1,
		},
		{
			name:           "invalid device type",
			args:           []string{"node", "x"},
			expectedStatus: 1,
		},
		{
			name:           "create named pipe",
			args:           []string{"mypipe", "p"},
			expectedStatus: 0,
			check: func(t *testing.T, fs afero.Fs) {
				_, err := fs.Stat("/mypipe")
				assert.NoError(t, err)
			},
		},
		{
			name:           "create character device",
			args:           []string{"char-dev", "c", "1", "3"},
			expectedStatus: 0,
			check: func(t *testing.T, fs afero.Fs) {
				_, err := fs.Stat("/char-dev")
				assert.NoError(t, err)
			},
		},
		{
			name:           "create block device",
			args:           []string{"block-dev", "b", "1", "3"},
			expectedStatus: 0,
			check: func(t *testing.T, fs afero.Fs) {
				_, err := fs.Stat("/block-dev")
				assert.NoError(t, err)
			},
		},
		{
			name:           "missing major/minor",
			args:           []string{"node", "c"},
			expectedStatus: 1,
		},
		{
			name:           "custom mode",
			args:           []string{"-m", "0600", "modepipe", "p"},
			expectedStatus: 0,
			check: func(t *testing.T, fs afero.Fs) {
				info, err := fs.Stat("/modepipe")
				require.NoError(t, err)
				assert.Equal(t, os.FileMode(0600), info.Mode().Perm())
			},
		},
		{
			name:           "invalid mode",
			args:           []string{"-m", "999", "fail", "p"},
			expectedStatus: 1,
		},
		{
			name:           "already exists",
			args:           []string{"exists", "p"},
			expectedStatus: 0,
			check: func(t *testing.T, fs afero.Fs) {
				m := New()
				env := &commands.Environment{
					FS:     fs,
					Cwd:    "/",
					Stderr: io.Discard,
				}
				status := m.Run(context.Background(), env, []string{"exists", "p"})
				assert.Equal(t, 1, status) // Error because os.O_EXCL
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stderr: io.Discard,
			}
			m := New()
			status := m.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.check != nil {
				tt.check(t, fs)
			}
		})
	}
}

func TestMknod_Metadata(t *testing.T) {
	m := New()
	assert.Equal(t, "mknod", m.Name())
}
