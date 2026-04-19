package disown

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDisown_Name(t *testing.T) {
	d := New()
	assert.Equal(t, "disown", d.Name())
}

func TestDisown_Run(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		initialJobs   []*commands.Job
		expectedCount int
		expectedCode  int
		wantErr       string
	}{
		{
			name: "disown specific job by id",
			args: []string{"1"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
				{ID: 2, PID: 102, Status: "Running", Command: "sleep 20"},
			},
			expectedCount: 1,
			expectedCode:  0,
		},
		{
			name: "disown specific job by jobspec",
			args: []string{"%2"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
				{ID: 2, PID: 102, Status: "Running", Command: "sleep 20"},
			},
			expectedCount: 1,
			expectedCode:  0,
		},
		{
			name: "disown all",
			args: []string{"-a"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
				{ID: 2, PID: 102, Status: "Running", Command: "sleep 20"},
			},
			expectedCount: 0,
			expectedCode:  0,
		},
		{
			name: "disown all running only",
			args: []string{"-a", "-r"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
				{ID: 2, PID: 102, Status: "Stopped", Command: "sleep 20"},
			},
			expectedCount: 1,
			expectedCode:  0,
		},
		{
			name: "disown no args (last job)",
			args: []string{},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
			},
			expectedCount: 0,
			expectedCode:  0,
		},
		{
			name: "disown no args running only (last is running)",
			args: []string{"-r"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
			},
			expectedCount: 0,
			expectedCode:  0,
		},
		{
			name: "disown no args running only (last is stopped)",
			args: []string{"-r"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Stopped", Command: "sleep 10"},
			},
			expectedCount: 1,
			expectedCode:  0,
		},
		{
			name: "disown no args hup (ignored)",
			args: []string{"-h"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
			},
			expectedCount: 1,
			expectedCode:  0,
		},
		{
			name: "job not found",
			args: []string{"999"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
			},
			expectedCount: 1,
			expectedCode:  1,
			wantErr:       "no such job",
		},
		{
			name: "invalid flag",
			args: []string{"--invalid"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
			},
			expectedCount: 1,
			expectedCode:  1,
			wantErr:       "unknown flag",
		},
		{
			name: "disown -r with specific non-running job",
			args: []string{"-r", "2"},
			initialJobs: []*commands.Job{
				{ID: 2, PID: 102, Status: "Stopped", Command: "sleep 20"},
			},
			expectedCount: 1,
			expectedCode:  0,
		},
		{
			name: "disown -h with specific job",
			args: []string{"-h", "2"},
			initialJobs: []*commands.Job{
				{ID: 2, PID: 102, Status: "Running", Command: "sleep 20"},
			},
			expectedCount: 1,
			expectedCode:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				Jobs:   tt.initialJobs,
				Stderr: errout,
			}

			d := New()
			status := d.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.expectedCode, status)
			assert.Equal(t, tt.expectedCount, len(env.Jobs))
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}
		})
	}
}
