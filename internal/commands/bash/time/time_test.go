package timecmd

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockExecutor struct{}

func (m *mockExecutor) Execute(ctx context.Context, line string) int {
	if line == "sleep" {
		time.Sleep(100 * time.Millisecond)
	}
	return 0
}

func TestTime_Run(t *testing.T) {
	stderr := &bytes.Buffer{}
	env := &commands.Environment{
		Stderr:   stderr,
		Executor: &mockExecutor{},
	}

	cmd := New()

	// Test timing a command
	status := cmd.Run(context.Background(), env, []string{"sleep"})
	assert.Equal(t, 0, status)
	output := stderr.String()
	assert.Contains(t, output, "real")
	assert.Contains(t, output, "0m")
	// Should be around 0.1s
	assert.True(t, strings.Contains(output, "0.1") || strings.Contains(output, "0.2"))
}
