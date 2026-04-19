package builtin

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockCmd struct {
	name    string
	runArgs []string
	retCode int
}

func (m *mockCmd) Name() string { return m.name }
func (m *mockCmd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	m.runArgs = args
	return m.retCode
}

func TestBuiltin_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		registryCmds   []commands.Command
		expectedStatus int
		expectedArgs   []string
		containsStderr string
	}{
		{
			name:           "run existing builtin",
			args:           []string{"echo", "hi"},
			registryCmds:   []commands.Command{&mockCmd{name: "echo", retCode: 0}},
			expectedStatus: 0,
			expectedArgs:   []string{"hi"},
		},
		{
			name:           "not a builtin",
			args:           []string{"invalid"},
			registryCmds:   []commands.Command{},
			expectedStatus: 1,
			containsStderr: "not a shell builtin",
		},
		{
			name:           "no args",
			args:           []string{},
			expectedStatus: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			registry := commands.New()
			var targetMock *mockCmd
			for _, c := range tt.registryCmds {
				_ = registry.Register(c)
				if len(tt.args) > 0 && c.Name() == tt.args[0] {
					targetMock = c.(*mockCmd)
				}
			}

			env := &commands.Environment{
				Stdout:   stdout,
				Stderr:   stderr,
				Registry: registry,
			}

			b := New()
			status := b.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if targetMock != nil {
				assert.Equal(t, tt.expectedArgs, targetMock.runArgs)
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestBuiltin_Metadata(t *testing.T) {
	b := New()
	assert.Equal(t, "builtin", b.Name())
}
