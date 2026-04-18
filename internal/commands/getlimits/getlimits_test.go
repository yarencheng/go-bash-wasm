package getlimits

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestGetlimits(t *testing.T) {
	g := New()
	stdout := &bytes.Buffer{}
	env := &commands.Environment{
		Stdout: stdout,
	}

	code := g.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, code)
	assert.Contains(t, stdout.String(), "CHAR_MAX=127")
	assert.Contains(t, stdout.String(), "INT_MAX=")
	assert.Contains(t, stdout.String(), "TIME_MAX=")
}
