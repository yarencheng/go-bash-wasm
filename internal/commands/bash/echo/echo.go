package echo

import (
	"context"
	"fmt"
	"strconv"
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
				i++
			case 'b':
				b.WriteByte('\b')
				i++
			case 'e', 'E':
				b.WriteByte('\x1b')
				i++
			case 'f':
				b.WriteByte('\f')
				i++
			case 'n':
				b.WriteByte('\n')
				i++
			case 'r':
				b.WriteByte('\r')
				i++
			case 't':
				b.WriteByte('\t')
				i++
			case 'v':
				b.WriteByte('\v')
				i++
			case 'c':
				return b.String()
			case 'x':
				// hex: \xHH
				if i+2 < len(s) {
					hex := ""
					for j := 0; j < 2 && i+2+j < len(s); j++ {
						c := s[i+2+j]
						if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
							hex += string(c)
						} else {
							break
						}
					}
					if hex != "" {
						val, _ := strconv.ParseUint(hex, 16, 8)
						b.WriteByte(byte(val))
						i += 1 + len(hex)
					} else {
						b.WriteByte('\\')
						b.WriteByte('x')
						i++
					}
				} else {
					b.WriteByte('\\')
					b.WriteByte('x')
					i++
				}
			case '0':
				// octal: \0NNN (up to 3 digits)
				if i+2 < len(s) {
					octal := ""
					for j := 0; j < 3 && i+2+j < len(s); j++ {
						c := s[i+2+j]
						if c >= '0' && c <= '7' {
							octal += string(c)
						} else {
							break
						}
					}
					if octal != "" {
						val, _ := strconv.ParseUint(octal, 8, 8)
						b.WriteByte(byte(val))
						i += 1 + len(octal)
					} else {
						b.WriteByte(0)
						i++
					}
				} else {
					b.WriteByte(0)
					i++
				}
			case '\\':
				b.WriteByte('\\')
				i++
			default:
				b.WriteByte(s[i])
				// Don't increment i here if we don't recognize the sequence
			}
		} else {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}
