package expr

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Expr struct{}

func New() *Expr {
	return &Expr{}
}

func (e *Expr) Name() string {
	return "expr"
}

func (e *Expr) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		fmt.Fprintf(env.Stderr, "expr: missing operand\n")
		return 2
	}

	// Handle string operators first
	if args[0] == "length" && len(args) == 2 {
		fmt.Fprintln(env.Stdout, len(args[1]))
		return 0
	}
	if args[0] == "index" && len(args) == 3 {
		str := args[1]
		chars := args[2]
		idx := strings.IndexAny(str, chars)
		if idx == -1 {
			fmt.Fprintln(env.Stdout, 0)
		} else {
			fmt.Fprintln(env.Stdout, idx+1)
		}
		return 0
	}
	if args[0] == "substr" && len(args) == 4 {
		str := args[1]
		start, _ := strconv.Atoi(args[2])
		length, _ := strconv.Atoi(args[3])
		if start < 1 || start > len(str) {
			fmt.Fprintln(env.Stdout, "")
		} else {
			end := start - 1 + length
			if end > len(str) {
				end = len(str)
			}
			fmt.Fprintln(env.Stdout, str[start-1:end])
		}
		return 0
	}

	if len(args) == 3 {
		s1 := args[0]
		op := args[1]
		s2 := args[2]

		// Logical operators
		if op == "|" {
			if s1 != "" && s1 != "0" {
				fmt.Fprintln(env.Stdout, s1)
			} else {
				fmt.Fprintln(env.Stdout, s2)
			}
			return 0
		}
		if op == "&" {
			if s1 != "" && s1 != "0" && s2 != "" && s2 != "0" {
				fmt.Fprintln(env.Stdout, s1)
			} else {
				fmt.Fprintln(env.Stdout, 0)
			}
			return 0
		}

		op1, err1 := strconv.Atoi(s1)
		op2, err2 := strconv.Atoi(s2)

		if err1 == nil && err2 == nil {
			switch op {
			case "+":
				fmt.Fprintln(env.Stdout, op1+op2)
				return 0
			case "-":
				fmt.Fprintln(env.Stdout, op1-op2)
				return 0
			case "*":
				fmt.Fprintln(env.Stdout, op1*op2)
				return 0
			case "/":
				if op2 == 0 {
					fmt.Fprintf(env.Stderr, "expr: division by zero\n")
					return 2
				}
				fmt.Fprintln(env.Stdout, op1/op2)
				return 0
			case "%":
				if op2 == 0 {
					fmt.Fprintf(env.Stderr, "expr: division by zero\n")
					return 2
				}
				fmt.Fprintln(env.Stdout, op1%op2)
				return 0
			case ">":
				if op1 > op2 {
					fmt.Fprintln(env.Stdout, 1)
				} else {
					fmt.Fprintln(env.Stdout, 0)
				}
				return 0
			case "<":
				if op1 < op2 {
					fmt.Fprintln(env.Stdout, 1)
				} else {
					fmt.Fprintln(env.Stdout, 0)
				}
				return 0
			case ">=":
				if op1 >= op2 {
					fmt.Fprintln(env.Stdout, 1)
				} else {
					fmt.Fprintln(env.Stdout, 0)
				}
				return 0
			case "<=":
				if op1 <= op2 {
					fmt.Fprintln(env.Stdout, 1)
				} else {
					fmt.Fprintln(env.Stdout, 0)
				}
				return 0
			case "=", "==":
				if op1 == op2 {
					fmt.Fprintln(env.Stdout, 1)
				} else {
					fmt.Fprintln(env.Stdout, 0)
				}
				return 0
			case "!=":
				if op1 != op2 {
					fmt.Fprintln(env.Stdout, 1)
				} else {
					fmt.Fprintln(env.Stdout, 0)
				}
				return 0
			}
		} else {
			// String comparison
			switch op {
			case "=", "==":
				if s1 == s2 {
					fmt.Fprintln(env.Stdout, 1)
				} else {
					fmt.Fprintln(env.Stdout, 0)
				}
				return 0
			case "!=":
				if s1 != s2 {
					fmt.Fprintln(env.Stdout, 1)
				} else {
					fmt.Fprintln(env.Stdout, 0)
				}
				return 0
			}
		}
	}

	// Default: just print the first arg (very basic)
	fmt.Fprintln(env.Stdout, args[0])
	return 0
}
