package groups

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestGroups_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		User:   "wasm",
		Groups: []int{1000},
	}

	g := New()
	status := g.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "wasm\n", stdout.String())
}
