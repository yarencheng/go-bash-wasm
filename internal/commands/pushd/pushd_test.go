package pushd

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPushd_Name(t *testing.T) {
	p := New()
	assert.Equal(t, "pushd", p.Name())
}

func TestPushd_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = fs.MkdirAll("/dir1", 0755)
	_ = fs.MkdirAll("/dir2", 0755)

	tests := []struct {
		name          string
		args          []string
		initialCwd    string
		initialStack  []string
		wantStatus    int
		wantCwd       string
		wantStack     []string
		wantOut       string
		wantErr       string
	}{
		{
			name:         "push specific dir",
			args:         []string{"/dir1"},
			initialCwd:   "/",
			initialStack: []string{},
			wantStatus:   0,
			wantCwd:      "/dir1",
			wantStack:    []string{"/"},
			wantOut:      "/dir1 /",
		},
		{
			name:         "push no-chdir",
			args:         []string{"-n", "/dir1"},
			initialCwd:   "/",
			initialStack: []string{},
			wantStatus:   0,
			wantCwd:      "/",
			wantStack:    []string{"/dir1"},
			wantOut:      "/ /dir1",
		},
		{
			name:         "swap with no args",
			args:         []string{},
			initialCwd:   "/",
			initialStack: []string{"/dir1"},
			wantStatus:   0,
			wantCwd:      "/dir1",
			wantStack:    []string{"/"},
			wantOut:      "/dir1 /",
		},
		{
			name:         "swap with no args multiple",
			args:         []string{},
			initialCwd:   "/dir1",
			initialStack: []string{"/dir2", "/"},
			wantStatus:   0,
			wantCwd:      "/",
			wantStack:    []string{"/dir2", "/dir1"},
			wantOut:      "/ /dir2 /dir1",
		},
		{
			name:         "no other dir error",
			args:         []string{},
			initialCwd:   "/",
			initialStack: []string{},
			wantStatus:   1,
			wantErr:      "no other directory",
		},
		{
			name:         "directory not found",
			args:         []string{"/nonexistent"},
			initialCwd:   "/",
			wantStatus:   1,
			wantErr:      "No such directory",
		},
		{
			name:         "help flag",
			args:         []string{"--help"},
			wantStatus:   0,
			wantOut:      "Usage: pushd",
		},
		{
			name:         "version flag",
			args:         []string{"--version"},
			wantStatus:   0,
			wantOut:      "pushd",
		},
		{
			name:         "invalid flag",
			args:         []string{"--invalid"},
			wantStatus:   1,
			wantErr:      "unknown flag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				FS:       fs,
				Cwd:      tt.initialCwd,
				DirStack: tt.initialStack,
				Stdout:   out,
				Stderr:   errout,
			}

			p := New()
			status := p.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantStatus == 0 && tt.name != "help flag" && tt.name != "version flag" {
				assert.Equal(t, tt.wantCwd, env.Cwd)
				assert.Equal(t, tt.wantStack, env.DirStack)
			}
			if tt.wantOut != "" {
				assert.Contains(t, out.String(), tt.wantOut)
			}
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}
		})
	}
}
