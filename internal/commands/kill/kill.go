package kill

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Kill struct{}

func New() *Kill {
	return &Kill{}
}

func (k *Kill) Name() string {
	return "kill"
}

var signals = map[int]string{
	1:  "HUP",
	2:  "INT",
	3:  "QUIT",
	4:  "ILL",
	5:  "TRAP",
	6:  "ABRT",
	7:  "BUS",
	8:  "FPE",
	9:  "KILL",
	10: "USR1",
	11: "SEGV",
	12: "USR2",
	13: "PIPE",
	14: "ALRM",
	15: "TERM",
	17: "CHLD",
	18: "CONT",
	19: "STOP",
	20: "TSTP",
}

func (k *Kill) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("kill", pflag.ContinueOnError)
	list := flags.BoolP("list", "l", false, "list signal names")
	
	// kill has weird flag parsing where -SIGNAL is common.
	// pflag doesn't handle -9 or -TERM easily as flags.
	// We'll manually check for those.
	
	if len(args) > 0 && strings.HasPrefix(args[0], "-") && !strings.HasPrefix(args[0], "--") && len(args[0]) > 1 {
		// Might be -9 or -TERM
		if !*list {
			// handled below
		}
	}

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "kill: %v\n", err)
		return 1
	}

	if *list {
		h := flags.Args()
		if len(h) == 0 {
			// print all
			keys := make([]int, 0, len(signals))
			for k := range signals {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			for i, key := range keys {
				fmt.Fprintf(env.Stdout, "%2d) SIG%-8s", key, signals[key])
				if (i+1)%4 == 0 {
					fmt.Fprintln(env.Stdout, "")
				}
			}
			if len(keys)%4 != 0 {
				fmt.Fprintln(env.Stdout, "")
			}
			return 0
		}
		// print specific
		for _, arg := range h {
			n, err := strconv.Atoi(arg)
			if err == nil {
				if sig, ok := signals[n]; ok {
					fmt.Fprintln(env.Stdout, sig)
				} else {
					fmt.Fprintf(env.Stderr, "kill: %s: invalid signal number\n", arg)
				}
			} else {
				// name to number? 
				found := false
				name := strings.ToUpper(strings.TrimPrefix(arg, "SIG"))
				for num, sig := range signals {
					if sig == name {
						fmt.Fprintln(env.Stdout, num)
						found = true
						break
					}
				}
				if !found {
					fmt.Fprintf(env.Stderr, "kill: %s: invalid signal specification\n", arg)
				}
			}
		}
		return 0
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "kill: usage: kill [-s sigspec | -n signum | -sigspec] pid | jobspec ... or kill -l [sigspec]\n")
		return 2
	}

	exitCode := 0
	for _, pidStr := range targets {
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			fmt.Fprintf(env.Stderr, "kill: %s: arguments must be process or job IDs\n", pidStr)
			exitCode = 1
			continue
		}

		// In simulator, we only have 'process' 1 (the shell itself) or nothing.
		if pid == 1 {
			// Signaling self? 
			// For now just ignore but return success if it's a known signal.
		} else {
			fmt.Fprintf(env.Stderr, "kill: (%d) - no such process\n", pid)
			exitCode = 1
		}
	}

	return exitCode
}
