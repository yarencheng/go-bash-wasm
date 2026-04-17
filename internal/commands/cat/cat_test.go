package cat

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCat_Run_Flags(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		args     []string
		expected string
	}{
		{
			name:     "basic cat",
			content:  "line1\nline2\n",
			args:     []string{"test.txt"},
			expected: "line1\nline2\n",
		},
		{
			name:     "number lines",
			content:  "line1\nline2\n",
			args:     []string{"-n", "test.txt"},
			expected: "     1\tline1\n     2\tline2\n",
		},
		{
			name:     "number non-blank",
			content:  "line1\n\nline2\n",
			args:     []string{"-b", "test.txt"},
			expected: "     1\tline1\n\n     2\tline2\n",
		},
		{
			name:     "squeeze blank",
			content:  "line1\n\n\nline2\n",
			args:     []string{"-s", "test.txt"},
			expected: "line1\n\nline2\n",
		},
		{
			name:     "show ends",
			content:  "line1\nline2\n",
			args:     []string{"-E", "test.txt"},
			expected: "line1$\nline2$\n",
		},
		{
			name:     "show tabs",
			content:  "line1\tline2",
			args:     []string{"-T", "test.txt"},
			expected: "line1^Iline2",
		},
		{
			name:     "show non-printing",
			content:  "line1\x01\x1f\x7f\x80\xff",
			args:     []string{"-v", "test.txt"},
			expected: "line1^A^_^?M-^@M-^?",
		},
		{
			name:     "show all",
			content:  "line1\t\x01\n",
			args:     []string{"-A", "test.txt"},
			expected: "line1^I^A$\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte(tt.content), 0644))

			var stdout bytes.Buffer
			env := &commands.Environment{
				FS:     fs,
				Stdout: &stdout,
				Cwd:    "/",
			}

			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, 0, status)
			assert.Equal(t, tt.expected, stdout.String())
		})
	}
}
