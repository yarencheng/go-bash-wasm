package stdbuf

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockCmd struct {
	name string
	ret  int
}

func (m *mockCmd) Name() string { return m.name }
func (m *mockCmd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	return m.ret
}

func TestStdbuf_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		commands       []commands.Command
		expectedStatus int
	}{
		{
			name:           "missing operand",
			args:           []string{},
			expectedStatus: 125,
		},
		{
			name:           "command not found",
			args:           []string{"echo"},
			expectedStatus: 127,
		},
		{
			name:           "successful execution",
			args:           []string{"testcmd", "arg1"},
			commands:       []commands.Command{&mockCmd{name: "testcmd", ret: 0}},
			expectedStatus: 0,
		},
		{
			name:           "with flags",
			args:           []string{"-o", "L", "testcmd"},
			commands:       []commands.Command{&mockCmd{name: "testcmd", ret: 0}},
			expectedStatus: 0,
		},
		{
			name:           "invalid flag",
			args:           []string{"--invalid"},
			expectedStatus: 125,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := commands.New()
			for _, cmd := range tt.commands {
				_ = registry.Register(cmd)
			}

			stderr := &bytes.Buffer{}
			env := &commands.Environment{
				Stderr:   stderr,
				Registry: registry,
			}
			s := New()
			status := s.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
		})
	}
}

func TestStdbuf_RegistryNotInitialized(t *testing.T) {
	env := &commands.Environment{
		Stderr: io.Discard,
	}
	s := New()
	status := s.Run(context.Background(), env, []string{"echo"})
	assert.Equal(t, 1, status)
}

func TestStdbuf_Metadata(t *testing.T) {
	s := New()
	assert.Equal(t, "stdbuf", s.Name())
}
