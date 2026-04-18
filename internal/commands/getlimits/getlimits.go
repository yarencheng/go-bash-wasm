package getlimits

import (
	"context"
	"fmt"
	"math"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Getlimits struct{}

func New() *Getlimits {
	return &Getlimits{}
}

func (g *Getlimits) Name() string {
	return "getlimits"
}

func (g *Getlimits) Run(ctx context.Context, env *commands.Environment, args []string) int {
	// Simple implementation of getlimits following coreutils output format

	// Integer limits
	fmt.Fprintf(env.Stdout, "CHAR_MAX=%d\n", 127)
	fmt.Fprintf(env.Stdout, "CHAR_OFLOW=%d\n", 128)
	fmt.Fprintf(env.Stdout, "CHAR_MIN=%d\n", -128)
	fmt.Fprintf(env.Stdout, "CHAR_UFLOW=%d\n", -129)

	fmt.Fprintf(env.Stdout, "SCHAR_MAX=%d\n", 127)
	fmt.Fprintf(env.Stdout, "SCHAR_OFLOW=%d\n", 128)
	fmt.Fprintf(env.Stdout, "SCHAR_MIN=%d\n", -128)
	fmt.Fprintf(env.Stdout, "SCHAR_UFLOW=%d\n", -129)

	fmt.Fprintf(env.Stdout, "UCHAR_MAX=%d\n", 255)
	fmt.Fprintf(env.Stdout, "UCHAR_OFLOW=%d\n", 256)

	fmt.Fprintf(env.Stdout, "SHRT_MAX=%d\n", math.MaxInt16)
	fmt.Fprintf(env.Stdout, "SHRT_OFLOW=%d\n", math.MaxInt16+1)
	fmt.Fprintf(env.Stdout, "SHRT_MIN=%d\n", math.MinInt16)
	fmt.Fprintf(env.Stdout, "SHRT_UFLOW=%d\n", math.MinInt16-1)

	fmt.Fprintf(env.Stdout, "INT_MAX=%d\n", math.MaxInt32)
	fmt.Fprintf(env.Stdout, "INT_OFLOW=%d\n", int64(math.MaxInt32)+1)
	fmt.Fprintf(env.Stdout, "INT_MIN=%d\n", math.MinInt32)
	fmt.Fprintf(env.Stdout, "INT_UFLOW=%d\n", int64(math.MinInt32)-1)

	fmt.Fprintf(env.Stdout, "UINT_MAX=%d\n", math.MaxUint32)
	fmt.Fprintf(env.Stdout, "UINT_OFLOW=%d\n", uint64(math.MaxUint32)+1)

	// Time, off_t etc are usually 64-bit in modern systems even on 32-bit arch
	fmt.Fprintf(env.Stdout, "TIME_MAX=%d\n", math.MaxInt64)
	fmt.Fprintf(env.Stdout, "OFF_T_MAX=%d\n", math.MaxInt64)

	fmt.Fprintf(env.Stdout, "SSIZE_MAX=%d\n", math.MaxInt32) // Assuming 32-bit size_t for wasm32
	fmt.Fprintf(env.Stdout, "SIZE_MAX=%d\n", uint32(math.MaxUint32))

	return 0
}
