package disown

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDisown_Run(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		initialJobs  []*commands.Job
		expectedJobs []*commands.Job
		expectedCode int
	}{
		{
			name: "disown specific job",
			args: []string{"1"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
				{ID: 2, PID: 102, Status: "Running", Command: "sleep 20"},
			},
			expectedJobs: []*commands.Job{
				{ID: 2, PID: 102, Status: "Running", Command: "sleep 20"},
			},
			expectedCode: 0,
		},
		{
			name: "disown all",
			args: []string{"-a"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Running", Command: "sleep 10"},
				{ID: 2, PID: 102, Status: "Running", Command: "sleep 20"},
			},
			expectedJobs: []*commands.Job{},
			expectedCode: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := &commands.Environment{
				Jobs:   tt.initialJobs,
				Stdout: io.Discard,
				Stderr: io.Discard,
			}
			d := New()
			status := d.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedCode, status)
			assert.Equal(t, len(tt.expectedJobs), len(env.Jobs))
		})
	}
}
