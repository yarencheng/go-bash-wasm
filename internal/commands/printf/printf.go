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
	format = expandEscapes(format)
	args = args[1:]

	if len(args) == 0 {
		fmt.Fprint(env.Stdout, format)
		return 0
	}

	// Bash printf reuses the format string until all arguments are exhausted
	for len(args) > 0 {
		// We need to parse the format string to handle %b and %q manually
		// because Go's fmt doesn't know about them.
		
		processedFormat := ""
		currentArgs := []interface{}{}
		argIdx := 0
		
		for i := 0; i < len(format); i++ {
			if format[i] == '%' && i+1 < len(format) {
				spec := format[i+1]
				if spec == '%' {
					processedFormat += "%%"
					i++
					continue
				}
				
				if argIdx >= len(args) {
					// Out of args for this spec, standard printf behavior defaults vary
					// but usually we just append the rest.
					break
				}
				
				arg := args[argIdx]
				argIdx++
				
				switch spec {
				case 'b':
					processedFormat += "%s"
					currentArgs = append(currentArgs, expandEscapes(arg))
					i++
				case 'q':
					processedFormat += "%s"
					currentArgs = append(currentArgs, shellQuote(arg))
					i++
				default:
					// Standard conversion
					// We'll just pass it to fmt for now, but need to find the full spec (e.g. %.2f)
					// For simplicity we'll just take the next char if it's a standard one.
					processedFormat += "%" + string(spec)
					currentArgs = append(currentArgs, arg)
					i++
				}
			} else {
				processedFormat += string(format[i])
			}
		}
		
		fmt.Fprintf(env.Stdout, processedFormat, currentArgs...)
		
		if argIdx == 0 { // avoid infinite loop if no specifiers
			break
		}
		args = args[argIdx:]
	}

	return 0
}

func expandEscapes(s string) string {
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, "\\t", "\t")
	s = strings.ReplaceAll(s, "\\\\", "\\")
	s = strings.ReplaceAll(s, "\\r", "\r")
	return s
}

func shellQuote(s string) string {
	if s == "" {
		return "''"
	}
	// Simple quoting: if contains special chars, wrap in single quotes and escape existing single quotes
	if strings.ContainsAny(s, " \t\n\r\"'$`\\!&*()[]{}<>?;|") {
		return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
	}
	return s
}
