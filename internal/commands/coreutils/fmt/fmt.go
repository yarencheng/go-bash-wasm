package fmt

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Fmt struct{}

func New() *Fmt {
	return &Fmt{}
}

func (f *Fmt) Name() string {
	return "fmt"
}

func (f *Fmt) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("fmt", pflag.ContinueOnError)
	width := flags.StringP("width", "w", "75", "maximum line width")
	splitOnly := flags.BoolP("split-only", "s", false, "split long lines, but do not refill")

	_ = flags.BoolP("crown-margin", "c", false, "preserve indentation of first two lines (ignored)")
	_ = flags.StringP("prefix", "p", "", "reformat only lines beginning with STRING (ignored)")
	_ = flags.BoolP("tagged-paragraph", "t", false, "indentation of first line different from second (ignored)")
	_ = flags.BoolP("uniform-spacing", "u", false, "one space between words, two after sentences (ignored)")
	_ = flags.StringP("goal", "g", "", "goal width (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	// Pre-parse numeric flags like -75
	var filtered []string
	for _, arg := range args {
		if len(arg) > 1 && arg[0] == '-' && arg[1] >= '0' && arg[1] <= '9' {
			if _, err := strconv.Atoi(arg[1:]); err == nil {
				filtered = append(filtered, "-w", arg[1:])
				continue
			}
		}
		filtered = append(filtered, arg)
	}

	if err := flags.Parse(filtered); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "fmt: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: fmt [-WIDTH] [OPTION]... [FILE]...\n")
		fmt.Fprintf(env.Stdout, "Reformat each paragraph in the FILE(s), writing to standard output.\n")
		fmt.Fprintf(env.Stdout, "The default width is 75 columns.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "fmt")
		return 0
	}

	maxWidth := 75
	if *width != "" {
		if val, err := strconv.Atoi(*width); err == nil {
			maxWidth = val
		}
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return f.process(env, env.Stdin, maxWidth, *splitOnly)
	}

	exitCode := 0
	for _, arg := range remaining {
		file, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "fmt: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := f.process(env, file, maxWidth, *splitOnly); status != 0 {
			exitCode = status
		}
		file.Close()
	}

	return exitCode
}

func (f *Fmt) process(env *commands.Environment, r io.Reader, maxWidth int, splitOnly bool) int {
	scanner := bufio.NewScanner(r)
	var currentParagraph []string

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if trimmed == "" {
			// End of paragraph
			if len(currentParagraph) > 0 {
				f.flushParagraph(env, currentParagraph, maxWidth, splitOnly)
				currentParagraph = nil
			}
			fmt.Fprintln(env.Stdout)
		} else {
			if splitOnly {
				// Just split long lines, don't combine
				f.splitLine(env, line, maxWidth)
			} else {
				// Collect for paragraph reconstruction
				currentParagraph = append(currentParagraph, trimmed)
			}
		}
	}

	if len(currentParagraph) > 0 {
		f.flushParagraph(env, currentParagraph, maxWidth, splitOnly)
	}

	return 0
}

func (f *Fmt) flushParagraph(env *commands.Environment, lines []string, maxWidth int, splitOnly bool) {
	if splitOnly {
		for _, line := range lines {
			f.splitLine(env, line, maxWidth)
		}
		return
	}

	// Combine all lines into a single stream of words
	var words []string
	for _, line := range lines {
		words = append(words, strings.Fields(line)...)
	}

	if len(words) == 0 {
		return
	}

	var currentLine strings.Builder
	for _, word := range words {
		if currentLine.Len() == 0 {
			currentLine.WriteString(word)
		} else if currentLine.Len()+1+len(word) <= maxWidth {
			currentLine.WriteString(" ")
			currentLine.WriteString(word)
		} else {
			fmt.Fprintln(env.Stdout, currentLine.String())
			currentLine.Reset()
			currentLine.WriteString(word)
		}
	}

	if currentLine.Len() > 0 {
		fmt.Fprintln(env.Stdout, currentLine.String())
	}
}

func (f *Fmt) splitLine(env *commands.Environment, line string, maxWidth int) {
	if len(line) <= maxWidth {
		fmt.Fprintln(env.Stdout, line)
		return
	}

	words := strings.Fields(line)
	var currentLine strings.Builder
	for _, word := range words {
		if currentLine.Len() == 0 {
			currentLine.WriteString(word)
		} else if currentLine.Len()+1+len(word) <= maxWidth {
			currentLine.WriteString(" ")
			currentLine.WriteString(word)
		} else {
			fmt.Fprintln(env.Stdout, currentLine.String())
			currentLine.Reset()
			currentLine.WriteString(word)
		}
	}
	if currentLine.Len() > 0 {
		fmt.Fprintln(env.Stdout, currentLine.String())
	}
}
