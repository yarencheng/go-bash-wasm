package arch

import (
	"context"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestArch_Run(t *testing.T) {
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdout: out,
	}

	a := New()
	status := a.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Equal(t, runtime.GOARCH+"\n", out.String())
}
