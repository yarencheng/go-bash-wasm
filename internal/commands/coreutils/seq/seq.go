package seq

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Seq struct{}

func New() *Seq {
	return &Seq{}
}

func (s *Seq) Name() string {
	return "seq"
}

func (s *Seq) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("seq", pflag.ContinueOnError)
	separator := flags.StringP("separator", "s", "\n", "use STRING to separate numbers (default: \\n)")
	equalWidth := flags.BoolP("equal-width", "w", false, "equalize width by padding with leading zeroes")
	formatFlag := flags.StringP("format", "f", "", "use printf style floating-point FORMAT")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "seq: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "seq: missing operand\n")
		return 1
	}

	start := 1.0
	increment := 1.0
	end := 0.0

	var err error
	if len(targets) == 1 {
		end, err = strconv.ParseFloat(targets[0], 64)
	} else if len(targets) == 2 {
		start, err = strconv.ParseFloat(targets[0], 64)
		end, _ = strconv.ParseFloat(targets[1], 64)
	} else if len(targets) == 3 {
		start, err = strconv.ParseFloat(targets[0], 64)
		increment, _ = strconv.ParseFloat(targets[1], 64)
		end, _ = strconv.ParseFloat(targets[2], 64)
	}

	if err != nil {
		fmt.Fprintf(env.Stderr, "seq: invalid operand\n")
		return 1
	}

	if increment == 0 {
		fmt.Fprintf(env.Stderr, "seq: increment cannot be 0\n")
		return 1
	}

	format := "%g"
	if *formatFlag != "" {
		format = *formatFlag
	} else if *equalWidth {
		// Calculate max width
		maxW := 0
		for val := start; (increment > 0 && val <= end) || (increment < 0 && val >= end); val += increment {
			w := len(fmt.Sprintf("%g", val))
			if w > maxW {
				maxW = w
			}
		}
		format = fmt.Sprintf("%%0%dg", maxW)
	}

	first := true
	for val := start; (increment > 0 && val <= end) || (increment < 0 && val >= end); val += increment {
		if !first {
			fmt.Fprint(env.Stdout, *separator)
		}
		fmt.Fprintf(env.Stdout, format, val)
		first = false
	}
	if !first {
		fmt.Fprint(env.Stdout, "\n")
	}

	return 0
}
