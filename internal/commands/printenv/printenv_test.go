package printenv

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPrintenv_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		EnvVars: map[string]string{
			"FOO": "BAR",
		},
	}

	p := New()
	status := p.Run(context.Background(), env, []string{"FOO"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "BAR\n", stdout.String())
}
