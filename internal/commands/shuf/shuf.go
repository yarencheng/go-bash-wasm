package shuf

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Shuf struct{}

func New() *Shuf {
	return &Shuf{}
}

func (s *Shuf) Name() string {
	return "shuf"
}

func (s *Shuf) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("shuf", pflag.ContinueOnError)
	echo := flags.BoolP("echo", "e", false, "treat each ARG as an input line")
	inputRange := flags.StringP("input-range", "i", "", "treat each number LO through HI as an input line")
	count := flags.IntP("head-count", "n", -1, "output at most COUNT lines")
	output := flags.StringP("output", "o", "", "write result to FILE instead of standard output")
	repeat := flags.BoolP("repeat", "r", false, "output lines can be repeated")
	zero := flags.BoolP("zero", "z", false, "end each output line with NUL, not newline")
	_ = flags.String("random-source", "", "get random bytes from FILE")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "shuf: %v\n", err)
		}
		return 1
	}

	remaining := flags.Args()
	var lines []string

	if *echo {
		lines = remaining
	} else if *inputRange != "" {
		parts := strings.Split(*inputRange, "-")
		if len(parts) != 2 {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "shuf: invalid input range: %s\n", *inputRange)
			}
			return 1
		}
		lo, err1 := strconv.Atoi(parts[0])
		hi, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil || lo > hi {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "shuf: invalid input range: %s\n", *inputRange)
			}
			return 1
		}
		for i := lo; i <= hi; i++ {
			lines = append(lines, strconv.Itoa(i))
		}
	} else {
		var r io.Reader
		if len(remaining) == 0 || remaining[0] == "-" {
			r = env.Stdin
		} else {
			f, err := env.FS.Open(remaining[0])
			if err != nil {
				if env.Stderr != nil {
					fmt.Fprintf(env.Stderr, "shuf: %s: %v\n", remaining[0], err)
				}
				return 1
			}
			defer f.Close()
			r = f
		}
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	if len(lines) == 0 {
		return 0
	}

	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	var out io.Writer = env.Stdout
	if *output != "" {
		f, err := env.FS.Create(*output)
		if err != nil {
			if env.Stderr != nil {
				fmt.Fprintf(env.Stderr, "shuf: %s: %v\n", *output, err)
			}
			return 1
		}
		defer f.Close()
		out = f
	}

	term := "\n"
	if *zero {
		term = "\x00"
	}

	if *repeat {
		limit := *count
		for i := 0; limit < 0 || i < limit; i++ {
			fmt.Fprintf(out, "%s%s", lines[randSource.Intn(len(lines))], term)
		}
	} else {
		randSource.Shuffle(len(lines), func(i, j int) {
			lines[i], lines[j] = lines[j], lines[i]
		})
		
		limit := len(lines)
		if *count >= 0 && *count < limit {
			limit = *count
		}
		for i := 0; i < limit; i++ {
			fmt.Fprintf(out, "%s%s", lines[i], term)
		}
	}

	return 0
}
