package shopt

import (
	"context"
	"fmt"
	"sort"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Shopt struct{}

func New() *Shopt {
	return &Shopt{}
}

func (s *Shopt) Name() string {
	return "shopt"
}

func (s *Shopt) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("shopt", pflag.ContinueOnError)
	set := flags.BoolP("set", "s", false, "enable each optname")
	unset := flags.BoolP("unset", "u", false, "disable each optname")
	quiet := flags.BoolP("quiet", "q", false, "suppress output")
	print := flags.BoolP("print", "p", false, "print each shell option with an indication of its status")
	// -o flag for set -o options is skipped for now to keep it simple

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "shopt: %v\n", err)
		return 1
	}

	targets := flags.Args()

	if *set && *unset {
		fmt.Fprintf(env.Stderr, "shopt: cannot set and unset shell options simultaneously\n")
		return 1
	}

	if len(targets) == 0 {
		// List options
		keys := make([]string, 0, len(env.Shopts))
		for k := range env.Shopts {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			val := env.Shopts[k]
			if *set && !val {
				continue
			}
			if *unset && val {
				continue
			}

			if !*quiet {
				if *print {
					status := "-u"
					if val {
						status = "-s"
					}
					fmt.Fprintf(env.Stdout, "shopt %s %s\n", status, k)
				} else {
					status := "off"
					if val {
						status = "on"
					}
					fmt.Fprintf(env.Stdout, "%-15s\t%s\n", k, status)
				}
			}
		}
		return 0
	}

	exitCode := 0
	for _, target := range targets {
		val, ok := env.Shopts[target]
		if !ok {
			fmt.Fprintf(env.Stderr, "shopt: %s: invalid shell option name\n", target)
			exitCode = 1
			continue
		}

		if *set {
			env.Shopts[target] = true
		} else if *unset {
			env.Shopts[target] = false
		} else {
			// Just query
			if !val {
				exitCode = 1
			}
			if !*quiet {
				if *print {
					status := "-u"
					if val {
						status = "-s"
					}
					fmt.Fprintf(env.Stdout, "shopt %s %s\n", status, target)
				} else {
					status := "off"
					if val {
						status = "on"
					}
					fmt.Fprintf(env.Stdout, "%-15s\t%s\n", target, status)
				}
			}
		}
	}

	return exitCode
}
