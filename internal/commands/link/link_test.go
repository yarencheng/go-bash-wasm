package link

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type MockLinkerFs struct {
	afero.Fs
	linkErr error
}

func (m *MockLinkerFs) Link(oldname, newname string) error {
	return m.linkErr
}

func TestLink_Name(t *testing.T) {
	l := New()
	assert.Equal(t, "link", l.Name())
}

func TestLink_Run(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		fs         afero.Fs
		wantStatus int
		wantErr    string
	}{
		{
			name:       "invalid arguments",
			args:       []string{"f1"},
			wantStatus: 1,
			wantErr:    "properly",
		},
		{
			name:       "link not supported",
			args:       []string{"f1", "f2"},
			fs:         afero.NewMemMapFs(),
			wantStatus: 1,
			wantErr:    "hard links not supported",
		},
		{
			name:       "link success",
			args:       []string{"f1", "f2"},
			fs:         &MockLinkerFs{Fs: afero.NewMemMapFs()},
			wantStatus: 0,
		},
		{
			name:       "link failed",
			args:       []string{"f1", "f2"},
			fs:         &MockLinkerFs{Fs: afero.NewMemMapFs(), linkErr: errors.New("perm denied")},
			wantStatus: 1,
			wantErr:    "cannot create link",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     tt.fs,
				Cwd:    "/",
				Stderr: errout,
			}

			ln := New()
			status := ln.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}
		})
	}
}
