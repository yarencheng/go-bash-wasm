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
	noRenumber := flags.BoolP("no-renumber", "p", false, "do not reset line numbers at logical pages")
	joinBlank := flags.IntP("join-blank-lines", "l", 1, "group of NUMBER empty lines counted as one")
	delim := flags.StringP("section-delimiter", "d", "\\:", "use STRING to separate logical pages")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "nl: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	opts := nlOptions{
		bStyle: *bodyStyle, hStyle: *headerStyle, fStyle: *footerStyle,
		startNum: *startNum, increment: *increment, width: *width,
		sep: *sep, format: *format, noRenumber: *noRenumber,
		joinBlank: *joinBlank, delim: *delim,
	}

	if len(remaining) == 0 {
		return n.process(env, env.Stdin, opts)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "nl: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := n.process(env, f, opts); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

type nlOptions struct {
	bStyle, hStyle, fStyle string
	startNum, increment    int
	width                  int
	sep, format            string
	noRenumber             bool
	joinBlank              int
	delim                  string
}

func (n *Nl) process(env *commands.Environment, r io.Reader, opts nlOptions) int {
	scanner := bufio.NewScanner(r)
	lineNum := opts.startNum
	section := "body" // default
	blankCount := 0

	d1 := ":"
	d2 := ":"
	if len(opts.delim) >= 1 { d1 = string(opts.delim[0]) }
	if len(opts.delim) >= 2 { d2 = string(opts.delim[1]) }

	pair := d1 + d2
	hDelim := pair + pair + pair
	bDelim := pair + pair
	fDelim := pair

	for scanner.Scan() {
		line := scanner.Text()
		
		isDelim := false
		if line == hDelim {
			section = "header"
			if !opts.noRenumber { lineNum = opts.startNum }
			isDelim = true
		} else if line == bDelim {
			section = "body"
			if !opts.noRenumber { lineNum = opts.startNum }
			isDelim = true
		} else if line == fDelim {
			section = "footer"
			if !opts.noRenumber { lineNum = opts.startNum }
			isDelim = true
		}

		if isDelim {
			fmt.Fprintln(env.Stdout)
			continue
		}

		style := opts.bStyle
		if section == "header" { style = opts.hStyle }
		if section == "footer" { style = opts.fStyle }

		shouldNumber := false
		isEmpty := strings.TrimSpace(line) == ""
		
		if isEmpty {
			blankCount++
			if blankCount >= opts.joinBlank {
				blankCount = 0
				if style == "a" || style == "t" { // GNU nl: t numbers non-empty, but if joined blanks reach threshold it might number? 
					// Actually 't' only numbers non-empty. 'a' numbers all.
				}
			}
		} else {
			blankCount = 0
		}

		switch style {
		case "a":
			shouldNumber = true
		case "t":
			if !isEmpty {
				shouldNumber = true
			}
		case "n":
			shouldNumber = false
		}

		if shouldNumber {
			fmt.Fprintf(env.Stdout, "%s%s%s\n", n.formatNumber(lineNum, opts.width, opts.format), opts.sep, line)
			lineNum += opts.increment
		} else {
			fmt.Fprintf(env.Stdout, "%*s%s%s\n", opts.width, "", "", line)
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
