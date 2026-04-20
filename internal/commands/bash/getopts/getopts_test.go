package getopts

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestGetopts_Run(t *testing.T) {
	env := &commands.Environment{
		EnvVars: map[string]string{
			"OPTIND": "1",
		},
		Stdout: io.Discard,
		Stderr: io.Discard,
	}

	g := New()

	// Positional arguments to parse
	posArgs := []string{"-a", "-b", "val", "file"}

	// First call: should get 'a'
	status := g.Run(context.Background(), env, append([]string{"ab:", "var"}, posArgs...))
	assert.Equal(t, 0, status)
	assert.Equal(t, "a", env.EnvVars["var"])
	assert.Equal(t, "2", env.EnvVars["OPTIND"])

	// Second call: should get 'b' with OPTARG='val'
	status = g.Run(context.Background(), env, append([]string{"ab:", "var"}, posArgs...))
	assert.Equal(t, 0, status)
	assert.Equal(t, "b", env.EnvVars["var"])
	assert.Equal(t, "val", env.EnvVars["OPTARG"])
	assert.Equal(t, "4", env.EnvVars["OPTIND"])

	// Third call: should return 1 (no more options)
	status = g.Run(context.Background(), env, append([]string{"ab:", "var"}, posArgs...))
	assert.Equal(t, 1, status)
	assert.Equal(t, "?", env.EnvVars["var"])
}

func TestGetopts_Fallback(t *testing.T) {
	env := &commands.Environment{
		EnvVars: map[string]string{
			"OPTIND": "1",
		},
		PositionalArgs: []string{"-x", "extra"},
		Stdout:         io.Discard,
		Stderr:         io.Discard,
	}
	g := New()

	// No args provided, should use env.PositionalArgs
	status := g.Run(context.Background(), env, []string{"x", "var"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "x", env.EnvVars["var"])
	assert.Equal(t, "2", env.EnvVars["OPTIND"])
}
