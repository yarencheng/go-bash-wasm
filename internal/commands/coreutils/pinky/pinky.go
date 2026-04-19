package pinky

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Pinky struct{}

func New() *Pinky {
	return &Pinky{}
}

func (p *Pinky) Name() string {
	return "pinky"
}

func (p *Pinky) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("pinky", pflag.ContinueOnError)
	_ = flags.BoolP("long", "l", false, "produce long format output")
	_ = flags.BoolP("short", "s", false, "produce short format output")
	_ = flags.BoolP("no-heading", "f", false, "omit the line of column headings in short format")
	_ = flags.BoolP("no-project", "w", false, "omit the user's project file in long format")
	_ = flags.BoolP("no-plan", "i", false, "omit the user's plan file in long format")
	_ = flags.BoolP("no-real-name", "q", false, "omit the user's real name in short format")
	_ = flags.BoolP("no-idle", "b", false, "omit the user's idle time in short format")
	_ = flags.BoolP("no-host", "h", false, "omit the user's remote host in short format")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "pinky: %v\n", err)
		return 1
	}

	// Stub: In this simulator, we only have one user "user".
	user := os.Getenv("USER")
	if user == "" {
		user = "user"
	}

	fmt.Fprintf(env.Stdout, "Login    Name                 TTY      Idle   When             Where\n")
	fmt.Fprintf(env.Stdout, "%-8s %-20s %-8s %-6s %-16s %-s\n", user, user, "tty1", "0:00", "2026-04-18 09:57", "localhost")

	return 0
}
