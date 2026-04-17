package expr

import (
	"context"
	"fmt"
	"strconv"

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

	if len(args) == 3 {
		op1, err1 := strconv.Atoi(args[0])
		operator := args[1]
		op2, err2 := strconv.Atoi(args[2])

		if err1 == nil && err2 == nil {
			switch operator {
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
			case "=":
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
		}
	}

	// Default: just print the first arg (very basic)
	fmt.Fprintln(env.Stdout, args[0])
	return 0
}
