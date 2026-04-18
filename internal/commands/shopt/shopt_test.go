package shopt

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestShopt(t *testing.T) {
	s := New()
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	env := &commands.Environment{
		Stdout: stdout,
		Stderr: stderr,
		Shopts: map[string]bool{
			"testopt": false,
			"setopt":  true,
		},
	}

	t.Run("list all", func(t *testing.T) {
		stdout.Reset()
		stderr.Reset()
		code := s.Run(context.Background(), env, []string{})
		assert.Equal(t, 0, code)
		assert.Contains(t, stdout.String(), "testopt        	off")
		assert.Contains(t, stdout.String(), "setopt         	on")
	})

	t.Run("set option", func(t *testing.T) {
		stdout.Reset()
		stderr.Reset()
		code := s.Run(context.Background(), env, []string{"-s", "testopt"})
		assert.Equal(t, 0, code)
		assert.True(t, env.Shopts["testopt"])
	})

	t.Run("unset option", func(t *testing.T) {
		stdout.Reset()
		stderr.Reset()
		code := s.Run(context.Background(), env, []string{"-u", "setopt"})
		assert.Equal(t, 0, code)
		assert.False(t, env.Shopts["setopt"])
	})

	t.Run("invalid option", func(t *testing.T) {
		stdout.Reset()
		stderr.Reset()
		code := s.Run(context.Background(), env, []string{"-s", "invalid"})
		assert.Equal(t, 1, code)
		assert.Contains(t, stderr.String(), "invalid shell option name")
	})

	t.Run("print mode", func(t *testing.T) {
		stdout.Reset()
		stderr.Reset()
		env.Shopts["testopt"] = true
		code := s.Run(context.Background(), env, []string{"-p", "testopt"})
		assert.Equal(t, 0, code)
		assert.Equal(t, "shopt -s testopt\n", stdout.String())
	})

	t.Run("quiet mode enabled", func(t *testing.T) {
		stdout.Reset()
		stderr.Reset()
		env.Shopts["setopt"] = true
		code := s.Run(context.Background(), env, []string{"-q", "setopt"})
		assert.Equal(t, 0, code)
		assert.Empty(t, stdout.String())
	})

	t.Run("quiet mode disabled", func(t *testing.T) {
		stdout.Reset()
		stderr.Reset()
		env.Shopts["testopt"] = false
		code := s.Run(context.Background(), env, []string{"-q", "testopt"})
		assert.Equal(t, 1, code)
		assert.Empty(t, stdout.String())
	})

	t.Run("cannot set and unset simultaneously", func(t *testing.T) {
		stdout.Reset()
		stderr.Reset()
		code := s.Run(context.Background(), env, []string{"-s", "-u", "testopt"})
		assert.Equal(t, 1, code)
		assert.Contains(t, stderr.String(), "cannot set and unset shell options simultaneously")
	})
}
