package shell

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestCompleter_Do_Commands(t *testing.T) {
	registry := commands.New()
	registry.Register(&mockCommand{name: "ls"})
	registry.Register(&mockCommand{name: "alias"})
	registry.Register(&mockCommand{name: "bash"})

	env := &commands.Environment{
		Registry: registry,
		FS:       afero.NewMemMapFs(),
	}

	completer := NewCompleter(env)

	// Complete 'l'
	matches, length := completer.Do([]rune("l"), 1)
	assert.Equal(t, 1, length)
	assert.Len(t, matches, 1)
	assert.Equal(t, "s ", string(matches[0]))

	// Complete 'a'
	matches, length = completer.Do([]rune("a"), 1)
	assert.Equal(t, 1, length)
	assert.Len(t, matches, 1)
	assert.Equal(t, "lias ", string(matches[0]))

	// Complete 'b'
	matches, length = completer.Do([]rune("b"), 1)
	assert.Equal(t, 1, length)
	assert.Len(t, matches, 1)
	assert.Equal(t, "ash ", string(matches[0]))
}

func TestCompleter_Do_Files(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = fs.Mkdir("/test", 0755)
	_ = afero.WriteFile(fs, "/test/file1.txt", []byte(""), 0644)
	_ = afero.WriteFile(fs, "/test/file2.txt", []byte(""), 0644)
	_ = fs.Mkdir("/test/subdir", 0755)

	env := &commands.Environment{
		Registry: commands.New(),
		FS:       fs,
		Cwd:      "/test",
	}

	completer := NewCompleter(env)

	// Complete 'f' in 'cat f'
	matches, length := completer.Do([]rune("cat f"), 5)
	assert.Equal(t, 1, length)
	assert.Len(t, matches, 2)
	assert.Equal(t, "ile1.txt", string(matches[0]))
	assert.Equal(t, "ile2.txt", string(matches[1]))

	// Complete 's' in 'ls s'
	matches, length = completer.Do([]rune("ls s"), 4)
	assert.Equal(t, 1, length)
	assert.Len(t, matches, 1)
	assert.Equal(t, "ubdir/", string(matches[0]))
}
