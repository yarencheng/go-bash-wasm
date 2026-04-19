package popd

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPopd_Name(t *testing.T) {
	p := New()
	assert.Equal(t, "popd", p.Name())
}

func TestPopd_Run(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		initialCwd    string
		initialStack  []string
		wantStatus    int
		wantCwd       string
		wantStack     []string
		wantOut       string
		wantErr       string
	}{
		{
			name:         "pop basic",
			args:         []string{},
			initialCwd:   "/dir1",
			initialStack: []string{"/"},
			wantStatus:   0,
			wantCwd:      "/",
			wantStack:    []string{},
			wantOut:      "/\n",
		},
		{
			name:         "pop no-chdir",
			args:         []string{"-n"},
			initialCwd:   "/dir1",
			initialStack: []string{"/", "/dir2"},
			wantStatus:   0,
			wantCwd:      "/dir1",
			wantStack:    []string{"/dir2"},
			wantOut:      "/dir1 /dir2\n",
		},
		{
			name:         "empty stack error",
			args:         []string{},
			initialCwd:   "/",
			initialStack: []string{},
			wantStatus:   1,
			wantErr:      "directory stack empty",
		},
		{
			name:         "invalid flag",
			args:         []string{"--invalid"},
			initialCwd:   "/",
			initialStack: []string{"/dir1"},
			wantStatus:   1,
			wantErr:      "unknown flag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				Cwd:      tt.initialCwd,
				DirStack: tt.initialStack,
				Stdout:   out,
				Stderr:   errout,
			}

			p := New()
			status := p.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantStatus == 0 {
				assert.Equal(t, tt.wantCwd, env.Cwd)
				assert.Equal(t, tt.wantStack, env.DirStack)
				if tt.wantOut != "" {
					assert.Equal(t, tt.wantOut, out.String())
				}
			}
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}
		})
	}
}
