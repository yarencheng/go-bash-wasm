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
	table := flags.BoolP("table", "t", false, "print a table of signal information")
	sigName := flags.StringP("signal", "s", "", "specify the signal to be sent")
	sigNum := flags.IntP("signum", "n", -1, "specify the signal number to be sent")

	// kill has weird flag parsing where -SIGNAL is common.
	// We'll manually extract the signal if it's in the first argument.
	var customSig string
	var pflagArgs = args
	if len(args) > 0 && strings.HasPrefix(args[0], "-") && !strings.HasPrefix(args[0], "--") && len(args[0]) > 1 {
		potentialSig := args[0][1:]
		// Check if it's a number
		if _, err := strconv.Atoi(potentialSig); err == nil {
			customSig = potentialSig
			pflagArgs = args[1:]
		} else {
			// Check if it's a known signal name
			name := strings.ToUpper(strings.TrimPrefix(potentialSig, "SIG"))
			for _, s := range signals {
				if s == name {
					customSig = potentialSig
					pflagArgs = args[1:]
					break
				}
			}
		}
	}

	if err := flags.Parse(pflagArgs); err != nil {
		fmt.Fprintf(env.Stderr, "kill: %v\n", err)
		return 1
	}

	if *list || *table {
		h := flags.Args()
		if len(h) == 0 {
			// print all
			keys := make([]int, 0, len(signals))
			for k := range signals {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			for i, key := range keys {
				if *table {
					fmt.Fprintf(env.Stdout, "%d\t%s\n", key, signals[key])
				} else {
					fmt.Fprintf(env.Stdout, "%2d) SIG%-8s", key, signals[key])
					if (i+1)%4 == 0 {
						fmt.Fprintln(env.Stdout, "")
					}
				}
			}
			if !*table && len(keys)%4 != 0 {
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

	signal := 15 // Default SIGTERM
	if customSig != "" {
		if n, err := strconv.Atoi(customSig); err == nil {
			signal = n
		} else {
			name := strings.ToUpper(strings.TrimPrefix(customSig, "SIG"))
			for num, s := range signals {
				if s == name {
					signal = num
					break
				}
			}
		}
	} else if *sigNum != -1 {
		signal = *sigNum
	} else if *sigName != "" {
		name := strings.ToUpper(strings.TrimPrefix(*sigName, "SIG"))
		found := false
		for num, s := range signals {
			if s == name {
				signal = num
				found = true
				break
			}
		}
		if !found {
			fmt.Fprintf(env.Stderr, "kill: %s: invalid signal specification\n", *sigName)
			return 1
		}
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
			_ = signal
		} else {
			// Check if it's a job ID
			foundJob := false
			for _, job := range env.Jobs {
				if job.PID == pid {
					foundJob = true
					// Mock signaling the job
					if signal == 9 || signal == 15 {
						job.Status = "Done"
					} else if signal == 19 || signal == 20 {
						job.Status = "Stopped"
					} else if signal == 18 {
						job.Status = "Running"
					}
					break
				}
			}
			if !foundJob {
				fmt.Fprintf(env.Stderr, "kill: (%d) - no such process\n", pid)
				exitCode = 1
			}
		}
	}

	return exitCode
}
