package head

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Head struct{}

func New() *Head {
	return &Head{}
}

func (h *Head) Name() string {
	return "head"
}

func (h *Head) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("head", pflag.ContinueOnError)
	lines := flags.IntP("lines", "n", 10, "print the first K lines instead of the first 10")
	bytes := flags.IntP("bytes", "c", 0, "print the first K bytes of each file")
	quiet := flags.BoolP("quiet", "q", false, "never print headers giving file names")
	verbose := flags.BoolP("verbose", "v", false, "always print headers giving file names")
	zero := flags.BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "head: %v\n", err)
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
				fmt.Fprintf(env.Stderr, "head: cannot open '%s' for reading: %v\n", target, err)
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
			_, err := io.CopyN(env.Stdout, reader, int64(*bytes))
			if err != nil && err != io.EOF {
				fmt.Fprintf(env.Stderr, "head: error reading '%s': %v\n", target, err)
				exitCode = 1
			}
		} else {
			count := *lines
			scanner := bufio.NewScanner(reader)
			if *zero {
				scanner.Split(scanNull)
			}
			terminator := "\n"
			if *zero {
				terminator = "\x00"
			}
			for j := 0; j < count && scanner.Scan(); j++ {
				fmt.Fprintf(env.Stdout, "%s%s", scanner.Text(), terminator)
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(env.Stderr, "head: error reading '%s': %v\n", target, err)
				exitCode = 1
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

