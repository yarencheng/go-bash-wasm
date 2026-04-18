package uptime

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Uptime struct{}

func New() *Uptime {
	return &Uptime{}
}

func (u *Uptime) Name() string {
	return "uptime"
}

func (u *Uptime) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("uptime", pflag.ContinueOnError)
	pretty := flags.BoolP("pretty", "p", false, "show uptime in pretty format")
	since := flags.BoolP("since", "s", false, "system up since")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "uptime: %v\n", err)
		return 1
	}

	now := time.Now()
	up := now.Sub(env.StartTime)

	if *since {
		fmt.Fprintln(env.Stdout, env.StartTime.Format("2006-01-02 15:04:05"))
		return 0
	}

	if *pretty {
		days := int(up.Hours()) / 24
		hours := int(up.Hours()) % 24
		minutes := int(up.Minutes()) % 60

		var res []string
		if days > 0 {
			res = append(res, fmt.Sprintf("%d days", days))
		}
		if hours > 0 {
			res = append(res, fmt.Sprintf("%d hours", hours))
		}
		if minutes > 0 {
			res = append(res, fmt.Sprintf("%d minutes", minutes))
		}
		if len(res) == 0 {
			res = append(res, "0 minutes")
		}
		fmt.Fprintf(env.Stdout, "up %s\n", strings.Join(res, ", "))
		return 0
	}

	// Format: 17:26:15 up 1:23, 1 user, load average: 0.01, 0.05, 0.02
	hours := int(up.Hours())
	minutes := int(up.Minutes()) % 60

	fmt.Fprintf(env.Stdout, " %s up %d:%02d, 1 user, load average: 0.00, 0.00, 0.00\n",
		now.Format("15:04:05"), hours, minutes)

	return 0
}
