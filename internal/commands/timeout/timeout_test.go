package timeout

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockSleep struct{}
func (m *mockSleep) Name() string { return "sleep" }
func (m *mockSleep) Run(ctx context.Context, env *commands.Environment, args []string) int {
	select {
	case <-time.After(2 * time.Second):
		return 0
	case <-ctx.Done():
		return 124
	}
}

func TestTimeout_Basic(t *testing.T) {
	registry := commands.New()
	registry.Register(&mockSleep{})

	env := &commands.Environment{
		Stdout:   &bytes.Buffer{},
		Stderr:   io.Discard,
		Registry: registry,
	}

	cmd := New()
	// Timeout after 100ms, but sleep takes 2s.
	status := cmd.Run(context.Background(), env, []string{"0.1s", "sleep"})
	assert.Equal(t, 124, status)
}
