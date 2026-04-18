package wait

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestWait_Run(t *testing.T) {
	env := &commands.Environment{
		Jobs: []*commands.Job{
			{ID: 1, PID: 1234, Command: "sleep 10", Status: "Running"},
		},
		Stdout: io.Discard,
		Stderr: io.Discard,
	}

	cmd := New()

	// Test wait for existing PID
	status := cmd.Run(context.Background(), env, []string{"1234"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "Done", env.Jobs[0].Status)

	// Test wait for non-existent PID
	status = cmd.Run(context.Background(), env, []string{"9999"})
	assert.Equal(t, 127, status)

	// Test wait for all (no args)
	status = cmd.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
}
