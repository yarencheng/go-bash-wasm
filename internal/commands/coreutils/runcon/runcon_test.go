package runcon

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

func TestRuncon_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		commands       []commands.Command
		expectedStatus int
	}{
		{
			name:           "missing operand",
			args:           []string{},
			expectedStatus: 1,
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
			name:           "execution failure",
			args:           []string{"failcmd"},
			commands:       []commands.Command{&mockCmd{name: "failcmd", ret: 1}},
			expectedStatus: 1,
		},
		{
			name:           "with flags",
			args:           []string{"-u", "user", "testcmd"},
			commands:       []commands.Command{&mockCmd{name: "testcmd", ret: 0}},
			expectedStatus: 0,
		},
		{
			name:           "invalid flag",
			args:           []string{"--invalid"},
			expectedStatus: 1,
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
			r := New()
			status := r.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
		})
	}
}

func TestRuncon_RegistryNotInitialized(t *testing.T) {
	env := &commands.Environment{
		Stderr: io.Discard,
	}
	r := New()
	status := r.Run(context.Background(), env, []string{"echo"})
	assert.Equal(t, 1, status)
}

func TestRuncon_Metadata(t *testing.T) {
	r := New()
	assert.Equal(t, "runcon", r.Name())
}
