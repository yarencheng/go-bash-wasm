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
	
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "cat: %v\n", err)
		return 1
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

		if !*number && !*numberNonBlank && !*squeezeBlank {
			_, err := io.Copy(env.Stdout, reader)
			if err != nil {
				fmt.Fprintf(env.Stderr, "cat: %s: %v\n", target, err)
				exitCode = 1
			}
			continue
		}

		// Handle flags using line-by-line processing
		scanner := bufio.NewScanner(reader)
		lineNumber := 1
		lastLineEmpty := false
		for scanner.Scan() {
			line := scanner.Text()
			isEmpty := len(line) == 0

			if *squeezeBlank && isEmpty && lastLineEmpty {
				continue
			}

			if *numberNonBlank {
				if !isEmpty {
					fmt.Fprintf(env.Stdout, "%6d\t%s\n", lineNumber, line)
					lineNumber++
				} else {
					fmt.Fprintln(env.Stdout)
				}
			} else if *number {
				fmt.Fprintf(env.Stdout, "%6d\t%s\n", lineNumber, line)
				lineNumber++
			} else {
				fmt.Fprintln(env.Stdout, line)
			}
			
			lastLineEmpty = isEmpty
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(env.Stderr, "cat: %s: %v\n", target, err)
			exitCode = 1
		}
	}

	return exitCode
}

