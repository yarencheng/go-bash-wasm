package bg

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestBg_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		jobs           []*commands.Job
		expectedStatus int
		expectedOutput string
		checkJobStatus func(t *testing.T, jobs []*commands.Job)
	}{
		{
			name:           "no jobs",
			args:           []string{},
			jobs:           []*commands.Job{},
			expectedStatus: 1,
		},
		{
			name: "run last job",
			args: []string{},
			jobs: []*commands.Job{
				{ID: 1, Command: "sleep 10", Status: "Stopped"},
			},
			expectedStatus: 0,
			expectedOutput: "[1]+ sleep 10 &\n",
			checkJobStatus: func(t *testing.T, jobs []*commands.Job) {
				assert.Equal(t, "Running", jobs[0].Status)
			},
		},
		{
			name: "run specific job",
			args: []string{"2"},
			jobs: []*commands.Job{
				{ID: 1, Command: "sleep 10", Status: "Stopped"},
				{ID: 2, Command: "sleep 20", Status: "Stopped"},
			},
			expectedStatus: 0,
			expectedOutput: "[2]+ sleep 20 &\n",
			checkJobStatus: func(t *testing.T, jobs []*commands.Job) {
				assert.Equal(t, "Running", jobs[1].Status)
				assert.Equal(t, "Stopped", jobs[0].Status)
			},
		},
		{
			name:           "invalid job id format",
			args:           []string{"abc"},
			jobs:           []*commands.Job{{ID: 1}},
			expectedStatus: 1,
		},
		{
			name:           "job not found",
			args:           []string{"5"},
			jobs:           []*commands.Job{{ID: 1}},
			expectedStatus: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout: stdout,
				Stderr: stderr,
				Jobs:   tt.jobs,
			}
			b := New()
			status := b.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if tt.expectedStatus == 0 {
				assert.Equal(t, tt.expectedOutput, stdout.String())
			}
			if tt.checkJobStatus != nil {
				tt.checkJobStatus(t, tt.jobs)
			}
		})
	}
}

func TestBg_Metadata(t *testing.T) {
	b := New()
	assert.Equal(t, "bg", b.Name())
}
