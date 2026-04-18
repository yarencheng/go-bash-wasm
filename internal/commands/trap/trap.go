package trap

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Trap struct{}

func New() *Trap {
	return &Trap{}
}

func (t *Trap) Name() string {
	return "trap"
}

func (t *Trap) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("trap", pflag.ContinueOnError)
	print := flags.BoolP("print", "p", false, "display current traps")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "trap: %v\n", err)
		return 1
	}

	targets := flags.Args()

	if *print || (len(targets) == 0 && len(args) == 0) {
		// List traps
		sigs := make([]string, 0, len(env.Traps))
		for sig := range env.Traps {
			sigs = append(sigs, sig)
		}
		sort.Strings(sigs)

		for _, sig := range sigs {
			cmd := env.Traps[sig]
			fmt.Fprintf(env.Stdout, "trap -- '%s' %s\n", strings.ReplaceAll(cmd, "'", "'\\''"), sig)
		}
		return 0
	}

	if len(targets) == 0 {
		return 0
	}

	// trap [COMMAND] SIGSPEC ...
	command := targets[0]
	signals := targets[1:]

	// If the first argument is a signal name, then we are resetting
	// But bash is complex here. If it's a number, it's a signal.
	// If it's a string, it might be a command.

	// Check if the first argument is a signal
	isFirstArgSignal := false
	if _, ok := t.getSignal(command); ok {
		isFirstArgSignal = true
	}

	if isFirstArgSignal {
		// Reset signals
		for _, sig := range targets {
			if s, ok := t.getSignal(sig); ok {
				delete(env.Traps, s)
			} else {
				fmt.Fprintf(env.Stderr, "trap: %s: invalid signal specification\n", sig)
			}
		}
		return 0
	}

	if len(signals) == 0 {
		// Just trap command with no signals? Error in bash
		return 0
	}

	for _, sig := range signals {
		if s, ok := t.getSignal(sig); ok {
			if command == "-" || command == "" {
				delete(env.Traps, s)
			} else {
				env.Traps[s] = command
			}
		} else {
			fmt.Fprintf(env.Stderr, "trap: %s: invalid signal specification\n", sig)
		}
	}

	return 0
}

func (t *Trap) getSignal(sig string) (string, bool) {
	sig = strings.ToUpper(sig)
	if strings.HasPrefix(sig, "SIG") {
		sig = sig[3:]
	}

	// Common signals
	valid := map[string]bool{
		"EXIT": true, "DEBUG": true, "RETURN": true, "ERR": true,
		"HUP": true, "INT": true, "QUIT": true, "ILL": true, "TRAP": true,
		"ABRT": true, "BUS": true, "FPE": true, "KILL": true, "USR1": true,
		"SEGV": true, "USR2": true, "PIPE": true, "ALRM": true, "TERM": true,
		"STKFLT": true, "CHLD": true, "CONT": true, "STOP": true, "TSTP": true,
		"TTIN": true, "TTOU": true, "URG": true, "XCPU": true, "XFSZ": true,
		"VTALRM": true, "PROF": true, "WINCH": true, "IO": true, "PWR": true,
		"SYS": true,
	}

	if valid[sig] {
		return sig, true
	}

	return "", false
}
