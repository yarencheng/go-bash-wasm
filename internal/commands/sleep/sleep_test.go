package sleep

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestSleep_Run(t *testing.T) {
	env := &commands.Environment{}

	s := New()
	start := time.Now()
	status := s.Run(context.Background(), env, []string{"0.01"})
	duration := time.Since(start)

	assert.Equal(t, 0, status)
	assert.GreaterOrEqual(t, duration, 10*time.Millisecond)
}
