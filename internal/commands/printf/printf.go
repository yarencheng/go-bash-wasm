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
	var varName string
	
	// Printf flags must come before the format string
	i := 0
	for ; i < len(args); i++ {
		if args[i] == "-v" && i+1 < len(args) {
			varName = args[i+1]
			i++
		} else if strings.HasPrefix(args[i], "-v") && len(args[i]) > 2 {
			varName = args[i][2:]
		} else if args[i] == "--" {
			i++
			break
		} else if strings.HasPrefix(args[i], "-") {
			// ignore unknown flags or handle them if needed
		} else {
			break
		}
	}

	if i >= len(args) {
		return 0
	}

	format := args[i]
	format = expandEscapes(format)
	remainingArgs := args[i+1:]

	var output strings.Builder
	writer := env.Stdout
	if varName != "" {
		writer = &output
	}

	if len(remainingArgs) == 0 {
		fmt.Fprint(writer, format)
	} else {
		// Bash printf reuses the format string until all arguments are exhausted
		for len(remainingArgs) > 0 {
			processedFormat := ""
			currentArgs := []interface{}{}
			argIdx := 0
			
			for j := 0; j < len(format); j++ {
				if format[j] == '%' && j+1 < len(format) {
					spec := format[j+1]
					if spec == '%' {
						processedFormat += "%%"
						j++
						continue
					}
					
					if argIdx >= len(remainingArgs) {
						// Out of args for this spec
						// Bash usually prints nothing or default for the type
						// We'll just stop processing format
						break
					}
					
					arg := remainingArgs[argIdx]
					argIdx++
					
					switch spec {
					case 'b':
						processedFormat += "%s"
						currentArgs = append(currentArgs, expandEscapes(arg))
						j++
					case 'q':
						processedFormat += "%s"
						currentArgs = append(currentArgs, shellQuote(arg))
						j++
					default:
						// Standard conversion
						processedFormat += "%" + string(spec)
						currentArgs = append(currentArgs, arg)
						j++
					}
				} else {
					processedFormat += string(format[j])
				}
			}
			
			fmt.Fprintf(writer, processedFormat, currentArgs...)
			
			if argIdx == 0 { // avoid infinite loop if no specifiers
				break
			}
			remainingArgs = remainingArgs[argIdx:]
		}
	}

	if varName != "" {
		if env.EnvVars == nil {
			env.EnvVars = make(map[string]string)
		}
		env.EnvVars[varName] = output.String()
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
