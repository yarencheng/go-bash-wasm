package compgen

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type mockCmd struct {
	name string
}

func (m *mockCmd) Name() string { return m.name }
func (m *mockCmd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	return 0
}

func TestCompgen_Run(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		files          map[string]bool // path -> isDir
		commands       []string
		expectedOutput string
	}{
		{
			name:           "wordlist completion",
			args:           []string{"-W", "apple banana apricot", "ap"},
			expectedOutput: "apple\napricot\n",
		},
		{
			name:           "command completion",
			args:           []string{"-c", "l"},
			commands:       []string{"ls", "grep", "link"},
			expectedOutput: "ls\nlink\n",
		},
		{
			name:  "file completion",
			args:  []string{"-f", "te"},
			files: map[string]bool{"/test.txt": false, "/temp": true, "/other.txt": false},
			expectedOutput: "temp/\ntest.txt\n",
		},
		{
			name:  "directory completion",
			args:  []string{"-d", "te"},
			files: map[string]bool{"/test.txt": false, "/temp": true},
			expectedOutput: "temp/\n",
		},
		{
			name:  "nested file completion",
			args:  []string{"-f", "/usr/b"},
			files: map[string]bool{"/usr/bin": true, "/usr/box": false},
			expectedOutput: "/usr/bin/\n/usr/box\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			fs := afero.NewMemMapFs()
			for path, isDir := range tt.files {
				if isDir {
					_ = fs.MkdirAll(path, 0755)
				} else {
					_ = afero.WriteFile(fs, path, []byte(""), 0644)
				}
			}

			registry := commands.New()
			for _, cmdName := range tt.commands {
				_ = registry.Register(&mockCmd{name: cmdName})
			}

			env := &commands.Environment{
				Stdout:   stdout,
				Stderr:   stderr,
				FS:       fs,
				Cwd:      "/",
				Registry: registry,
			}

			c := New()
			status := c.Run(context.Background(), env, tt.args)
			assert.Equal(t, 0, status)

			// Sort lines for consistent comparison if needed, though they should be in order of discovery
			output := stdout.String()
			expectedLines := strings.Split(strings.TrimSpace(tt.expectedOutput), "\n")
			actualLines := strings.Split(strings.TrimSpace(output), "\n")
			
			// If empty, Split returns [""]
			if tt.expectedOutput == "" {
				assert.Equal(t, "", strings.TrimSpace(output))
			} else {
				assert.ElementsMatch(t, expectedLines, actualLines)
			}
		})
	}
}

func TestCompgen_Metadata(t *testing.T) {
	c := New()
	assert.Equal(t, "compgen", c.Name())
}
