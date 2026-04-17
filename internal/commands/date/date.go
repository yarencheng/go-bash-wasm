package date

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Date struct{}

func New() *Date {
	return &Date{}
}

func (d *Date) Name() string {
	return "date"
}

func (d *Date) Run(ctx context.Context, env *commands.Environment, args []string) int {
	now := time.Now()
	format := time.UnixDate

	if len(args) > 0 {
		arg := args[0]
		if strings.HasPrefix(arg, "+") {
			// Basic support for format string
			// We map some common ones to Go format
			f := arg[1:]
			f = strings.ReplaceAll(f, "%Y", "2006")
			f = strings.ReplaceAll(f, "%m", "01")
			f = strings.ReplaceAll(f, "%d", "02")
			f = strings.ReplaceAll(f, "%H", "15")
			f = strings.ReplaceAll(f, "%M", "04")
			f = strings.ReplaceAll(f, "%S", "05")
			fmt.Fprintln(env.Stdout, now.Format(f))
			return 0
		}
	}

	fmt.Fprintln(env.Stdout, now.Format(format))
	return 0
}
