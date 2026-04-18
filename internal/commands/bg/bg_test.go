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
		name         string
		args         []string
		initialJobs  []*commands.Job
		expectedStatus string
		expectedCode int
	}{
		{
			name: "bg specific job",
			args: []string{"1"},
			initialJobs: []*commands.Job{
				{ID: 1, PID: 101, Status: "Stopped", Command: "sleep 10"},
			},
			expectedStatus: "Running",
			expectedCode: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			env := &commands.Environment{
				Jobs:   tt.initialJobs,
				Stdout: &stdout,
				Stderr: &stderr,
			}
			b := New()
			status := b.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedCode, status)
			assert.Equal(t, tt.expectedStatus, env.Jobs[0].Status)
		})
	}
}
