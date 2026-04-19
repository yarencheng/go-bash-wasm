package pathchk

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPathchk_Basic(t *testing.T) {
	env := &commands.Environment{
		Stderr: io.Discard,
	}

	p := New()
	status := p.Run(context.Background(), env, []string{"/valid/path"})
	assert.Equal(t, 0, status)
}

func TestPathchk_Invalid(t *testing.T) {
	// POSIX portability - includes long paths or invalid chars
	env := &commands.Environment{
		Stderr: &strings.Builder{},
	}

	p := New()
	// path is too long
	longPath := strings.Repeat("a", 5000)
	status := p.Run(context.Background(), env, []string{"-p", longPath})
	assert.Equal(t, 1, status)
}
