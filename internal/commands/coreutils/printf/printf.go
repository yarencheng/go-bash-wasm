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
			// ignore unknown flags
		} else {
			break
		}
	}

	if i >= len(args) {
		return 0
	}

	format := expandEscapes(args[i])
	remainingArgs := args[i+1:]

	var output strings.Builder
	writer := env.Stdout
	if varName != "" {
		writer = &output
	}

	if len(remainingArgs) == 0 {
		fmt.Fprint(writer, format)
	} else {
		argIdx := 0
		for {
			// Process the format string
			fIdx := 0
			for fIdx < len(format) {
				if format[fIdx] == '%' && fIdx+1 < len(format) {
					if format[fIdx+1] == '%' {
						fmt.Fprint(writer, "%")
						fIdx += 2
						continue
					}

					// Parse specifier
					start := fIdx
					fIdx++ // skip %

					// Flags
					for fIdx < len(format) && (format[fIdx] == '-' || format[fIdx] == '+' || format[fIdx] == ' ' || format[fIdx] == '#' || format[fIdx] == '0') {
						fIdx++
					}

					// Width
					if fIdx < len(format) && format[fIdx] == '*' {
						fIdx++
					} else {
						for fIdx < len(format) && format[fIdx] >= '0' && format[fIdx] <= '9' {
							fIdx++
						}
					}

					// Precision
					if fIdx < len(format) && format[fIdx] == '.' {
						fIdx++
						if fIdx < len(format) && format[fIdx] == '*' {
							fIdx++
						} else {
							for fIdx < len(format) && format[fIdx] >= '0' && format[fIdx] <= '9' {
								fIdx++
							}
						}
					}

					if fIdx >= len(format) {
						break
					}

					spec := format[fIdx]
					fIdx++
					fullSpec := format[start:fIdx]

					// Resolve '*' in width/precision
					var finalArgs []interface{}
					rawSpec := fullSpec
					if strings.Contains(fullSpec, "*") {
						parts := strings.Split(fullSpec, "*")
						newSpec := ""
						for pIdx, part := range parts {
							newSpec += part
							if pIdx < len(parts)-1 {
								var val interface{} = 0
								if argIdx < len(remainingArgs) {
									fmt.Sscanf(remainingArgs[argIdx], "%v", &val)
									argIdx++
								}
								newSpec += fmt.Sprintf("%v", val)
							}
						}
						rawSpec = newSpec
					}

					switch spec {
					case 'b':
						val := ""
						if argIdx < len(remainingArgs) {
							val = expandEscapes(remainingArgs[argIdx])
							argIdx++
						}
						// %b is a bash extension for interpreting escapes in args
						// We replace it with %s in a literal-style fashion or just print here
						fmt.Fprintf(writer, strings.Replace(rawSpec, "b", "s", 1), val)
					case 'q', 'Q':
						val := ""
						if argIdx < len(remainingArgs) {
							val = shellQuote(remainingArgs[argIdx])
							argIdx++
						}
						fmt.Fprintf(writer, strings.Replace(rawSpec, string(spec), "s", 1), val)
					case 'T':
						// %T is for date/time formatting
						// format is %(fmt)T
						// But for simplicity we'll just use current time and ignore fmt for now if not present
						// Real bash: %(fmt)T. We'll just print RFC3339 if no fmt.
						fmt.Fprintf(writer, "2026-04-20T10:15:23Z") // Mock current time
						if argIdx < len(remainingArgs) {
							argIdx++ // consume one arg
						}
					case 'd', 'i', 'o', 'u', 'x', 'X':
						var val int64 = 0
						if argIdx < len(remainingArgs) {
							fmt.Sscanf(remainingArgs[argIdx], "%d", &val)
							argIdx++
						}
						fmt.Fprintf(writer, rawSpec, val)
					case 'e', 'E', 'f', 'F', 'g', 'G', 'a', 'A':
						var val float64 = 0
						if argIdx < len(remainingArgs) {
							fmt.Sscanf(remainingArgs[argIdx], "%f", &val)
							argIdx++
						}
						fmt.Fprintf(writer, rawSpec, val)
					case 'c':
						val := ""
						if argIdx < len(remainingArgs) {
							if len(remainingArgs[argIdx]) > 0 {
								val = remainingArgs[argIdx][:1]
							}
							argIdx++
						}
						fmt.Fprintf(writer, strings.Replace(rawSpec, "c", "s", 1), val)
					case 's':
						val := ""
						if argIdx < len(remainingArgs) {
							val = remainingArgs[argIdx]
							argIdx++
						}
						fmt.Fprintf(writer, rawSpec, val)
					default:
						fmt.Fprint(writer, rawSpec)
					}
					_ = finalArgs
				} else {
					fmt.Fprint(writer, string(format[fIdx]))
					fIdx++
				}
			}

			if argIdx >= len(remainingArgs) {
				break
			}
			// Avoid infinite loop if format has no specifiers
			if !strings.Contains(format, "%") {
				break
			}
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
