package declare

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Declare struct {
	name string
}

func New() *Declare {
	return &Declare{name: "declare"}
}

func NewWithName(name string) *Declare {
	return &Declare{name: name}
}

func (d *Declare) Name() string {
	return d.name
}

func (d *Declare) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet(d.name, pflag.ContinueOnError)
	printFlag := flags.BoolP("print", "p", false, "display the attributes and value of each NAME")
	exportFlag := flags.BoolP("export", "x", false, "make NAMEs export")
	readonlyFlag := flags.BoolP("readonly", "r", false, "make NAMEs readonly")
	integerFlag := flags.BoolP("integer", "i", false, "make NAMEs have the integer attribute")
	lower := flags.BoolP("lowercase", "l", false, "convert NAMEs to lowercase on assignment")
	upper := flags.BoolP("uppercase", "u", false, "convert NAMEs to uppercase on assignment")
	nameref := flags.BoolP("nameref", "n", false, "make NAME a reference to the variable named by its value")
	trace := flags.BoolP("trace", "t", false, "make NAMEs have the trace attribute")
	function := flags.BoolP("function", "f", false, "restrict action or display to function names and definitions")
	funcname := flags.BoolP("funcname", "F", false, "restrict display to function names only")
	global := flags.BoolP("global", "g", false, "create NAMEs in the global scope")
	array := flags.BoolP("array", "a", false, "make NAMEs indexed arrays")
	assoc := flags.BoolP("assoc", "A", false, "make NAMEs associative arrays")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "%s: %v\n", d.name, err)
		return 2
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: %s [-aAfFgilnrtux] [-p] [name[=value] ...]\n", d.name)
		fmt.Fprintf(env.Stdout, "Declare variables and give them attributes.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, d.name)
		return 0
	}

	if env.VarAttributes == nil {
		env.VarAttributes = make(map[string]uint32)
	}

	targets := flags.Args()

	// Handle listing attributes
	if len(targets) == 0 || (*printFlag && len(targets) == 0) {
		keys := make([]string, 0, len(env.EnvVars))
		for k := range env.EnvVars {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			attrStr := d.formatAttr(env.VarAttributes[k])
			if *printFlag {
				fmt.Fprintf(env.Stdout, "declare %s %s=\"%s\"\n", attrStr, k, env.EnvVars[k])
			} else {
				fmt.Fprintf(env.Stdout, "%s=\"%s\"\n", k, env.EnvVars[k])
			}
		}
		return 0
	}

	var attrMask uint32
	if *exportFlag {
		attrMask |= commands.AttrExport
	}
	if *readonlyFlag {
		attrMask |= commands.AttrReadonly
	}
	if *integerFlag {
		attrMask |= commands.AttrInteger
	}
	if *array {
		attrMask |= commands.AttrArray
	}
	if *assoc {
		attrMask |= commands.AttrAssoc
	}
	if *nameref {
		attrMask |= commands.AttrNameref
	}
	if *lower {
		attrMask |= commands.AttrLowercase
	}
	if *upper {
		attrMask |= commands.AttrUppercase
	}
	if *trace {
		// trace attribute not fully simulated
	}
	_ = function
	_ = funcname
	_ = global

	for _, arg := range targets {
		name := arg
		value := ""
		hasValue := false

		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			name = parts[0]
			value = parts[1]
			hasValue = true
		}

		// Check readonly
		if (env.VarAttributes[name] & commands.AttrReadonly) != 0 {
			fmt.Fprintf(env.Stderr, "%s: %s: readonly variable\n", d.name, name)
			return 1
		}

		// Set attributes
		env.VarAttributes[name] |= attrMask

		if hasValue {
			if (attrMask & commands.AttrLowercase) != 0 {
				value = strings.ToLower(value)
			}
			if (attrMask & commands.AttrUppercase) != 0 {
				value = strings.ToUpper(value)
			}
			env.EnvVars[name] = value
		} else if attrMask != 0 && (attrMask&commands.AttrArray) != 0 {
			if env.Arrays == nil {
				env.Arrays = make(map[string][]string)
			}
			if _, ok := env.Arrays[name]; !ok {
				env.Arrays[name] = []string{}
			}
		} else if attrMask != 0 && (attrMask&commands.AttrAssoc) != 0 {
			if env.AssocArrays == nil {
				env.AssocArrays = make(map[string]map[string]string)
			}
			if _, ok := env.AssocArrays[name]; !ok {
				env.AssocArrays[name] = make(map[string]string)
			}
		}
	}

	return 0
}

func (d *Declare) formatAttr(attr uint32) string {
	res := ""
	if (attr & commands.AttrArray) != 0 {
		res += "-a"
	} else if (attr & commands.AttrAssoc) != 0 {
		res += "-A"
	}
	if (attr & commands.AttrInteger) != 0 {
		res += "-i"
	}
	if (attr & commands.AttrReadonly) != 0 {
		res += "-r"
	}
	if (attr & commands.AttrExport) != 0 {
		res += "-x"
	}
	return res
}
