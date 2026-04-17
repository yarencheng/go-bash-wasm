package basename

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestBasename_Run(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"basic", []string{"/usr/bin/go"}, "go\n"},
		{"with suffix", []string{"/usr/bin/go", "o"}, "g\n"},
		{"flag -a", []string{"-a", "/usr/bin/go", "/etc/passwd"}, "go\npasswd\n"},
		{"flag -s", []string{"-s", ".sh", "script.sh", "other.sh"}, "script\nother\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			env := &commands.Environment{Stdout: &stdout}
			b := New()
			status := b.Run(context.Background(), env, tt.args)
			assert.Equal(t, 0, status)
			assert.Equal(t, tt.expected, stdout.String())
		})
	}
}
