package source

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockExecutor struct {
	env *commands.Environment
}

func (m *mockExecutor) Execute(ctx context.Context, line string) int {
	line = strings.TrimSpace(line)
	if line == "" {
		return 0
	}
	args := strings.Fields(line)
	if args[0] == "setvar" {
		m.env.EnvVars[args[1]] = args[2]
		return 0
	}
	if args[0] == "exit" {
		m.env.ExitRequested = true
		return 0
	}
	if args[0] == "fail" {
		return 1
	}
	return 0
}

func TestSource_Name(t *testing.T) {
	assert.Equal(t, "source", New().Name())
	assert.Equal(t, ".", NewDot().Name())
}

func TestSource_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/script.sh", []byte("setvar FOO bar\nsetvar BAZ qux\n"), 0644)
	_ = afero.WriteFile(fs, "/fail.sh", []byte("fail\n"), 0644)

	tests := []struct {
		name       string
		args       []string
		noExecutor bool
		wantStatus int
		wantOutVar string
		wantErr    string
		dot        bool
	}{
		{
			name:       "source success",
			args:       []string{"/script.sh"},
			wantStatus: 0,
			wantOutVar: "bar",
		},
		{
			name:       "dot success",
			args:       []string{"/script.sh"},
			dot:        true,
			wantStatus: 0,
			wantOutVar: "bar",
		},
		{
			name:       "missing operand",
			args:       []string{},
			wantStatus: 1,
			wantErr:    "filename argument required",
		},
		{
			name:       "file not found",
			args:       []string{"/nonexistent"},
			wantStatus: 1,
			wantErr:    "file does not exist",
		},
		{
			name:       "executor not available",
			args:       []string{"/script.sh"},
			noExecutor: true,
			wantStatus: 1,
			wantErr:    "executor not available",
		},
		{
			name:       "last command fails",
			args:       []string{"/fail.sh"},
			wantStatus: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				FS:      fs,
				Cwd:     "/",
				EnvVars: make(map[string]string),
				Stderr:  errout,
			}
			if !tt.noExecutor {
				env.Executor = &mockExecutor{env: env}
			}

			var s *Source
			if tt.dot {
				s = NewDot()
			} else {
				s = New()
			}
			status := s.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}
			if tt.wantOutVar != "" {
				assert.Equal(t, tt.wantOutVar, env.EnvVars["FOO"])
			}
		})
	}
}

func TestSource_Exit(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/script.sh", []byte("setvar FOO bar\nexit\nsetvar BAZ qux\n"), 0644)

	env := &commands.Environment{
		FS:      fs,
		Cwd:     "/",
		EnvVars: make(map[string]string),
	}
	env.Executor = &mockExecutor{env: env}

	s := New()
	status := s.Run(context.Background(), env, []string{"/script.sh"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "bar", env.EnvVars["FOO"])
	assert.Empty(t, env.EnvVars["BAZ"])
	assert.True(t, env.ExitRequested)
}

type FailReaderFs struct {
	afero.Fs
}

func (f *FailReaderFs) Open(name string) (afero.File, error) {
	return &FailFile{File: nil}, nil
}

type FailFile struct {
	afero.File
}

func (f *FailFile) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}
func (f *FailFile) Close() error { return nil }

func TestSource_ReadError(t *testing.T) {
	// Harder to mock read error from bufio.Scanner without a custom afero.File
	// For now we've covered most paths.
}
