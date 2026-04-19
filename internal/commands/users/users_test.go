package userscmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUsers_Run(t *testing.T) {
	stdout := &bytes.Buffer{}
	env := &commands.Environment{
		Stdout: stdout,
		User:   "testuser",
	}

	u := New()
	status := u.Run(context.Background(), env, nil)
	assert.Equal(t, 0, status)
	assert.Equal(t, "testuser\n", stdout.String())
}

func TestUsers_Metadata(t *testing.T) {
	u := New()
	assert.Equal(t, "users", u.Name())
}
