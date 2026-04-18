package csplit

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Csplit struct{}

func New() *Csplit {
	return &Csplit{}
}

func (c *Csplit) Name() string {
	return "csplit"
}

func (c *Csplit) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("csplit", pflag.ContinueOnError)
	prefix := flags.StringP("prefix", "f", "xx", "use PREFIX instead of 'xx'")
	suffixFormat := flags.StringP("suffix-format", "b", "%02d", "use FORMAT instead of '%02d'")
	digits := flags.IntP("digits", "n", 2, "use number of digits instead of 2")
	quiet := flags.BoolP("quiet", "s", false, "do not print counts of output file sizes")
	keepFiles := flags.BoolP("keep-files", "k", false, "do not remove output files on errors")
	elideEmpty := flags.BoolP("elide-empty-files", "z", false, "remove empty output files")
	suppressMatched := flags.Bool("suppress-matched", false, "do not echo lines matching PATTERN")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "csplit: %v\n", err)
		return 1
	}

	remainingArgs := flags.Args()
	if len(remainingArgs) < 2 {
		fmt.Fprintf(env.Stderr, "csplit: missing operand\n")
		return 1
	}

	inputPath := remainingArgs[0]
	patterns := remainingArgs[1:]

	var input io.ReadCloser
	if inputPath == "-" {
		input = io.NopCloser(env.Stdin)
	} else {
		f, err := env.FS.Open(c.absPath(env, inputPath))
		if err != nil {
			fmt.Fprintf(env.Stderr, "csplit: %v\n", err)
			return 1
		}
		input = f
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fileIdx := 0
	currentLine := 1
	var createdFiles []string

	cleanup := func() {
		if !*keepFiles {
			for _, f := range createdFiles {
				env.FS.Remove(f)
			}
		}
	}

	for _, pattern := range patterns {
		var splitLine int
		isRegex := false
		var re *regexp.Regexp

		if strings.HasPrefix(pattern, "/") && strings.HasSuffix(pattern, "/") {
			isRegex = true
			reStr := pattern[1 : len(pattern)-1]
			var err error
			re, err = regexp.Compile(reStr)
			if err != nil {
				fmt.Fprintf(env.Stderr, "csplit: invalid pattern: %v\n", err)
				cleanup()
				return 1
			}
		} else {
			val, err := strconv.Atoi(pattern)
			if err == nil {
				splitLine = val
			} else {
				fmt.Fprintf(env.Stderr, "csplit: %v: invalid number\n", pattern)
				cleanup()
				return 1
			}
		}

		if isRegex {
			found := false
			searchStart := currentLine - 1
			if searchStart < 0 {
				searchStart = 0
			}
			// For regex, if we are already at a line that matched previously,
			// we should probably start searching from the next line to find the NEXT occurrence.
			// However, the very first split search should include the current line.
			// Actually, GNU csplit: the next search starts at the line FOLLOWING the line that matched.
			if currentLine > 1 {
				searchStart = currentLine
			}

			for i := searchStart; i < len(lines); i++ {
				if re.MatchString(lines[i]) {
					splitLine = i + 1
					found = true
					break
				}
			}
			if !found {
				fmt.Fprintf(env.Stderr, "csplit: '%s': line not found\n", pattern)
				cleanup()
				return 1
			}
		}

		if splitLine <= currentLine && !isRegex {
			continue
		}

		// Handle the split
		chunk := lines[currentLine-1 : splitLine-1]
		if *elideEmpty && len(chunk) == 0 {
			// Skip empty
		} else {
			filename := fmt.Sprintf(*prefix+*suffixFormat, fileIdx)
			// Adjust format if %02d was requested but digits changed
			if *suffixFormat == "%02d" && *digits != 2 {
				filename = fmt.Sprintf("%s%0*d", *prefix, *digits, fileIdx)
			}

			absFilename := c.absPath(env, filename)
			size, err := c.writeLines(env.FS, absFilename, chunk)
			if err != nil {
				fmt.Fprintf(env.Stderr, "csplit: %v\n", err)
				cleanup()
				return 1
			}
			createdFiles = append(createdFiles, absFilename)
			if !*quiet {
				fmt.Fprintln(env.Stdout, size)
			}
			fileIdx++
		}

		if isRegex && *suppressMatched {
			currentLine = splitLine + 1
		} else {
			currentLine = splitLine
		}
	}

	// Write remaining lines
	chunk := lines[currentLine-1:]
	if !(*elideEmpty && len(chunk) == 0) {
		filename := fmt.Sprintf(*prefix+*suffixFormat, fileIdx)
		if *suffixFormat == "%02d" && *digits != 2 {
			filename = fmt.Sprintf("%s%0*d", *prefix, *digits, fileIdx)
		}
		absFilename := c.absPath(env, filename)
		size, err := c.writeLines(env.FS, absFilename, chunk)
		if err != nil {
			fmt.Fprintf(env.Stderr, "csplit: %v\n", err)
			cleanup()
			return 1
		}
		if !*quiet {
			fmt.Fprintln(env.Stdout, size)
		}
	}

	return 0
}

func (c *Csplit) writeLines(fs afero.Fs, path string, lines []string) (int64, error) {
	f, err := fs.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var total int64
	for _, line := range lines {
		n, _ := fmt.Fprintln(f, line)
		total += int64(n)
	}
	return total, nil
}

func (c *Csplit) absPath(env *commands.Environment, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(env.Cwd, path)
}
