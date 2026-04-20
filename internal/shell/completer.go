package shell

import (
	"path"
	"sort"
	"strings"

	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

// Completer handles tab completion logic for the shell.
type Completer struct {
	env *commands.Environment
}

// NewCompleter creates a new Completer with the given environment.
func NewCompleter(env *commands.Environment) *Completer {
	return &Completer{env: env}
}

// Do implements the readline.AutoCompleter interface.
func (c *Completer) Do(line []rune, pos int) (newLine [][]rune, length int) {
	// We only complete from the cursor position backwards.
	typed := string(line[:pos])

	if typed == "" && pos == 0 {
		return nil, 0
	}

	// Simple word splitting by space for completion.
	// In a real shell, this would need to handle quotes and escapes.
	parts := strings.Split(typed, " ")
	lastWord := parts[len(parts)-1]

	// Determine matching part for replacement length
	matchingPart := lastWord
	if strings.Contains(lastWord, "/") {
		matchingPart = path.Base(lastWord)
		if strings.HasSuffix(lastWord, "/") {
			matchingPart = ""
		}
	}
	length = len(matchingPart)

	matches := c.getMatches(typed, lastWord, parts)

	// If there's only one match and it's a command or file (not directory),
	// we append a space for convenience.
	if len(matches) == 1 {
		match := matches[0]
		// Don't append space if it's a directory (already ends in /)
		if !strings.HasSuffix(match, "/") {
			matches[0] = match + " "
		}
	}

	results := make([][]rune, len(matches))
	for i, m := range matches {
		// chzyer/readline expects only the suffix to be returned
		suffix := m
		if strings.HasPrefix(m, matchingPart) {
			suffix = m[len(matchingPart):]
		}
		results[i] = []rune(suffix)
	}

	return results, length
}

func (c *Completer) getMatches(line, lastWord string, parts []string) []string {
	var matches []string

	// Programmable completion
	cmdName := parts[0]
	if spec, ok := c.env.Completions[cmdName]; ok && len(parts) > 1 {
		matches = c.generateMatches(spec, lastWord)
	} else if len(parts) == 1 {
		// Complete commands from registry
		for _, name := range c.env.Registry.List() {
			if strings.HasPrefix(name, lastWord) {
				matches = append(matches, name)
			}
		}
	}

	if len(matches) == 0 {
		// Fallback to file completion
		matches = c.generateFileMatches(lastWord)
	}

	sort.Strings(matches)
	return matches
}

func (c *Completer) generateFileMatches(lastWord string) []string {
	var matches []string
	dir := "."
	prefix := lastWord

	if strings.Contains(lastWord, "/") {
		dir = path.Dir(lastWord)
		prefix = path.Base(lastWord)
		if strings.HasSuffix(lastWord, "/") {
			prefix = ""
		}
	}

	fullDir := dir
	if !path.IsAbs(dir) {
		fullDir = path.Join(c.env.Cwd, dir)
	}

	entries, err := afero.ReadDir(c.env.FS, fullDir)
	if err == nil {
		for _, entry := range entries {
			name := entry.Name()
			if strings.HasPrefix(name, prefix) {
				if entry.IsDir() {
					name += "/"
				}
				matches = append(matches, name)
			}
		}
	}
	return matches
}

func (c *Completer) generateMatches(spec *commands.CompSpec, lastWord string) []string {
	var matches []string

	// Actions
	if spec.Actions&(1<<2) != 0 { // -c: command
		for _, name := range c.env.Registry.List() {
			if strings.HasPrefix(name, lastWord) {
				matches = append(matches, name)
			}
		}
	}
	if spec.Actions&(1<<3) != 0 { // -d: directory
		fileMatches := c.generateFileMatches(lastWord)
		for _, m := range fileMatches {
			if strings.HasSuffix(m, "/") {
				matches = append(matches, m)
			}
		}
	}
	if spec.Actions&(1<<5) != 0 { // -f: file
		matches = append(matches, c.generateFileMatches(lastWord)...)
	}
	if spec.Actions&(1<<11) != 0 { // -v: variable
		for name := range c.env.EnvVars {
			if strings.HasPrefix(name, lastWord) {
				matches = append(matches, name)
			}
		}
	}

	// WordList
	if spec.WordList != "" {
		words := strings.Fields(spec.WordList)
		for _, word := range words {
			if strings.HasPrefix(word, lastWord) {
				matches = append(matches, word)
			}
		}
	}

	return matches
}

func commonPrefix(s1, s2 string) string {
	i := 0
	for i < len(s1) && i < len(s2) && s1[i] == s2[i] {
		i++
	}
	return s1[:i]
}
