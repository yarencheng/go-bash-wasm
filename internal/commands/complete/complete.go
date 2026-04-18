package complete

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Complete struct{}

func New() *Complete {
	return &Complete{}
}

func (c *Complete) Name() string {
	return "complete"
}

func (c *Complete) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("complete", pflag.ContinueOnError)
	pflag_p := flags.BoolP("print", "p", false, "print existing completion specifications")
	pflag_r := flags.BoolP("remove", "r", false, "remove a completion specification")

	// Complex flags that take arguments
	flag_A := flags.StringP("action", "A", "", "action")
	flag_G := flags.StringP("globpat", "G", "", "globpat")
	flag_W := flags.StringP("wordlist", "W", "", "wordlist")
	flag_P := flags.StringP("prefix", "P", "", "prefix")
	flag_S := flags.StringP("suffix", "S", "", "suffix")
	flag_X := flags.StringP("filterpat", "X", "", "filterpat")
	flag_F := flags.StringP("function", "F", "", "function")
	flag_C := flags.StringP("command", "C", "", "command")

	// Boolean flags for common actions
	flag_a := flags.BoolP("alias", "a", false, "alias")
	flag_b := flags.BoolP("builtin", "b", false, "builtin")
	flag_c := flags.BoolP("cmd", "c", false, "command")
	flag_d := flags.BoolP("directory", "d", false, "directory")
	flag_e := flags.BoolP("export", "e", false, "export")
	flag_f := flags.BoolP("file", "f", false, "file")
	flag_g := flags.BoolP("group", "g", false, "group")
	flag_j := flags.BoolP("job", "j", false, "job")
	flag_k := flags.BoolP("keyword", "k", false, "keyword")
	flag_s := flags.BoolP("service", "s", false, "service")
	flag_u := flags.BoolP("user", "u", false, "user")
	flag_v := flags.BoolP("variable", "v", false, "variable")

	_ = flags.BoolP("empty", "E", false, "apply to empty line (ignored)")
	_ = flags.BoolP("initial", "I", false, "apply to initial non-command word (ignored)")
	_ = flags.BoolP("default", "D", false, "apply to commands for which no spec exists (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "complete: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: complete [-abcdefgjklsv] [-A action] [-G globpat] [-W wordlist] [-P prefix] [-S suffix] [-X filterpat] [-F function] [-C command] [name ...]\n")
		fmt.Fprintf(env.Stdout, "Specify how arguments are to be completed by Readline.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "complete")
		return 0
	}

	targets := flags.Args()

	if *pflag_p || (len(args) == 0) {
		// List completions
		cmds := make([]string, 0, len(env.Completions))
		for cmd := range env.Completions {
			cmds = append(cmds, cmd)
		}
		sort.Strings(cmds)

		for _, cmd := range cmds {
			cs := env.Completions[cmd]
			fmt.Fprintf(env.Stdout, "complete %s %s\n", c.formatSpec(cs), cmd)
		}
		return 0
	}

	if *pflag_r {
		if len(targets) == 0 {
			env.Completions = make(map[string]*commands.CompSpec)
		} else {
			for _, target := range targets {
				delete(env.Completions, target)
			}
		}
		return 0
	}

	if len(targets) == 0 {
		return 0
	}

	// Create spec
	cs := &commands.CompSpec{}
	if *flag_a {
		cs.Actions |= 1 << 0
	}
	if *flag_b {
		cs.Actions |= 1 << 1
	}
	if *flag_c {
		cs.Actions |= 1 << 2
	}
	if *flag_d {
		cs.Actions |= 1 << 3
	}
	if *flag_e {
		cs.Actions |= 1 << 4
	}
	if *flag_f {
		cs.Actions |= 1 << 5
	}
	if *flag_g {
		cs.Actions |= 1 << 6
	}
	if *flag_j {
		cs.Actions |= 1 << 7
	}
	if *flag_k {
		cs.Actions |= 1 << 8
	}
	if *flag_s {
		cs.Actions |= 1 << 9
	}
	if *flag_u {
		cs.Actions |= 1 << 10
	}
	if *flag_v {
		cs.Actions |= 1 << 11
	}

	cs.FunctionName = *flag_F
	cs.Command = *flag_C
	cs.GlobPat = *flag_G
	cs.WordList = *flag_W
	cs.Prefix = *flag_P
	cs.Suffix = *flag_S
	cs.FilterPat = *flag_X

	if *flag_A != "" {
		// Basic mapping of action strings to flags if needed
	}

	for _, target := range targets {
		env.Completions[target] = cs
	}

	return 0
}

func (c *Complete) formatSpec(cs *commands.CompSpec) string {
	var parts []string
	if cs.Actions&(1<<0) != 0 {
		parts = append(parts, "-a")
	}
	if cs.Actions&(1<<1) != 0 {
		parts = append(parts, "-b")
	}
	if cs.Actions&(1<<2) != 0 {
		parts = append(parts, "-c")
	}
	if cs.Actions&(1<<3) != 0 {
		parts = append(parts, "-d")
	}
	if cs.Actions&(1<<4) != 0 {
		parts = append(parts, "-e")
	}
	if cs.Actions&(1<<5) != 0 {
		parts = append(parts, "-f")
	}
	if cs.Actions&(1<<6) != 0 {
		parts = append(parts, "-g")
	}
	if cs.Actions&(1<<7) != 0 {
		parts = append(parts, "-j")
	}
	if cs.Actions&(1<<8) != 0 {
		parts = append(parts, "-k")
	}
	if cs.Actions&(1<<9) != 0 {
		parts = append(parts, "-s")
	}
	if cs.Actions&(1<<10) != 0 {
		parts = append(parts, "-u")
	}
	if cs.Actions&(1<<11) != 0 {
		parts = append(parts, "-v")
	}

	if cs.FunctionName != "" {
		parts = append(parts, "-F "+cs.FunctionName)
	}
	if cs.Command != "" {
		parts = append(parts, "-C "+cs.Command)
	}
	if cs.GlobPat != "" {
		parts = append(parts, "-G "+cs.GlobPat)
	}
	if cs.WordList != "" {
		parts = append(parts, "-W "+cs.WordList)
	}
	if cs.Prefix != "" {
		parts = append(parts, "-P "+cs.Prefix)
	}
	if cs.Suffix != "" {
		parts = append(parts, "-S "+cs.Suffix)
	}
	if cs.FilterPat != "" {
		parts = append(parts, "-X "+cs.FilterPat)
	}

	return strings.Join(parts, " ")
}
