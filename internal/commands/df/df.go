package df

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Df struct{}

func New() *Df {
	return &Df{}
}

func (d *Df) Name() string {
	return "df"
}

func (d *Df) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("df", pflag.ContinueOnError)
	humanReadable := flags.BoolP("human-readable", "h", false, "print sizes in human readable format (e.g., 1K 234M 2G)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "df: %v\n", err)
		return 1
	}

	// Simulation for virtual environment
	fmt.Fprintf(env.Stdout, "%-15s %10s %10s %10s %5s %s\n", "Filesystem", "Size", "Used", "Avail", "Use%", "Mounted on")
	
	size := int64(1024 * 1024 * 1024) // 1GB
	used := int64(123 * 1024 * 1024) // 123MB
	avail := size - used
	usePct := "12%"

	if *humanReadable {
		fmt.Fprintf(env.Stdout, "%-15s %10s %10s %10s %5s %s\n", "shm", "1.0G", "123M", "901M", usePct, "/")
	} else {
		fmt.Fprintf(env.Stdout, "%-15s %10d %10d %10d %5s %s\n", "shm", size/1024, used/1024, avail/1024, usePct, "/")
	}

	return 0
}
