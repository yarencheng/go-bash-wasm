package kill

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestKill_List(t *testing.T) {
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdout: out,
	}

	k := New()
	status := k.Run(context.Background(), env, []string{"-l"})
	assert.Equal(t, 0, status)
	assert.Contains(t, out.String(), "TERM")
	assert.Contains(t, out.String(), "KILL")
}

func TestKill_NoSuchProcess(t *testing.T) {
	stderr := &strings.Builder{}
	env := &commands.Environment{
		Stderr: stderr,
	}

	k := New()
	status := k.Run(context.Background(), env, []string{"12345"})
	assert.Equal(t, 1, status)
	assert.Contains(t, stderr.String(), "no such process")
}
