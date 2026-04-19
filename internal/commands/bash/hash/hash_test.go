package hashcmd

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockCommand struct {
	name string
}

func (m *mockCommand) Name() string { return m.name }
func (m *mockCommand) Run(ctx context.Context, env *commands.Environment, args []string) int {
	return 0
}

func TestHash_Basic(t *testing.T) {
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdout: out,
		Stderr: io.Discard,
		Hash: map[string]string{
			"ls": "/bin/ls",
		},
	}

	h := New()
	status := h.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Contains(t, out.String(), "/bin/ls")
}

func TestHash_Clear(t *testing.T) {
	env := &commands.Environment{
		Stderr: io.Discard,
		Hash: map[string]string{
			"ls": "/bin/ls",
		},
	}

	h := New()
	status := h.Run(context.Background(), env, []string{"-r"})
	assert.Equal(t, 0, status)
	assert.Empty(t, env.Hash)
}

func TestHash_Add(t *testing.T) {
	registry := commands.New()
	registry.Register(&mockCommand{name: "ls"})
	env := &commands.Environment{
		Stderr:   io.Discard,
		Hash:     make(map[string]string),
		Registry: registry,
	}

	h := New()
	// Adding hash is usually done by shell during path lookup
	// but hash command can also 'remember' things?
	// GNU hash name... hits those names and adds to table if found.
	// Since we are sim, we'll just mock it.
	status := h.Run(context.Background(), env, []string{"ls"})
	assert.Equal(t, 0, status)
	assert.Contains(t, env.Hash, "ls")
}
