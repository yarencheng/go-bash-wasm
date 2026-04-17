package logout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestLogout_Basic(t *testing.T) {
	env := &commands.Environment{}

	l := New()
	status := l.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.True(t, env.ExitRequested)
}
