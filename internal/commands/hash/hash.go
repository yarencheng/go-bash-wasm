package hashcmd

import (
	"context"
	"fmt"
	"sort"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Hash struct{}

func New() *Hash {
	return &Hash{}
}

func (h *Hash) Name() string {
	return "hash"
}

func (h *Hash) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("hash", pflag.ContinueOnError)
	reset := flags.BoolP("reset", "r", false, "forget all remembered locations")

	_ = flags.StringP("delete", "d", "", "forget the remembered location of each NAME (ignored)")
	_ = flags.StringP("path", "p", "", "use PATH as the full pathname of NAME (ignored)")
	_ = flags.BoolP("list", "l", false, "display in a format that may be reused as input (ignored)")
	_ = flags.BoolP("table", "t", false, "print the remembered location of each NAME, preceding each location with the corresponding NAME (ignored)")
	help := flags.Bool("help", false, "display this help and exit")
	version := flags.Bool("version", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "hash: %v\n", err)
		return 1
	}

	if *help {
		fmt.Fprintf(env.Stdout, "Usage: hash [-lr] [-p pathname] [-dt] [name ...]\n")
		fmt.Fprintf(env.Stdout, "Remember or display program locations.\n\n")
		flags.PrintDefaults()
		return 0
	}

	if *version {
		commands.ShowVersion(env.Stdout, "hash")
		return 0
	}

	if *reset {
		for k := range env.Hash {
			delete(env.Hash, k)
		}
		return 0
	}

	targets := flags.Args()
	if len(targets) == 0 {
		if len(env.Hash) == 0 {
			fmt.Fprintln(env.Stdout, "hash: hash table empty")
			return 0
		}

		fmt.Fprintln(env.Stdout, "hits\tcommand")
		keys := make([]string, 0, len(env.Hash))
		for k := range env.Hash {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(env.Stdout, "0\t%s\n", env.Hash[k]) // simplified 0 hits
		}
		return 0
	}

	for _, name := range targets {
		if _, ok := env.Registry.Get(name); ok {
			env.Hash[name] = "/bin/" + name // simplified path
		} else {
			fmt.Fprintf(env.Stderr, "hash: %s: not found\n", name)
		}
	}

	return 0
}
