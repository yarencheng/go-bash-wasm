package bind

import (
	"context"
	"fmt"
	"sort"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Bind struct{}

func New() *Bind {
	return &Bind{}
}

func (b *Bind) Name() string {
	return "bind"
}

func (b *Bind) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("bind", pflag.ContinueOnError)
	list := flags.BoolP("list", "l", false, "list all readline functions")
	listVars := flags.BoolP("list-vars", "v", false, "list readline variable names and values")
	printStatus := flags.BoolP("print-status", "p", false, "print status of readline functions and bindings")
	listPublicVars := flags.BoolP("list-public-vars", "V", false, "list readline variable names and values in format for inputrc")
	printFuncs := flags.BoolP("print-funcs", "P", false, "print readline function names and bindings")
	listMacros := flags.BoolP("list-macros", "s", false, "list readline key sequences and macros")
	printMacros := flags.BoolP("print-macros", "S", false, "print readline key sequences and macros in format for inputrc")
	listKeyseq := flags.BoolP("list-keyseq", "X", false, "list key sequences bound to shell commands")
	_ = flags.StringP("file", "f", "", "read keybindings from FILE")
	_ = flags.StringP("query", "q", "", "query keys bound to FUNCTION")
	_ = flags.StringP("unbind", "u", "", "unbind all keys bound to FUNCTION")
	_ = flags.StringP("map", "m", "", "use KEYMAP")
	_ = flags.StringP("remove", "r", "", "remove binding for KEYSEQ")
	_ = flags.StringP("exec", "x", "", "execute SHELLCMD when KEYSEQ is entered")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "bind: %v\n", err)
		return 2
	}

	if *list {
		fmt.Fprintln(env.Stdout, "abort")
		fmt.Fprintln(env.Stdout, "accept-line")
		fmt.Fprintln(env.Stdout, "backward-char")
		fmt.Fprintln(env.Stdout, "forward-char")
		return 0
	}

	if *listVars || *listPublicVars {
		fmt.Fprintln(env.Stdout, "bell-style is audible")
		fmt.Fprintln(env.Stdout, "editing-mode is emacs")
		return 0
	}

	if *printStatus || *printFuncs || *listMacros || *printMacros || *listKeyseq {
		// Mock empty for now
		return 0
	}

	return 0
}
