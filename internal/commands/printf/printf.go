package printf

import (
	"context"
	"fmt"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Printf struct{}

func New() *Printf {
	return &Printf{}
}

func (p *Printf) Name() string {
	return "printf"
}

func (p *Printf) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		return 0
	}

	format := args[0]
	// Basic escape sequence replacement
	format = strings.ReplaceAll(format, "\\n", "\n")
	format = strings.ReplaceAll(format, "\\t", "\t")

	if len(args) == 1 {
		fmt.Fprint(env.Stdout, format)
		return 0
	}

	// Convert args to interface slices
	interfaceArgs := make([]interface{}, len(args)-1)
	for i, arg := range args[1:] {
		interfaceArgs[i] = arg
	}

	fmt.Fprintf(env.Stdout, format, interfaceArgs...)
	return 0
}
