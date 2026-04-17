package export

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestExport_Run(t *testing.T) {
	env := &commands.Environment{
		EnvVars: make(map[string]string),
	}

	e := New()
	status := e.Run(context.Background(), env, []string{"FOO=BAR"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "BAR", env.EnvVars["FOO"])
}
