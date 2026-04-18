package paste

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Paste struct{}

func New() *Paste {
	return &Paste{}
}

func (p *Paste) Name() string {
	return "paste"
}

func (p *Paste) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("paste", pflag.ContinueOnError)
	delims := flags.StringP("delimiters", "d", "\t", "reuse characters from LIST instead of TABs")
	serial := flags.BoolP("serial", "s", false, "paste one file at a time instead of in parallel")
	zero := flags.BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "paste: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return p.processParallel(env, []io.ReadCloser{io.NopCloser(env.Stdin)}, *delims, *zero)
	}

	files := make([]io.ReadCloser, 0, len(remaining))
	for _, arg := range remaining {
		if arg == "-" {
			files = append(files, io.NopCloser(env.Stdin))
			continue
		}
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "paste: %s: %v\n", arg, err)
			for _, pf := range files {
				pf.Close()
			}
			return 1
		}
		files = append(files, f)
	}
	defer func() {
		for _, f := range files {
			f.Close()
		}
	}()

	if *serial {
		return p.processSerial(env, files, *delims, *zero)
	}
	return p.processParallel(env, files, *delims, *zero)
}

func (p *Paste) processParallel(env *commands.Environment, files []io.ReadCloser, delims string, zero bool) int {
	scanners := make([]*bufio.Scanner, len(files))
	for i, f := range files {
		scanners[i] = bufio.NewScanner(f)
		if zero {
			scanners[i].Split(scanNull)
		}
	}

	terminator := "\n"
	if zero {
		terminator = "\x00"
	}

	delimChars := []rune(delims)

	for {
		active := false
		var builder strings.Builder
		for i, s := range scanners {
			if i > 0 {
				builder.WriteRune(delimChars[(i-1)%len(delimChars)])
			}
			if s.Scan() {
				builder.WriteString(s.Text())
				active = true
			}
		}

		if !active {
			break
		}
		fmt.Fprintf(env.Stdout, "%s%s", builder.String(), terminator)
	}

	return 0
}

func (p *Paste) processSerial(env *commands.Environment, files []io.ReadCloser, delims string, zero bool) int {
	terminator := "\n"
	if zero {
		terminator = "\x00"
	}

	delimChars := []rune(delims)

	for _, f := range files {
		scanner := bufio.NewScanner(f)
		if zero {
			scanner.Split(scanNull)
		}

		first := true
		i := 0
		for scanner.Scan() {
			if !first {
				fmt.Fprintf(env.Stdout, "%c", delimChars[i%len(delimChars)])
				i++
			}
			fmt.Fprint(env.Stdout, scanner.Text())
			first = false
		}
		fmt.Fprint(env.Stdout, terminator)
	}
	return 0
}

func scanNull(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.IndexByte(string(data), '\x00'); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
