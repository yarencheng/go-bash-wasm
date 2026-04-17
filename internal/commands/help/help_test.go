package help

import (
	"context"
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

func TestHelp_ListAll(t *testing.T) {
	registry := commands.New()
	registry.Register(&mockCommand{name: "test1"})
	registry.Register(&mockCommand{name: "test2"})
	
	out := &strings.Builder{}
	env := &commands.Environment{
		Stdout:   out,
		Registry: registry,
	}

	h := New()
	status := h.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Contains(t, out.String(), "test1")
	assert.Contains(t, out.String(), "test2")
}
