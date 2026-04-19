package boolcmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrue_Run(t *testing.T) {
	cmd := NewTrue()
	assert.Equal(t, "true", cmd.Name())
	status := cmd.Run(context.Background(), nil, nil)
	assert.Equal(t, 0, status)
}

func TestFalse_Run(t *testing.T) {
	cmd := NewFalse()
	assert.Equal(t, "false", cmd.Name())
	status := cmd.Run(context.Background(), nil, nil)
	assert.Equal(t, 1, status)
}
