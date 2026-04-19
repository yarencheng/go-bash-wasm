package testcmd

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTest_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/test.txt", []byte("hello"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/empty.txt", []byte(""), 0644))
	
	// Newer and older files
	require.NoError(t, afero.WriteFile(fs, "/old.txt", []byte("old"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/new.txt", []byte("new"), 0644))
	
	now := time.Now()
	require.NoError(t, fs.Chtimes("/old.txt", now.Add(-time.Hour), now.Add(-time.Hour)))
	require.NoError(t, fs.Chtimes("/new.txt", now, now))

	symlinkSupported := false
	if _, ok := fs.(afero.Linker); ok {
		err := fs.(afero.Linker).SymlinkIfPossible("/test.txt", "/link.txt")
		if err == nil {
			symlinkSupported = true
		}
	}

	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
		EnvVars: map[string]string{
			"VAR_EXIST": "1",
		},
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
		{"-s non-empty", []string{"-s", "/test.txt"}, 0},
		{"-s empty", []string{"-s", "/empty.txt"}, 1},
		{"-L symlink", []string{"-L", "/link.txt"}, 0},
		{"-h symlink", []string{"-h", "/link.txt"}, 0},
		{"-v var set", []string{"-v", "VAR_EXIST"}, 0},
		{"-v var not set", []string{"-v", "VAR_NOT_EXIST"}, 1},
		{"-nt newer", []string{"/new.txt", "-nt", "/old.txt"}, 0},
		{"-ot older", []string{"/old.txt", "-ot", "/new.txt"}, 0},
		{"-ef same file", []string{"/test.txt", "-ef", "/test.txt"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (tt.name == "-L symlink" || tt.name == "-h symlink") && !symlinkSupported {
				t.Skip("Symlinks not supported by FS")
			}
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
