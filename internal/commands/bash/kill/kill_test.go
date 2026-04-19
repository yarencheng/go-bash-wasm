package kill

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestKill_Name(t *testing.T) {
	k := New()
	assert.Equal(t, "kill", k.Name())
}

func TestKill_Run(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		jobs       []*commands.Job
		wantStatus int
		wantOut    string
		wantErr    string
	}{
		{
			name:       "list all signals",
			args:       []string{"-l"},
			wantStatus: 0,
			wantOut:    "TERM",
		},
		{
			name:       "list specific signal by number",
			args:       []string{"-l", "15"},
			wantStatus: 0,
			wantOut:    "TERM\n",
		},
		{
			name:       "list specific signal by name",
			args:       []string{"-l", "TERM"},
			wantStatus: 0,
			wantOut:    "15\n",
		},
		{
			name:       "list invalid signal number",
			args:       []string{"-l", "999"},
			wantStatus: 0,
			wantErr:    "invalid signal number",
		},
		{
			name:       "list invalid signal name",
			args:       []string{"-l", "INVALID"},
			wantStatus: 0,
			wantErr:    "invalid signal specification",
		},
		{
			name:       "usage error no args",
			args:       []string{},
			wantStatus: 2,
			wantErr:    "usage",
		},
		{
			name:       "invalid flag",
			args:       []string{"--invalid"},
			wantStatus: 1,
			wantErr:    "unknown flag",
		},
		{
			name:       "invalid signal specification",
			args:       []string{"-s", "INVALID", "123"},
			wantStatus: 1,
			wantErr:    "invalid signal specification",
		},
		{
			name:       "invalid pid",
			args:       []string{"abc"},
			wantStatus: 1,
			wantErr:    "arguments must be process or job IDs",
		},
		{
			name:       "no such process",
			args:       []string{"12345"},
			wantStatus: 1,
			wantErr:    "no such process",
		},
		{
			name:       "kill self (pid 1)",
			args:       []string{"1"},
			wantStatus: 0,
		},
		{
			name:       "kill job by pid",
			args:       []string{"100"},
			jobs:       []*commands.Job{{PID: 100, Status: "Running"}},
			wantStatus: 0,
		},
		{
			name:       "kill job by pid with signal number",
			args:       []string{"-n", "9", "100"},
			jobs:       []*commands.Job{{PID: 100, Status: "Running"}},
			wantStatus: 0,
		},
		{
			name:       "kill job by pid with signal name",
			args:       []string{"-s", "KILL", "100"},
			jobs:       []*commands.Job{{PID: 100, Status: "Running"}},
			wantStatus: 0,
		},
		{
			name:       "stop job by pid",
			args:       []string{"-s", "STOP", "100"},
			jobs:       []*commands.Job{{PID: 100, Status: "Running"}},
			wantStatus: 0,
		},
		{
			name:       "continue job by pid",
			args:       []string{"-s", "CONT", "100"},
			jobs:       []*commands.Job{{PID: 100, Status: "Stopped"}},
			wantStatus: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout: out,
				Stderr: errout,
				Jobs:   tt.jobs,
			}

			k := New()
			status := k.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantOut != "" {
				assert.Contains(t, out.String(), tt.wantOut)
			}
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}

			// Verify job status changes
			if tt.name == "kill job by pid" || tt.name == "kill job by pid with signal name" {
				assert.Equal(t, "Done", tt.jobs[0].Status)
			} else if tt.name == "stop job by pid" {
				assert.Equal(t, "Stopped", tt.jobs[0].Status)
			} else if tt.name == "continue job by pid" {
				assert.Equal(t, "Running", tt.jobs[0].Status)
			}
		})
	}
}
