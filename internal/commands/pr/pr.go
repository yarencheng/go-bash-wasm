package pr

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Pr struct{}

func New() *Pr {
	return &Pr{}
}

func (p *Pr) Name() string {
	return "pr"
}

func (p *Pr) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("pr", pflag.ContinueOnError)
	header := flags.StringP("header", "h", "", "use STRING as the header")
	length := flags.IntP("length", "l", 66, "page length in lines")
	omitHeader := flags.BoolP("omit-header", "t", false, "omit page headers and footers")
	width := flags.IntP("width", "w", 72, "page width in characters")
	doubleSpace := flags.BoolP("double-space", "d", false, "double space the output")
	numbering := flags.BoolP("numbering", "n", false, "number lines")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "pr: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return p.process(env, env.Stdin, "stdin", *header, *length, *omitHeader, *width, *doubleSpace, *numbering)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "pr: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := p.process(env, f, arg, *header, *length, *omitHeader, *width, *doubleSpace, *numbering); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

func (p *Pr) process(env *commands.Environment, r io.Reader, filename, customHeader string, length int, omitHeader bool, width int, doubleSpace, numbering bool) int {
	scanner := bufio.NewScanner(r)
	page := 1
	lineCount := 0
	
	printHeader := func() {
		if !omitHeader {
			dateStr := time.Now().Format("2006-01-02 15:04")
			h := filename
			if customHeader != "" {
				h = customHeader
			}
			fmt.Fprintf(env.Stdout, "\n\n%s  %s  Page %d\n\n\n", dateStr, h, page)
			lineCount = 5
		} else {
			lineCount = 0
		}
	}

	printHeader()
	globalLineNum := 1
	
	for scanner.Scan() {
		line := scanner.Text()
		
		if lineCount >= length-5 && !omitHeader {
			// Print footer (5 blank lines)
			for i := 0; i < 5; i++ {
				fmt.Fprintln(env.Stdout)
			}
			page++
			printHeader()
		}

		outLine := line
		if numbering {
			outLine = fmt.Sprintf("%5d\t%s", globalLineNum, outLine)
		}
		
		fmt.Fprintln(env.Stdout, outLine)
		lineCount++
		globalLineNum++

		if doubleSpace {
			fmt.Fprintln(env.Stdout)
			lineCount++
		}
	}

	if !omitHeader {
		// Fill the rest of the page
		for lineCount < length {
			fmt.Fprintln(env.Stdout)
			lineCount++
		}
	}

	return 0
}
