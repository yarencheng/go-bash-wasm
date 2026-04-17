package jobs

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestJobs_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Jobs: []*commands.Job{
			{ID: 1, PID: 123, Command: "echo hello", Status: "Running"},
			{ID: 2, PID: 124, Command: "sleep 10", Status: "Stopped"},
		},
	}

	j := New()
	status := j.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Contains(t, stdout.String(), "[1]-  Running                  echo hello")
	assert.Contains(t, stdout.String(), "[2]+  Stopped                  sleep 10")
}
