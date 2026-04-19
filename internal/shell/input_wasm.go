//go:build wasip1 || js

package shell

import (
	"bufio"
	"fmt"
	"io"
	"path"
	"sort"
	"strings"

	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type wasmReader struct {
	env     *commands.Environment
	buf     []rune
	histIdx int
}

func (w *wasmReader) Readline() (string, error) {
	fmt.Fprint(w.env.Stdout, "$ ")
	w.buf = []rune{}
	w.histIdx = len(w.env.History)

	reader := bufio.NewReader(w.env.Stdin)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			return "", err
		}

		switch b {
		case '\r', '\n':
			fmt.Fprint(w.env.Stdout, "\r\n")
			return string(w.buf), nil

		case 127, 8: // Backspace
			if len(w.buf) > 0 {
				w.buf = w.buf[:len(w.buf)-1]
				fmt.Fprint(w.env.Stdout, "\b \b")
			}

		case '\t':
			w.handleTab()

		case 27: // Escape
			// Handle arrow keys
			if b2, _ := reader.ReadByte(); b2 == '[' {
				if b3, _ := reader.ReadByte(); b3 == 'A' { // Up
					if w.histIdx > 0 {
						w.histIdx--
						w.replaceLine(w.env.History[w.histIdx])
					}
				} else if b3 == 'B' { // Down
					if w.histIdx < len(w.env.History)-1 {
						w.histIdx++
						w.replaceLine(w.env.History[w.histIdx])
					} else if w.histIdx == len(w.env.History)-1 {
						w.histIdx++
						w.replaceLine("")
					}
				}
			}

		case 3: // Ctrl+C
			fmt.Fprint(w.env.Stdout, "^C\r\n$ ")
			w.buf = []rune{}
			w.histIdx = len(w.env.History)

		case 4: // Ctrl+D
			if len(w.buf) == 0 {
				return "", io.EOF
			}

		default:
			if b >= 32 {
				w.buf = append(w.buf, rune(b))
				fmt.Fprint(w.env.Stdout, string(b))
			}
		}
	}
}

func (w *wasmReader) handleTab() {
	line := string(w.buf)
	if line == "" {
		return
	}

	// Simple word splitting by space for completion
	parts := strings.Split(line, " ")
	lastWord := parts[len(parts)-1]

	// Programmable completion
	cmdName := parts[0]
	var matches []string
	if spec, ok := w.env.Completions[cmdName]; ok && len(parts) > 1 {
		matches = w.generateMatches(spec, lastWord)
	} else if len(parts) == 1 {
		// Complete commands
		for _, name := range w.env.Registry.List() {
			if strings.HasPrefix(name, lastWord) {
				matches = append(matches, name)
			}
		}
	}

	if len(matches) == 0 {
		// Fallback to file completion if no programmable matches or not a command
		matches = w.generateFileMatches(lastWord)
	}

	if len(matches) == 1 {
		// Unique match
		completion := matches[0]
		// Determine common prefix to know what to append
		// This is a bit simplified: handle path completion prefix logic
		prefix := lastWord
		if strings.Contains(lastWord, "/") {
			prefix = path.Base(lastWord)
			if strings.HasSuffix(lastWord, "/") {
				prefix = ""
			}
		}
		
		if strings.HasPrefix(completion, prefix) {
			completion = completion[len(prefix):]
		}

		w.buf = append(w.buf, []rune(completion)...)
		fmt.Fprint(w.env.Stdout, completion)
	} else if len(matches) > 1 {
		// Find common prefix among matches
		prefix := lastWord
		if strings.Contains(lastWord, "/") {
			prefix = path.Base(lastWord)
			if strings.HasSuffix(lastWord, "/") {
				prefix = ""
			}
		}

		common := matches[0]
		for _, m := range matches[1:] {
			common = commonPrefix(common, m)
		}
		if len(common) > len(prefix) {
			completion := common[len(prefix):]
			w.buf = append(w.buf, []rune(completion)...)
			fmt.Fprint(w.env.Stdout, completion)
		} else {
			// Show all matches
			fmt.Fprint(w.env.Stdout, "\r\n")
			sort.Strings(matches)
			for i, m := range matches {
				fmt.Fprint(w.env.Stdout, m)
				if (i+1)%4 == 0 {
					fmt.Fprint(w.env.Stdout, "\r\n")
				} else {
					fmt.Fprint(w.env.Stdout, "\t")
				}
			}
			fmt.Fprintf(w.env.Stdout, "\r\n$ %s", string(w.buf))
		}
	}
}

func (w *wasmReader) generateFileMatches(lastWord string) []string {
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
		fullDir = path.Join(w.env.Cwd, dir)
	}

	entries, err := afero.ReadDir(w.env.FS, fullDir)
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

func (w *wasmReader) generateMatches(spec *commands.CompSpec, lastWord string) []string {
	var matches []string

	// Actions
	if spec.Actions&(1<<2) != 0 { // -c: command
		for _, name := range w.env.Registry.List() {
			if strings.HasPrefix(name, lastWord) {
				matches = append(matches, name)
			}
		}
	}
	if spec.Actions&(1<<3) != 0 { // -d: directory
		fileMatches := w.generateFileMatches(lastWord)
		for _, m := range fileMatches {
			if strings.HasSuffix(m, "/") {
				matches = append(matches, m)
			}
		}
	}
	if spec.Actions&(1<<5) != 0 { // -f: file
		matches = append(matches, w.generateFileMatches(lastWord)...)
	}
	if spec.Actions&(1<<11) != 0 { // -v: variable
		for name := range w.env.EnvVars {
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

func (w *wasmReader) replaceLine(newLine string) {
	// Clear current line
	for i := 0; i < len(w.buf); i++ {
		fmt.Fprint(w.env.Stdout, "\b \b")
	}
	// Print new line
	w.buf = []rune(newLine)
	fmt.Fprint(w.env.Stdout, newLine)
}

func (w *wasmReader) Close() error {
	return nil
}

func newLineReader(env *commands.Environment) (LineReader, error) {
	return &wasmReader{
		env:     env,
		histIdx: len(env.History),
	}, nil
}
