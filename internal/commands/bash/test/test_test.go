package testcmd

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTest_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello"), 0644))

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}
	// Test cases
	tests := []struct {
		name     string
		args     []string
		expected int
	}{
		{"-e exists", []string{"-e", "/test.txt"}, 0},
		{"-e not exists", []string{"-e", "/nonexistent"}, 1},
		{"-f is file", []string{"-f", "/test.txt"}, 0},
		{"-d is not directory", []string{"-d", "/test.txt"}, 1},
		{"-z empty", []string{"-z", ""}, 0},
		{"-z not empty", []string{"-z", "abc"}, 1},
		{"-n not empty", []string{"-n", "abc"}, 0},
		{"-n empty", []string{"-n", ""}, 1},
		{"string equal", []string{"abc", "=", "abc"}, 0},
		{"string not equal", []string{"abc", "=", "def"}, 1},
		{"string != true", []string{"abc", "!=", "def"}, 0},
		{"string != false", []string{"abc", "!=", "abc"}, 1},
		{"num -eq true", []string{"10", "-eq", "10"}, 0},
		{"num -eq false", []string{"10", "-eq", "20"}, 1},
		{"num -ne true", []string{"10", "-ne", "20"}, 0},
		{"num -lt true", []string{"10", "-lt", "20"}, 0},
		{"num -le true", []string{"10", "-le", "10"}, 0},
		{"num -gt true", []string{"20", "-gt", "10"}, 0},
		{"num -ge true", []string{"10", "-ge", "10"}, 0},
		{"unary negation", []string{"!", "-z", "abc"}, 0},
		{"binary negation", []string{"!", "10", "-eq", "20"}, 0},
		{"logical and", []string{"10", "-eq", "10", "-a", "20", "-eq", "20"}, 0},
		{"logical or", []string{"10", "-eq", "20", "-o", "20", "-eq", "20"}, 0},
		{"[ ] alias", []string{"10", "-eq", "10", "]"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdName := "test"
			if tt.name == "[ ] alias" {
				cmdName = "["
			}
			tc := New(cmdName)
			status := tc.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expected, status)
		})
	}
}
