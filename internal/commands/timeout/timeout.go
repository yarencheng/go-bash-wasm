package timeout

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Timeout struct{}

func New() *Timeout {
	return &Timeout{}
}

func (t *Timeout) Name() string {
	return "timeout"
}

func (t *Timeout) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("timeout", pflag.ContinueOnError)
	_ = flags.Bool("foreground", false, "run command in the foreground")
	_ = flags.StringP("kill-after", "k", "", "also send a KILL signal if command is still running")
	_ = flags.StringP("signal", "s", "TERM", "specify the signal to be sent on timeout")
	preserveCode := flags.Bool("preserve-status", false, "exit with the same status as COMMAND")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "timeout: %v\n", err)
		return 125
	}

	remaining := flags.Args()
	if len(remaining) < 2 {
		fmt.Fprintf(env.Stderr, "timeout: too few arguments\n")
		return 125
	}

	durationStr := remaining[0]
	cmdName := remaining[1]
	cmdArgs := remaining[1:]

	duration, err := parseDuration(durationStr)
	if err != nil {
		fmt.Fprintf(env.Stderr, "timeout: invalid time interval '%s'\n", durationStr)
		return 125
	}

	childCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	targetCmd, exists := env.Registry.Get(cmdName)
	if !exists {
		fmt.Fprintf(env.Stderr, "timeout: failed to run command '%s': No such file or directory\n", cmdName)
		return 127
	}

	// We don't have real processes or signals here, so we just rely on context cancellation.
	// In a real system, timeout would send a signal. Here we just expect the command to respect ctx.
	exitCode := targetCmd.Run(childCtx, env, cmdArgs)

	if childCtx.Err() == context.DeadlineExceeded {
		if !*preserveCode {
			return 124
		}
	}

	return exitCode
}

func parseDuration(s string) (time.Duration, error) {
	if s == "" {
		return 0, fmt.Errorf("empty duration")
	}
	
	unit := s[len(s)-1]
	valStr := s
	multiplier := time.Second
	
	switch unit {
	case 's':
		valStr = s[:len(s)-1]
		multiplier = time.Second
	case 'm':
		valStr = s[:len(s)-1]
		multiplier = time.Minute
	case 'h':
		valStr = s[:len(s)-1]
		multiplier = time.Hour
	case 'd':
		valStr = s[:len(s)-1]
		multiplier = 24 * time.Hour
	default:
		// default is seconds if no unit
		multiplier = time.Second
	}

	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return 0, err
	}
	
	return time.Duration(val * float64(multiplier)), nil
}
