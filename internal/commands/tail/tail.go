package tail

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Tail struct{}

func New() *Tail {
	return &Tail{}
}

func (t *Tail) Name() string {
	return "tail"
}

func (t *Tail) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("tail", pflag.ContinueOnError)
	lines := flags.IntP("lines", "n", 10, "output the last K lines, instead of the last 10")
	bytes := flags.IntP("bytes", "c", 0, "output the last K bytes")
	quiet := flags.BoolP("quiet", "q", false, "never output headers giving file names")
	verbose := flags.BoolP("verbose", "v", false, "always output headers giving file names")
	zero := flags.BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "tail: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		targets = []string{"-"}
	}

	showHeaders := (len(targets) > 1 || *verbose) && !*quiet
	exitCode := 0

	for i, target := range targets {
		var reader io.Reader
		if target == "-" {
			reader = env.Stdin
		} else {
			fullPath := target
			if !filepath.IsAbs(target) {
				fullPath = filepath.Join(env.Cwd, target)
			}
			file, err := env.FS.Open(fullPath)
			if err != nil {
				fmt.Fprintf(env.Stderr, "tail: cannot open '%s' for reading: %v\n", target, err)
				exitCode = 1
				continue
			}
			defer file.Close()
			reader = file
		}

		if showHeaders {
			if i > 0 {
				fmt.Fprintln(env.Stdout)
			}
			name := target
			if target == "-" {
				name = "standard input"
			}
			fmt.Fprintf(env.Stdout, "==> %s <==\n", name)
		}

		if *bytes > 0 {
			data, err := io.ReadAll(reader)
			if err != nil {
				fmt.Fprintf(env.Stderr, "tail: error reading '%s': %v\n", target, err)
				exitCode = 1
				continue
			}
			start := len(data) - *bytes
			if start < 0 {
				start = 0
			}
			env.Stdout.Write(data[start:])
		} else {
			var allLines []string
			scanner := bufio.NewScanner(reader)
			if *zero {
				scanner.Split(scanNull)
			}
			for scanner.Scan() {
				allLines = append(allLines, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(env.Stderr, "tail: error reading '%s': %v\n", target, err)
				exitCode = 1
				continue
			}

			count := *lines
			start := len(allLines) - count
			if start < 0 {
				start = 0
			}
			
			terminator := "\n"
			if *zero {
				terminator = "\x00"
			}
			for _, line := range allLines[start:] {
				fmt.Fprintf(env.Stdout, "%s%s", line, terminator)
			}
		}
	}

	return exitCode
}

func scanNull(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	for i := 0; i < len(data); i++ {
		if data[i] == '\x00' {
			return i + 1, data[0:i], nil
		}
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

