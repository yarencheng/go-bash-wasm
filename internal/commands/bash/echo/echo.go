package echo

import (
	"context"
	"fmt"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Echo struct{}

func New() *Echo {
	return &Echo{}
}

func (e *Echo) Name() string {
	return "echo"
}

func (e *Echo) Run(ctx context.Context, env *commands.Environment, args []string) int {
	noNewline := false
	interpretEscapes := false

	// Handle flags manually since echo is special about them (only if they are at the beginning)
	i := 0
	for i < len(args) {
		arg := args[i]
		if len(arg) < 2 || arg[0] != '-' {
			break
		}

		isFlag := true
		for j := 1; j < len(arg); j++ {
			switch arg[j] {
			case 'n':
				noNewline = true
			case 'e':
				interpretEscapes = true
			case 'E':
				interpretEscapes = false
			default:
				isFlag = false
				break
			}
			if !isFlag {
				break
			}
		}

		if !isFlag {
			break
		}
		i++
	}

	text := strings.Join(args[i:], " ")
	if interpretEscapes {
		text = unescape(text)
	}

	fmt.Fprint(env.Stdout, text)
	if !noNewline {
		fmt.Fprintln(env.Stdout)
	}

	return 0
}

func unescape(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case 'a':
				b.WriteByte('\a')
			case 'b':
				b.WriteByte('\b')
			case 'e', 'E':
				b.WriteByte('\x1b')
			case 'f':
				b.WriteByte('\f')
			case 'n':
				b.WriteByte('\n')
			case 'r':
				b.WriteByte('\r')
			case 't':
				b.WriteByte('\t')
			case 'v':
				b.WriteByte('\v')
			case '\\':
				b.WriteByte('\\')
			default:
				b.WriteByte(s[i])
				b.WriteByte(s[i+1])
			}
			i++
		} else {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}
