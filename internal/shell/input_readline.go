//go:build !wasip1

package shell

import (
	"github.com/chzyer/readline"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type readlineReader struct {
	instance *readline.Instance
}

func (r *readlineReader) Readline() (string, error) {
	return r.instance.Readline()
}

func (r *readlineReader) Close() error {
	return r.instance.Close()
}

func newLineReader(env *commands.Environment) (LineReader, error) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "$ ",
		HistoryFile:     "", // Disable history file for WASM/mock simplicity
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		Stdin:           env.Stdin,
		Stdout:          env.Stdout,
		Stderr:          env.Stderr,
	})
	if err != nil {
		return nil, err
	}
	return &readlineReader{instance: rl}, nil
}
