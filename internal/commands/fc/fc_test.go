package fc

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockExecutor struct {
	executedLine string
	retCode      int
}

func (m *mockExecutor) Execute(ctx context.Context, line string) int {
	m.executedLine = line
	return m.retCode
}

func TestFc_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		history        []string
		expectedStatus int
		containsOutput string
		executedLine   string
	}{
		{
			name:           "list history",
			args:           []string{"-l"},
			history:        []string{"ls", "echo hi", "pwd"},
			expectedStatus: 0,
			containsOutput: "2\techo hi",
		},
		{
			name:           "list history no numbers",
			args:           []string{"-ln"},
			history:        []string{"ls", "echo hi"},
			expectedStatus: 0,
			containsOutput: "echo hi",
		},
		{
			name:           "list history reverse",
			args:           []string{"-lr", "1", "2"},
			history:        []string{"cmd1", "cmd2", "cmd3"},
			expectedStatus: 0,
			containsOutput: "2\tcmd2\n1\tcmd1",
		},
		{
			name:           "re-execute last",
			args:           []string{"-s"},
			history:        []string{"ls -l", "whoami"},
			expectedStatus: 0,
			executedLine:   "whoami",
		},
		{
			name:           "re-execute by string",
			args:           []string{"-s", "ls"},
			history:        []string{"ls -l", "whoami"},
			expectedStatus: 0,
			executedLine:   "ls -l",
		},
		{
			name:           "re-execute by index",
			args:           []string{"-s", "1"},
			history:        []string{"ls -l", "whoami"},
			expectedStatus: 0,
			executedLine:   "ls -l",
		},
		{
			name:           "re-execute by negative index",
			args:           []string{"-s", "--", "-1"},
			history:        []string{"ls -l", "whoami"},
			expectedStatus: 0,
			executedLine:   "whoami",
		},
		{
			name:           "help",
			args:           []string{"--help"},
			expectedStatus: 0,
			containsOutput: "Fix Command",
		},
		{
			name:           "version",
			args:           []string{"--version"},
			expectedStatus: 0,
			containsOutput: "Version",
		},
		{
			name:           "default editor warning",
			args:           []string{},
			history:        []string{"ls"},
			expectedStatus: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			executor := &mockExecutor{}
			env := &commands.Environment{
				Stdout:   stdout,
				Stderr:   stderr,
				History:  tt.history,
				Executor: executor,
			}
			f := New()
			status := f.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.containsOutput != "" {
				assert.Contains(t, strings.ToLower(stdout.String()), strings.ToLower(tt.containsOutput))
			}
			if tt.executedLine != "" {
				assert.Equal(t, tt.executedLine, executor.executedLine)
			}
		})
	}
}

func TestFc_Metadata(t *testing.T) {
	f := New()
	assert.Equal(t, "fc", f.Name())
}
