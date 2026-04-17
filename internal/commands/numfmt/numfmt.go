package numfmt

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Numfmt struct{}

func New() *Numfmt {
	return &Numfmt{}
}

func (n *Numfmt) Name() string {
	return "numfmt"
}

func (n *Numfmt) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("numfmt", pflag.ContinueOnError)
	from := flags.String("from", "none", "auto-scale input numbers; UNIT can be: none, auto, si, iec, iec-i")
	to := flags.String("to", "none", "auto-scale output numbers; UNIT can be: none, auto, si, iec, iec-i")
	header := flags.Int("header", 0, "skip N header lines")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "numfmt: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return n.process(env, env.Stdin, *from, *to, *header)
	}

	exitCode := 0
	for _, arg := range remaining {
		file, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "numfmt: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := n.process(env, file, *from, *to, *header); status != 0 {
			exitCode = status
		}
		file.Close()
	}

	return exitCode
}

func (n *Numfmt) process(env *commands.Environment, r io.Reader, from, to string, headers int) int {
	scanner := bufio.NewScanner(r)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		if lineCount <= headers {
			fmt.Fprintln(env.Stdout, line)
			continue
		}

		parts := strings.Fields(line)
		for i, part := range parts {
			val, err := n.parse(part, from)
			if err != nil {
				fmt.Fprint(env.Stdout, part)
			} else {
				fmt.Fprint(env.Stdout, n.format(val, to))
			}
			if i < len(parts)-1 {
				fmt.Fprint(env.Stdout, " ")
			}
		}
		fmt.Fprintln(env.Stdout)
	}
	return 0
}

func (n *Numfmt) parse(s string, unit string) (float64, error) {
	if unit == "none" {
		return strconv.ParseFloat(s, 64)
	}

	// Basic scaling parser
	suffix := ""
	valStr := s
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] >= '0' && s[i] <= '9') || s[i] == '.' {
			valStr = s[:i+1]
			suffix = s[i+1:]
			break
		}
	}

	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return 0, err
	}

	if suffix == "" {
		return val, nil
	}

	multiplier := 1.0
	base := 1000.0
	if unit == "iec" || unit == "iec-i" || unit == "auto" {
		base = 1024.0
	}

	switch strings.ToUpper(suffix) {
	case "K", "KB":
		multiplier = base
	case "M", "MB":
		multiplier = base * base
	case "G", "GB":
		multiplier = base * base * base
	case "T", "TB":
		multiplier = base * base * base * base
	case "P", "PB":
		multiplier = base * base * base * base * base
	case "E", "EB":
		multiplier = base * base * base * base * base * base
	case "Z", "ZB":
		multiplier = base * base * base * base * base * base * base
	case "Y", "YB":
		multiplier = base * base * base * base * base * base * base * base
	}

	return val * multiplier, nil
}

func (n *Numfmt) format(val float64, unit string) string {
	if unit == "none" {
		return strconv.FormatFloat(val, 'f', -1, 64)
	}

	base := 1000.0
	suffixes := []string{"", "k", "M", "G", "T", "P", "E", "Z", "Y"}
	if unit == "iec" || unit == "iec-i" || unit == "auto" {
		base = 1024.0
		suffixes = []string{"", "K", "M", "G", "T", "P", "E", "Z", "Y"}
	}

	if val < base {
		return strconv.FormatFloat(val, 'f', -1, 64)
	}

	exp := int(math.Log(val) / math.Log(base))
	if exp >= len(suffixes) {
		exp = len(suffixes) - 1
	}

	scaled := val / math.Pow(base, float64(exp))
	return fmt.Sprintf("%.1f%s", scaled, suffixes[exp])
}
