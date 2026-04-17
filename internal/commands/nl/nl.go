package nl

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Nl struct{}

func New() *Nl {
	return &Nl{}
}

func (n *Nl) Name() string {
	return "nl"
}

func (n *Nl) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("nl", pflag.ContinueOnError)
	bodyStyle := flags.StringP("body-numbering", "b", "t", "use STYLE for numbering body lines (a: all, t: non-empty, n: none)")
	headerStyle := flags.StringP("header-numbering", "h", "n", "use STYLE for numbering header lines")
	footerStyle := flags.StringP("footer-numbering", "f", "n", "use STYLE for numbering footer lines")
	increment := flags.IntP("line-increment", "i", 1, "line number increment at each line")
	startNum := flags.IntP("starting-line-number", "v", 1, "first line number on each logical page")
	width := flags.IntP("number-width", "w", 6, "use NUMBER characters for line numbers")
	sep := flags.StringP("number-separator", "s", "\t", "add STRING after line numbers")
	format := flags.StringP("number-format", "n", "rn", "insert line numbers according to FORMAT (ln: left justify, rn: right justify, rz: right justify with leading zeros)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "nl: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return n.process(env, env.Stdin, *bodyStyle, *headerStyle, *footerStyle, *startNum, *increment, *width, *sep, *format)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "nl: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := n.process(env, f, *bodyStyle, *headerStyle, *footerStyle, *startNum, *increment, *width, *sep, *format); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

func (n *Nl) process(env *commands.Environment, r io.Reader, bStyle, hStyle, fStyle string, startNum, increment, width int, sep, format string) int {
	scanner := bufio.NewScanner(r)
	lineNum := startNum
	
	for scanner.Scan() {
		line := scanner.Text()
		
		// For now we assume everything is "body". Supporting sectional delimiters (\:\:\:, \:\:, \:) is complex.
		// GNU nl default is bodystyle='t' (number non-empty lines).
		
		shouldNumber := false
		switch bStyle {
		case "a":
			shouldNumber = true
		case "t":
			if strings.TrimSpace(line) != "" {
				shouldNumber = true
			}
		case "n":
			shouldNumber = false
		}

		if shouldNumber {
			fmt.Fprintf(env.Stdout, "%s%s%s\n", n.formatNumber(lineNum, width, format), sep, line)
			lineNum += increment
		} else {
			fmt.Fprintf(env.Stdout, "%*s%s%s\n", width, "", "", line)
		}
	}

	return 0
}

func (n *Nl) formatNumber(num, width int, format string) string {
	switch format {
	case "ln":
		return fmt.Sprintf("%-*d", width, num)
	case "rn":
		return fmt.Sprintf("%*d", width, num)
	case "rz":
		return fmt.Sprintf("%0*d", width, num)
	default:
		return fmt.Sprintf("%*d", width, num)
	}
}
