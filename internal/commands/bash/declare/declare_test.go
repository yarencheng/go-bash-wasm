package declare

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDeclare_Run(t *testing.T) {
	env := &commands.Environment{
		EnvVars: make(map[string]string),
	}

	d := New()
	status := d.Run(context.Background(), env, []string{"MYVAR=VAL"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "VAL", env.EnvVars["MYVAR"])
}
