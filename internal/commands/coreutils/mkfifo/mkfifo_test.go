package mkfifo

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestMkfifo_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		check          func(t *testing.T, fs afero.Fs)
	}{
		{
			name:           "missing operand",
			args:           []string{},
			expectedStatus: 1,
		},
		{
			name: "create fifo",
			args: []string{"myfifo"},
			check: func(t *testing.T, fs afero.Fs) {
				exists, _ := afero.Exists(fs, "/myfifo")
				assert.True(t, exists)
			},
		},
		{
			name: "custom mode",
			args: []string{"-m", "0600", "modedfifo"},
			check: func(t *testing.T, fs afero.Fs) {
				info, err := fs.Stat("/modedfifo")
				assert.NoError(t, err)
				assert.Equal(t, uint32(0600), uint32(info.Mode().Perm()))
			},
		},
		{
			name:           "invalid mode",
			args:           []string{"-m", "999", "failfifo"},
			expectedStatus: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: stdout,
				Stderr: stderr,
				Stdin:  io.NopCloser(bytes.NewReader(nil)),
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

func TestMkfifo_Metadata(t *testing.T) {
	m := New()
	assert.Equal(t, "mkfifo", m.Name())
}
