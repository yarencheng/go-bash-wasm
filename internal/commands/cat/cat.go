package cat

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Cat struct{}

func New() *Cat {
	return &Cat{}
}

func (c *Cat) Name() string {
	return "cat"
}

func (c *Cat) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("cat", pflag.ContinueOnError)
	number := flags.BoolP("number", "n", false, "number all output lines")
	numberNonBlank := flags.BoolP("number-nonblank", "b", false, "number nonempty output lines, overrides -n")
	squeezeBlank := flags.BoolP("squeeze-blank", "s", false, "suppress repeated empty output lines")
	showAll := flags.BoolP("show-all", "A", false, "equivalent to -vET")
	showEnds := flags.BoolP("show-ends", "E", false, "display $ at end of each line")
	showTabs := flags.BoolP("show-tabs", "T", false, "display TAB characters as ^I")
	showNonPrinting := flags.BoolP("show-nonprinting", "v", false, "use ^ and M- notation, except for LFD and TAB")
	eFlag := flags.BoolP("e", "e", false, "equivalent to -vE")
	tFlag := flags.BoolP("t", "t", false, "equivalent to -vT")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "cat: %v\n", err)
		return 1
	}

	if *showAll {
		*showNonPrinting = true
		*showEnds = true
		*showTabs = true
	}
	if *eFlag {
		*showNonPrinting = true
		*showEnds = true
	}
	if *tFlag {
		*showNonPrinting = true
		*showTabs = true
	}

	targets := flags.Args()
	if len(targets) == 0 {
		targets = []string{"-"}
	}
	exitCode := 0

	for _, target := range targets {
		var reader io.Reader
		if target == "-" {
			reader = env.Stdin
		} else {
			fullPath := target
			if !strings.HasPrefix(target, "/") {
				fullPath = filepath.Join(env.Cwd, target)
			}
			file, err := env.FS.Open(fullPath)
			if err != nil {
				fmt.Fprintf(env.Stderr, "cat: %s: %v\n", target, err)
				exitCode = 1
				continue
			}
			defer file.Close()
			reader = file
		}

		if !*number && !*numberNonBlank && !*squeezeBlank && !*showEnds && !*showTabs && !*showNonPrinting {
			_, err := io.Copy(env.Stdout, reader)
			if err != nil {
				fmt.Fprintf(env.Stderr, "cat: %s: %v\n", target, err)
				exitCode = 1
			}
			continue
		}

		// Handle flags using line-by-line processing
		readerComp := bufio.NewReader(reader)
		lineNumber := 1
		lastLineEmpty := false
		atEOF := false
		for !atEOF {
			line, err := readerComp.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					if line == "" {
						break
					}
					atEOF = true
				} else {
					fmt.Fprintf(env.Stderr, "cat: %s: %v\n", target, err)
					exitCode = 1
					break
				}
			}

			hasNewline := strings.HasSuffix(line, "\n")
			cleanLine := strings.TrimSuffix(line, "\n")
			isEmpty := len(cleanLine) == 0

			if *squeezeBlank && isEmpty && lastLineEmpty {
				continue
			}

			if *numberNonBlank {
				if !isEmpty {
					fmt.Fprintf(env.Stdout, "%6d\t", lineNumber)
					lineNumber++
				}
			} else if *number {
				fmt.Fprintf(env.Stdout, "%6d\t", lineNumber)
				lineNumber++
			}

			processed := cleanLine
			if *showTabs {
				processed = strings.ReplaceAll(processed, "\t", "^I")
			}
			if *showNonPrinting {
				processed = toNonPrinting(processed)
			}

			fmt.Fprint(env.Stdout, processed)

			if *showEnds && hasNewline {
				fmt.Fprint(env.Stdout, "$")
			}
			if hasNewline {
				fmt.Fprint(env.Stdout, "\n")
			}

			lastLineEmpty = isEmpty
		}
	}

	return exitCode
}

func toNonPrinting(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < 32 {
			if c == 9 || c == 10 { // TAB and LF
				b.WriteByte(c)
			} else {
				b.WriteByte('^')
				b.WriteByte(c + 64)
			}
		} else if c == 127 {
			b.WriteString("^?")
		} else if c >= 128 {
			b.WriteString("M-")
			if c >= 128+32 {
				if c == 255 {
					b.WriteString("^?")
				} else {
					b.WriteByte(c - 128)
				}
			} else {
				b.WriteByte('^')
				b.WriteByte(c - 128 + 64)
			}
		} else {
			b.WriteByte(c)
		}
	}
	return b.String()
}
