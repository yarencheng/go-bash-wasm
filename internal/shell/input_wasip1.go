//go:build wasip1

package shell

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type wasmReader struct {
	scanner *bufio.Scanner
	stdout  io.Writer
}

func (w *wasmReader) Readline() (string, error) {
	fmt.Fprint(w.stdout, "$ ")
	if ok := w.scanner.Scan(); !ok {
		if err := w.scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return w.scanner.Text(), nil
}

func (w *wasmReader) Close() error {
	return nil
}

func newLineReader(env *commands.Environment) (LineReader, error) {
	return &wasmReader{
		scanner: bufio.NewScanner(env.Stdin),
		stdout:  env.Stdout,
	}, nil
}
