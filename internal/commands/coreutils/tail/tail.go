package tail

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"time"

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
	follow := flags.BoolP("follow", "f", false, "output appended data as the file grows")
	_ = flags.BoolP("retry", "F", false, "keep trying to open a file if it is inaccessible (stub)")

	_ = flags.Float64P("sleep-interval", "s", 1.0, "with -f, sleep for approximately N seconds between iterations (ignored)")
	_ = flags.Int("pid", 0, "with -f, terminate after process ID, PID dies (ignored)")
	_ = flags.Int("max-unchanged-stats", 5, "with --follow=name, reopen a FILE (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "tail: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: tail [OPTION]... [FILE]...\n")
		fmt.Fprintf(env.Stdout, "Print the last 10 lines of each FILE to standard output.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "tail")
		return 0
	}

	targets := flags.Args()
	if len(targets) == 0 {
		targets = []string{"-"}
	}

	showHeaders := (len(targets) > 1 || *verbose) && !*quiet
	exitCode := 0

	for i, target := range targets {
		var input io.Reader

		if target == "-" {
			input = env.Stdin
			t.processReader(env, input, *lines, *bytes, *zero, showHeaders, i > 0, target)
			continue
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

			// Try to seek if possible
			input = file
		}

		t.processReader(env, input, *lines, *bytes, *zero, showHeaders, i > 0, target)

		if *follow && target != "-" {
			for {
				select {
				case <-ctx.Done():
					return exitCode
				default:
					newData, err := io.ReadAll(input)
					if err != nil {
						return exitCode
					}
					if len(newData) > 0 {
						env.Stdout.Write(newData)
					}
					time.Sleep(1 * time.Second)
				}
			}
		}
	}

	return exitCode
}

func (t *Tail) processReader(env *commands.Environment, reader io.Reader, lines, bytes int, zero bool, showHeaders, interHeader bool, target string) {
	if showHeaders {
		if interHeader {
			fmt.Fprintln(env.Stdout)
		}
		name := target
		if target == "-" {
			name = "standard input"
		}
		fmt.Fprintf(env.Stdout, "==> %s <==\n", name)
	}

	if bytes > 0 {
		data, err := io.ReadAll(reader)
		if err != nil {
			fmt.Fprintf(env.Stderr, "tail: error reading '%s': %v\n", target, err)
			return
		}
		start := len(data) - bytes
		if start < 0 {
			start = 0
		}
		env.Stdout.Write(data[start:])
	} else {
		var allLines []string
		scanner := bufio.NewScanner(reader)
		if zero {
			scanner.Split(scanNull)
		}
		for scanner.Scan() {
			allLines = append(allLines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(env.Stderr, "tail: error reading '%s': %v\n", target, err)
			return
		}

		start := len(allLines) - lines
		if start < 0 {
			start = 0
		}

		terminator := "\n"
		if zero {
			terminator = "\x00"
		}
		for _, line := range allLines[start:] {
			fmt.Fprintf(env.Stdout, "%s%s", line, terminator)
		}
	}
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
