package read

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestRead_Basic(t *testing.T) {
	in := io.NopCloser(strings.NewReader("hello world\n"))
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdin:   in,
		Stdout:  out,
		EnvVars: make(map[string]string),
	}

	r := New()
	status := r.Run(context.Background(), env, []string{"VAR1", "VAR2"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello", env.EnvVars["VAR1"])
	assert.Equal(t, "world", env.EnvVars["VAR2"])
}

func TestRead_SingleVar(t *testing.T) {
	in := io.NopCloser(strings.NewReader("hello world\n"))
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdin:   in,
		Stdout:  out,
		EnvVars: make(map[string]string),
	}

	r := New()
	status := r.Run(context.Background(), env, []string{"VAR"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello world", env.EnvVars["VAR"])
}

func TestRead_DefaultVar(t *testing.T) {
	in := io.NopCloser(strings.NewReader("some data\n"))
	env := &commands.Environment{
		Stdin:   in,
		EnvVars: make(map[string]string),
	}

	r := New()
	status := r.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Equal(t, "some data", env.EnvVars["REPLY"])
}
func TestRead_Array(t *testing.T) {
	in := io.NopCloser(strings.NewReader("a b c d\n"))
	env := &commands.Environment{
		Stdin:  in,
		Arrays: make(map[string][]string),
	}

	r := New()
	status := r.Run(context.Background(), env, []string{"-a", "MYARR"})
	assert.Equal(t, 0, status)
	assert.Equal(t, []string{"a", "b", "c", "d"}, env.Arrays["MYARR"])
}

func TestRead_RawAndDelim(t *testing.T) {
	in := io.NopCloser(strings.NewReader("hello\\world:next"))
	env := &commands.Environment{
		Stdin:   in,
		EnvVars: make(map[string]string),
	}

	r := New()
	// -r preserves backslash, -d : stops at colon
	status := r.Run(context.Background(), env, []string{"-r", "-d", ":", "VAR"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello\\world", env.EnvVars["VAR"])
}

func TestRead_NChars(t *testing.T) {
	in := io.NopCloser(strings.NewReader("123456789"))
	env := &commands.Environment{
		Stdin:   in,
		EnvVars: make(map[string]string),
	}

	r := New()
	status := r.Run(context.Background(), env, []string{"-n", "3", "VAR"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "123", env.EnvVars["VAR"])
}
