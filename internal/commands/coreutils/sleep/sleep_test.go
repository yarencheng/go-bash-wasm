package sleep

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestSleep_Run(t *testing.T) {
	env := &commands.Environment{
		Stderr: io.Discard,
	}

	cmd := New()

	// Test basic sleep (fast)
	start := time.Now()
	status := cmd.Run(context.Background(), env, []string{"0.01"})
	assert.Equal(t, 0, status)
	assert.True(t, time.Since(start) >= 10*time.Millisecond)

	// Test suffix
	start = time.Now()
	status = cmd.Run(context.Background(), env, []string{"0.01s"})
	assert.Equal(t, 0, status)
	assert.True(t, time.Since(start) >= 10*time.Millisecond)

	// Test multiple args
	start = time.Now()
	status = cmd.Run(context.Background(), env, []string{"0.01", "0.01"})
	assert.Equal(t, 0, status)
	assert.True(t, time.Since(start) >= 20*time.Millisecond)

	// Test invalid
	status = cmd.Run(context.Background(), env, []string{"abc"})
	assert.Equal(t, 1, status)
}
