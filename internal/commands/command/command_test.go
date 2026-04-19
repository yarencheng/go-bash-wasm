package command

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

func TestCommand_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		registryCmds   []commands.Command
		expectedStatus int
		expectedArgs   []string
		containsOutput string
		containsStderr string
	}{
		{
			name:           "execute builtin",
			args:           []string{"ls", "-l"},
			registryCmds:   []commands.Command{&mockCmd{name: "ls", retCode: 0}},
			expectedStatus: 0,
			expectedArgs:   []string{"-l"},
		},
		{
			name:           "verbose identify",
			args:           []string{"-v", "ls"},
			registryCmds:   []commands.Command{&mockCmd{name: "ls"}},
			expectedStatus: 0,
			containsOutput: "ls",
		},
		{
			name:           "Verbose identify",
			args:           []string{"-V", "ls"},
			registryCmds:   []commands.Command{&mockCmd{name: "ls"}},
			expectedStatus: 0,
			containsOutput: "ls is a builtin command",
		},
		{
			name:           "not found",
			args:           []string{"unknown"},
			expectedStatus: 127,
			containsStderr: "not found",
		},
		{
			name:           "identify not found",
			args:           []string{"-v", "unknown"},
			expectedStatus: 1,
			containsStderr: "not found",
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
				if mc, ok := c.(*mockCmd); ok {
					targetMock = mc
				}
			}

			env := &commands.Environment{
				Stdout:   stdout,
				Stderr:   stderr,
				Registry: registry,
				EnvVars:  make(map[string]string),
			}

			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)
			if targetMock != nil && tt.name != "verbose identify" && tt.name != "Verbose identify" {
				assert.Equal(t, tt.expectedArgs, targetMock.runArgs)
			}
			if tt.containsOutput != "" {
				assert.Contains(t, stdout.String(), tt.containsOutput)
			}
			if tt.containsStderr != "" {
				assert.Contains(t, stderr.String(), tt.containsStderr)
			}
		})
	}
}

func TestCommand_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "command", c.Name())
}
