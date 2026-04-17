package factor

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Factor struct{}

func New() *Factor {
	return &Factor{}
}

func (f *Factor) Name() string {
	return "factor"
}

func (f *Factor) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		return 0
	}

	for _, arg := range args {
		n, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(env.Stderr, "factor: '%s' is not a valid positive integer\n", arg)
			continue
		}

		fmt.Fprintf(env.Stdout, "%d:", n)
		if n > 1 {
			factors := f.getFactors(n)
			for _, fact := range factors {
				fmt.Fprintf(env.Stdout, " %d", fact)
			}
		}
		fmt.Fprintln(env.Stdout)
	}

	return 0
}

func (f *Factor) getFactors(n uint64) []uint64 {
	var factors []uint64
	d := uint64(2)
	for n%d == 0 {
		factors = append(factors, d)
		n /= d
	}
	d = 3
	for d*d <= n {
		for n%d == 0 {
			factors = append(factors, d)
			n /= d
		}
		d += 2
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
