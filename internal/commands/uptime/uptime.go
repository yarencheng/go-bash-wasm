package uptime

import (
	"context"
	"fmt"
	"time"

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
	now := time.Now()
	up := now.Sub(env.StartTime)
	
	// Format: 17:26:15 up 1:23, 1 user, load average: 0.01, 0.05, 0.02
	hours := int(up.Hours())
	minutes := int(up.Minutes()) % 60
	
	fmt.Fprintf(env.Stdout, " %s up %d:%02d, 1 user, load average: 0.00, 0.00, 0.00\n",
		now.Format("15:04:05"), hours, minutes)
	
	return 0
}
