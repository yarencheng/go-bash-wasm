package history

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type History struct{}

func New() *History {
	return &History{}
}

func (h *History) Name() string {
	return "history"
}

func (h *History) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("history", pflag.ContinueOnError)
	clear := flags.BoolP("clear", "c", false, "clear the history list")
	delete := flags.IntP("delete", "d", -1, "delete the history entry at offset OFFSET")
	doAppend := flags.BoolP("append", "a", false, "append history lines from this session to the history file")
	doRead := flags.BoolP("read", "r", false, "read the history file and append the contents to the history list")
	write := flags.BoolP("write", "w", false, "write the current history to the history file")
	readNew := flags.BoolP("read-new", "n", false, "read the history lines not already read from the history file")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "history: %v\n", err)
		}
		return 1
	}

	if *clear {
		env.History = nil
		return 0
	}

	if *delete != -1 {
		idx := *delete - 1
		if idx < 0 || idx >= len(env.History) {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "history: %d: history position out of range\n", *delete)
			}
			return 1
		}
		env.History = append(env.History[:idx], env.History[idx+1:]...)
		return 0
	}

	histFile := env.EnvVars["HISTFILE"]
	if histFile == "" {
		histFile = "/home/wasm/.bash_history"
	}

	if *write {
		content := ""
		for _, line := range env.History {
			content += line + "\n"
		}
		if err := afero.WriteFile(env.FS, histFile, []byte(content), 0644); err != nil {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "history: %v\n", err)
			}
			return 1
		}
		return 0
	}

	if *doAppend {
		// For now, same as write since we don't track session-only history separately yet
		f, err := env.FS.OpenFile(histFile, 1|64, 0644) // O_WRONLY|O_APPEND|O_CREATE
		if err != nil {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "history: %v\n", err)
			}
			return 1
		}
		defer f.Close()
		for _, line := range env.History {
			fmt.Fprintln(f, line)
		}
		return 0
	}

	if *doRead || *readNew {
		data, err := afero.ReadFile(env.FS, histFile)
		if err != nil {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "history: %v\n", err)
			}
			return 1
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if line != "" {
				env.History = append(env.History, line)
			}
		}
		return 0
	}

	targets := flags.Args()
	limit := len(env.History)
	if len(targets) > 0 {
		n, err := strconv.Atoi(targets[0])
		if err == nil && n < limit {
			limit = n
		}
	}

	start := len(env.History) - limit
	if start < 0 {
		start = 0
	}

	for i := start; i < len(env.History); i++ {
		fmt.Fprintf(env.Stdout, "%5d  %s\n", i+1, env.History[i])
	}

	return 0
}
