package grep

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"regexp"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Grep struct{}

func New() *Grep {
	return &Grep{}
}

func (g *Grep) Name() string {
	return "grep"
}

func (g *Grep) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("grep", pflag.ContinueOnError)
	ignoreCase := flags.BoolP("ignore-case", "i", false, "ignore case distinctions")
	invertMatch := flags.BoolP("invert-match", "v", false, "select non-matching lines")
	lineReg := flags.BoolP("line-number", "n", false, "print line number with output lines")
	count := flags.BoolP("count", "c", false, "print only a count of matching lines per file")
	filesWithMatches := flags.BoolP("files-with-matches", "l", false, "print only names of files with matches")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "grep: %v\n", err)
		return 2
	}

	remaining := flags.Args()
	if len(remaining) < 1 {
		fmt.Fprintf(env.Stderr, "grep: pattern is required\n")
		return 2
	}

	pattern := remaining[0]
	files := remaining[1:]

	regFlags := ""
	if *ignoreCase {
		regFlags = "(?i)"
	}

	re, err := regexp.Compile(regFlags + pattern)
	if err != nil {
		fmt.Fprintf(env.Stderr, "grep: invalid regex: %v\n", err)
		return 2
	}

	status := 1 // Default to no match

	if len(files) == 0 {
		// Read from stdin
		matched, _ := g.processReader(env.Stdout, env.Stdin, re, "stdin", *invertMatch, *lineReg, *count, *filesWithMatches, false)
		if matched {
			status = 0
		}
	} else {
		multipleFiles := len(files) > 1
		for _, file := range files {
			fullPath := file
			if !filepath.IsAbs(file) {
				fullPath = filepath.Join(env.Cwd, file)
			}

			f, err := env.FS.Open(fullPath)
			if err != nil {
				fmt.Fprintf(env.Stderr, "grep: %s: %v\n", file, err)
				continue
			}

			matched, _ := g.processReader(env.Stdout, f, re, file, *invertMatch, *lineReg, *count, *filesWithMatches, multipleFiles)
			f.Close()
			if matched {
				status = 0
			}
		}
	}

	return status
}

func (g *Grep) processReader(stdout io.Writer, r io.Reader, re *regexp.Regexp, filename string, invert, lineNum, countOnly, listFiles, showFilename bool) (bool, error) {
	scanner := bufio.NewScanner(r)
	matchCount := 0
	matchedAny := false

	lineIdx := 0
	for scanner.Scan() {
		lineIdx++
		text := scanner.Text()
		match := re.MatchString(text)
		if invert {
			match = !match
		}

		if match {
			matchCount++
			matchedAny = true
			if !countOnly && !listFiles {
				prefix := ""
				if showFilename {
					prefix = filename + ":"
				}
				if lineNum {
					prefix = fmt.Sprintf("%s%d:", prefix, lineIdx)
				}
				fmt.Fprintf(stdout, "%s%s\n", prefix, text)
			}
		}
	}

	if listFiles && matchedAny {
		fmt.Fprintln(stdout, filename)
	} else if countOnly {
		prefix := ""
		if showFilename {
			prefix = filename + ":"
		}
		fmt.Fprintf(stdout, "%s%d\n", prefix, matchCount)
	}

	return matchedAny, scanner.Err()
}
