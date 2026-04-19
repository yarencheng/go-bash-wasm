package touch

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTouch_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		setup          func(fs afero.Fs)
		expectedStatus int
		check          func(t *testing.T, fs afero.Fs)
	}{
		{
			name: "create new file",
			args: []string{"newfile"},
			check: func(t *testing.T, fs afero.Fs) {
				exists, _ := afero.Exists(fs, "/newfile")
				assert.True(t, exists)
			},
			expectedStatus: 0,
		},
		{
			name: "update existing file time",
			args: []string{"oldfile"},
			setup: func(fs afero.Fs) {
				_ = afero.WriteFile(fs, "/oldfile", []byte("data"), 0644)
				oldTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
				_ = fs.Chtimes("/oldfile", oldTime, oldTime)
			},
			check: func(t *testing.T, fs afero.Fs) {
				info, _ := fs.Stat("/oldfile")
				assert.True(t, info.ModTime().After(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)))
			},
			expectedStatus: 0,
		},
		{
			name: "no create flag",
			args: []string{"-c", "nonexistent"},
			check: func(t *testing.T, fs afero.Fs) {
				exists, _ := afero.Exists(fs, "/nonexistent")
				assert.False(t, exists)
			},
			expectedStatus: 0,
		},
		{
			name: "reference file",
			args: []string{"-r", "ref", "target"},
			setup: func(fs afero.Fs) {
				refTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
				_ = afero.WriteFile(fs, "/ref", []byte("ref"), 0644)
				_ = fs.Chtimes("/ref", refTime, refTime)
				_ = afero.WriteFile(fs, "/target", []byte("target"), 0644)
			},
			check: func(t *testing.T, fs afero.Fs) {
				info, _ := fs.Stat("/target")
				assert.True(t, info.ModTime().Equal(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)))
			},
			expectedStatus: 0,
		},
		{
			name: "date string",
			args: []string{"-d", "2022-02-02", "datefile"},
			check: func(t *testing.T, fs afero.Fs) {
				info, _ := fs.Stat("/datefile")
				assert.Equal(t, 2022, info.ModTime().Year())
				assert.Equal(t, time.February, info.ModTime().Month())
				assert.Equal(t, 2, info.ModTime().Day())
			},
			expectedStatus: 0,
		},
		{
			name: "time string",
			args: []string{"-t", "202303030303", "timefile"},
			check: func(t *testing.T, fs afero.Fs) {
				info, _ := fs.Stat("/timefile")
				assert.Equal(t, 2023, info.ModTime().Year())
				assert.Equal(t, 3, int(info.ModTime().Month()))
				assert.Equal(t, 3, info.ModTime().Day())
				assert.Equal(t, 3, info.ModTime().Hour())
				assert.Equal(t, 3, info.ModTime().Minute())
			},
			expectedStatus: 0,
		},
		{
			name:           "missing reference",
			args:           []string{"-r", "missing", "target"},
			expectedStatus: 1,
		},
		{
			name:           "invalid date",
			args:           []string{"-d", "invalid", "file"},
			expectedStatus: 1,
		},
		{
			name:           "invalid time",
			args:           []string{"-t", "invalid", "file"},
			expectedStatus: 1,
		},
		{
			name: "time with seconds",
			args: []string{"-t", "202303030303.59", "secfile"},
			check: func(t *testing.T, fs afero.Fs) {
				info, _ := fs.Stat("/secfile")
				assert.Equal(t, 59, info.ModTime().Second())
			},
			expectedStatus: 0,
		},
		{
			name: "short year format",
			args: []string{"-t", "9903030303", "shortyrfile"},
			check: func(t *testing.T, fs afero.Fs) {
				info, _ := fs.Stat("/shortyrfile")
				assert.Equal(t, 1999, info.ModTime().Year())
			},
			expectedStatus: 0,
		},
		{
			name: "future short year format",
			args: []string{"-t", "0103030303", "futureyrfile"},
			check: func(t *testing.T, fs afero.Fs) {
				info, _ := fs.Stat("/futureyrfile")
				assert.Equal(t, 2001, info.ModTime().Year())
			},
			expectedStatus: 0,
		},
		{
			name: "no year format",
			args: []string{"-t", "03030303", "noyrfile"},
			check: func(t *testing.T, fs afero.Fs) {
				info, _ := fs.Stat("/noyrfile")
				assert.Equal(t, time.Now().Year(), info.ModTime().Year())
			},
			expectedStatus: 0,
		},
		{
			name: "modification time only",
			args: []string{"-m", "mfile"},
			setup: func(fs afero.Fs) {
				_ = afero.WriteFile(fs, "/mfile", []byte("data"), 0644)
			},
			expectedStatus: 0,
		},
		{
			name: "access time only",
			args: []string{"-a", "afile"},
			setup: func(fs afero.Fs) {
				_ = afero.WriteFile(fs, "/afile", []byte("data"), 0644)
			},
			expectedStatus: 0,
		},
		{
			name:           "invalid pflag",
			args:           []string{"--invalid"},
			expectedStatus: 1,
		},
		{
			name:           "missing file",
			args:           []string{},
			expectedStatus: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			if tt.setup != nil {
				tt.setup(fs)
			}

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdout: stdout,
				Stderr: stderr,
			}

			touch := New()
			status := touch.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			if tt.check != nil {
				tt.check(t, fs)
			}
		})
	}
}

func TestTouch_Metadata(t *testing.T) {
	touch := New()
	assert.Equal(t, "touch", touch.Name())
}
