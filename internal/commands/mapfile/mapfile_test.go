package mapfile

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestMapfile_Basic(t *testing.T) {
	input := "line1\nline2\nline3\n"
	env := &commands.Environment{
		Stdin:  io.NopCloser(strings.NewReader(input)),
		Arrays: make(map[string][]string),
	}

	m := New()
	status := m.Run(context.Background(), env, []string{"myarray"})
	assert.Equal(t, 0, status)
	assert.Equal(t, []string{"line1\n", "line2\n", "line3\n"}, env.Arrays["myarray"])
}

func TestMapfile_Trim(t *testing.T) {
	input := "line1\nline2\n"
	env := &commands.Environment{
		Stdin:  io.NopCloser(strings.NewReader(input)),
		Arrays: make(map[string][]string),
	}

	m := New()
	status := m.Run(context.Background(), env, []string{"-t", "myarray"})
	assert.Equal(t, 0, status)
	assert.Equal(t, []string{"line1", "line2"}, env.Arrays["myarray"])
}
