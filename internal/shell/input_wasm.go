//go:build wasip1 || js

package shell

import (
	"bufio"
	"fmt"
	"io"

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
