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
	maxLineLen int
}

func (w *Wc) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("wc", pflag.ContinueOnError)
	l := flags.BoolP("lines", "l", false, "print the newline counts")
	words := flags.BoolP("words", "w", false, "print the word counts")
	bytes := flags.BoolP("bytes", "c", false, "print the byte counts")
	chars := flags.BoolP("chars", "m", false, "print the character counts")
	maxLineLen := flags.BoolP("max-line-length", "L", false, "print the maximum display width")
	files0From := flags.String("files0-from", "", "read input from the files specified by NUL-terminated names in file F")
	_ = flags.String("total", "auto", "when to print a line with a total; WHEN is: auto, always, only, never")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "wc: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if *files0From != "" {
		fullPath := *files0From
		if !filepath.IsAbs(*files0From) {
			fullPath = filepath.Join(env.Cwd, *files0From)
		}
		data, err := afero.ReadFile(env.FS, fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "wc: %v\n", err)
			return 1
		}
		targets = append(targets, strings.Split(string(data), "\x00")...)
	}

	if len(targets) == 0 {
		targets = []string{"-"}
	}

	if !*l && !*words && !*bytes && !*chars && !*maxLineLen {
		*l, *words, *bytes = true, true, true
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

		w.printCounts(env, c, target, *l, *words, *bytes, *chars, *maxLineLen)
		total.lines += c.lines
		total.words += c.words
		total.bytes += c.bytes
		if c.maxLineLen > total.maxLineLen {
			total.maxLineLen = c.maxLineLen
		}
	}

	if len(targets) > 1 {
		w.printCounts(env, total, "total", *l, *words, *bytes, *chars, *maxLineLen)
	}

	return exitCode
}

func (w *Wc) count(r io.Reader) (counts, error) {
	c := counts{}
	reader := bufio.NewReader(r)
	inWord := false
	curr := 0
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
			if curr > c.maxLineLen {
				c.maxLineLen = curr
			}
			curr = 0
		} else {
			curr++
		}
		if unicode.IsSpace(r) {
			inWord = false
		} else if !inWord {
			inWord = true
			c.words++
		}
	}
	if curr > c.maxLineLen {
		c.maxLineLen = curr
	}
	return c, nil
}

func (w *Wc) printCounts(env *commands.Environment, c counts, name string, showL, showW, showB, showC, showM bool) {
	var res []string
	if showL {
		res = append(res, fmt.Sprintf("%d", c.lines))
	}
	if showW {
		res = append(res, fmt.Sprintf("%d", c.words))
	}
	if showC {
		res = append(res, fmt.Sprintf("%d", c.bytes))
	}
	if showB {
		res = append(res, fmt.Sprintf("%d", c.bytes))
	}
	if showM {
		res = append(res, fmt.Sprintf("%d", c.maxLineLen))
	}
	out := strings.Join(res, " ")
	if name != "-" {
		out += " " + name
	}
	fmt.Fprintln(env.Stdout, out)
}

