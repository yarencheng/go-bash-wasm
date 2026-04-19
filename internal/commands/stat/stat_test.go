package stat

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestStat_Name(t *testing.T) {
	s := New()
	assert.Equal(t, "stat", s.Name())
}

func TestStat_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	mtime := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello"), 0644))
	require.NoError(t, fs.Chtimes("/test.txt", mtime, mtime))

	tests := []struct {
		name       string
		args       []string
		wantStatus int
		wantOut    string
		wantErr    string
	}{
		{
			name:       "basic stat",
			args:       []string{"/test.txt"},
			wantStatus: 0,
			wantOut:    "File: /test.txt\n  Size: 5",
		},
		{
			name:       "terse mode",
			args:       []string{"-t", "/test.txt"},
			wantStatus: 0,
			wantOut:    "/test.txt 5 -rw-r--r-- 2026-01-01 00:00:00",
		},
		{
			name:       "format mode",
			args:       []string{"-c", "%n %s %A %y", "/test.txt"},
			wantStatus: 0,
			wantOut:    "/test.txt 5 -rw-r--r-- 2026-01-01 00:00:00.000000000 +0000",
		},
		{
			name:       "dereference (basic)",
			args:       []string{"-L", "/test.txt"},
			wantStatus: 0,
			wantOut:    "File: /test.txt",
		},
		{
			name:       "help flag",
			args:       []string{"--help"},
			wantStatus: 0,
			wantOut:    "Usage: stat",
		},
		{
			name:       "version flag",
			args:       []string{"--version"},
			wantStatus: 0,
			wantOut:    "stat", // ShowVersion output usually contains the name
		},
		{
			name:       "missing operand",
			args:       []string{},
			wantStatus: 1,
			wantErr:    "missing operand",
		},
		{
			name:       "invalid flag",
			args:       []string{"--invalid"},
			wantStatus: 1,
			wantErr:    "unknown flag",
		},
		{
			name:       "file not found",
			args:       []string{"/nonexistent"},
			wantStatus: 1,
			wantErr:    "cannot stat '/nonexistent'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: out,
				Stderr: errout,
			}

			s := New()
			status := s.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantOut != "" {
				assert.Contains(t, out.String(), tt.wantOut)
			}
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}
		})
	}
}
