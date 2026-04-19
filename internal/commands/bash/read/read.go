package read

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Read struct{}

func New() *Read {
	return &Read{}
}

func (r *Read) Name() string {
	return "read"
}

func (r *Read) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("read", pflag.ContinueOnError)
	prompt := flags.StringP("prompt", "p", "", "display PROMPT without a trailing newline, before attempting to read")
	raw := flags.BoolP("raw", "r", false, "do not allow backslashes to escape any characters")
	delim := flags.StringP("delimiter", "d", "\n", "terminate after reading DELIM, rather than newline")
	nchars := flags.IntP("nchars", "n", -1, "return after reading NCHARS characters")
	ncharsExact := flags.IntP("nchars-exact", "N", -1, "return only after reading exactly NCHARS characters")
	array := flags.StringP("array", "a", "", "assign the words read to sequential indices of the array variable ARRAY")
	_ = flags.BoolP("silent", "s", false, "do not echo input coming from a terminal")
	timeout := flags.Float64P("timeout", "t", 0, "terminate after TIMEOUT seconds")
	_ = flags.IntP("fd", "u", 0, "read from file descriptor FD instead of the standard input")
	_ = flags.BoolP("readline", "e", false, "use Readline to obtain the line")
	_ = flags.StringP("initial", "i", "", "use TEXT as initial text for Readline")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "read: %v\n", err)
		}
		return 1
	}

	if *timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(*timeout*float64(time.Second)))
		defer cancel()
	}

	if *prompt != "" {
		fmt.Fprint(env.Stdout, *prompt)
	}

	reader := bufio.NewReader(env.Stdin)
	var line strings.Builder
	var lastChar rune
	count := 0
	limit := *nchars
	if *ncharsExact != -1 {
		limit = *ncharsExact
	}

	d := '\n'
	if len(*delim) > 0 {
		d = rune((*delim)[0])
	}

	readErr := make(chan error, 1)
	readDone := make(chan struct{}, 1)

	go func() {
		for {
			if limit != -1 && count >= limit {
				break
			}

			char, _, err := reader.ReadRune()
			if err != nil {
				readErr <- err
				return
			}

			if !*raw && lastChar == '\\' {
				line.WriteRune(char)
				lastChar = 0
				count++
				continue
			}

			if char == d {
				break
			}

			if !*raw && char == '\\' {
				lastChar = '\\'
				continue
			}

			line.WriteRune(char)
			lastChar = char
			count++
		}
		readDone <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			return 142 // Common bash timeout exit code (128 + 14)
		}
		return 1
	case <-readErr:
		// EOF or other error
	case <-readDone:
	}

	content := line.String()
	fields := flags.Args()

	ifs := env.EnvVars["IFS"]
	if ifs == "" {
		ifs = " \t\n"
	}

	splitFn := func(r rune) bool {
		return strings.ContainsRune(ifs, r)
	}

	if *array != "" {
		words := strings.FieldsFunc(content, splitFn)
		if env.Arrays == nil {
			env.Arrays = make(map[string][]string)
		}
		env.Arrays[*array] = words
		return 0
	}

	if len(fields) == 0 {
		env.EnvVars["REPLY"] = content
		return 0
	}

	words := strings.FieldsFunc(content, splitFn)
	for i, field := range fields {
		if i < len(words) {
			if i == len(fields)-1 {
				// Last field gets the rest of the line including spaces/separators
				// We need to find the actual start of this word in the content
				// This is simplified:
				remaining := strings.Join(words[i:], " ")
				env.EnvVars[field] = remaining
			} else {
				env.EnvVars[field] = words[i]
			}
		} else {
			env.EnvVars[field] = ""
		}
	}

	return 0
}
