package dirname

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDirname_Run(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"basic", []string{"/usr/bin/go"}, "/usr/bin\n"},
		{"multiple", []string{"/usr/bin/go", "/etc/passwd"}, "/usr/bin\n/etc\n"},
		{"flag -z", []string{"-z", "/usr/bin/go"}, "/usr/bin\x00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			env := &commands.Environment{Stdout: &stdout}
			d := New()
			status := d.Run(context.Background(), env, tt.args)
			assert.Equal(t, 0, status)
			assert.Equal(t, tt.expected, stdout.String())
		})
	}
}
