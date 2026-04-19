package letcmd

import (
	"context"
	"strconv"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Let struct{}

func New() *Let {
	return &Let{}
}

func (l *Let) Name() string {
	return "let"
}

func (l *Let) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		return 1
	}

	lastStatus := 0
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) == 2 {
			varName := parts[0]
			expr := parts[1]
			// Very basic evaluation for now
			val := l.eval(env, expr)
			env.EnvVars[varName] = strconv.FormatInt(val, 10)
			if val == 0 {
				lastStatus = 1
			} else {
				lastStatus = 0
			}
		} else {
			// Arithmetic without assignment
			val := l.eval(env, arg)
			if val == 0 {
				lastStatus = 1
			} else {
				lastStatus = 0
			}
		}
	}
	return lastStatus
}

func (l *Let) eval(env *commands.Environment, expr string) int64 {
	// For now, only handles simple numbers or variable names
	if val, err := strconv.ParseInt(expr, 10, 64); err == nil {
		return val
	}
	if valStr, ok := env.EnvVars[expr]; ok {
		if val, err := strconv.ParseInt(valStr, 10, 64); err == nil {
			return val
		}
	}
	return 0
}
