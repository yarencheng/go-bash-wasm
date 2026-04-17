package wc

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Wc struct{}

func New() *Wc {
	return &Wc{}
}

func (w *Wc) Name() string {
	return "wc"
}

type counts struct {
	lines int
	words int
	bytes int
}

func (w *Wc) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("wc", pflag.ContinueOnError)
	l := flags.BoolP("lines", "l", false, "print the newline counts")
	words := flags.BoolP("words", "w", false, "print the word counts")
	bytes := flags.BoolP("bytes", "c", false, "print the byte counts")
	m := flags.BoolP("chars", "m", false, "print the character counts")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "wc: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		targets = []string{"-"}
	}

	// Default behavior if no flags are specified
	if !*l && !*words && !*bytes && !*m {
		*l = true
		*words = true
		*bytes = true
	}

	total := counts{}
	exitCode := 0

	for _, target := range targets {
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
				fmt.Fprintf(env.Stderr, "wc: %s: %v\n", target, err)
				exitCode = 1
				continue
			}
			defer file.Close()
			reader = file
		}

		c, err := w.count(reader)
		if err != nil {
			fmt.Fprintf(env.Stderr, "wc: %s: %v\n", target, err)
			exitCode = 1
			continue
		}

		w.printCounts(env, c, target, *l, *words, *bytes || *m)
		total.lines += c.lines
		total.words += c.words
		total.bytes += c.bytes
	}

	if len(targets) > 1 {
		w.printCounts(env, total, "total", *l, *words, *bytes || *m)
	}

	return exitCode
}

func (w *Wc) count(r io.Reader) (counts, error) {
	c := counts{}
	reader := bufio.NewReader(r)
	inWord := false

	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return c, err
		}

		c.bytes += size
		if r == '\n' {
			c.lines++
		}

		if unicode.IsSpace(r) {
			inWord = false
		} else if !inWord {
			inWord = true
			c.words++
		}
	}

	return c, nil
}

func (w *Wc) printCounts(env *commands.Environment, c counts, name string, showLines, showWords, showBytes bool) {
	var results []string
	if showLines {
		results = append(results, fmt.Sprintf("%d", c.lines))
	}
	if showWords {
		results = append(results, fmt.Sprintf("%d", c.words))
	}
	if showBytes {
		results = append(results, fmt.Sprintf("%d", c.bytes))
	}

	out := strings.Join(results, " ")
	if name != "-" {
		out += " " + name
	}
	fmt.Fprintln(env.Stdout, out)
}
