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
	si := flags.BoolP("si", "H", false, "lik -h but use powers of 1000 not 1024")
	kb := flags.BoolP("kilobytes", "k", false, "like --block-size=1K")
	all := flags.BoolP("all", "a", false, "include pseudo, duplicate, inaccessible file systems")
	printType := flags.BoolP("print-type", "T", false, "print file system type")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "df: %v\n", err)
		}
		return 1
	}

	// Simulation for virtual environment
	header := fmt.Sprintf("%-15s", "Filesystem")
	if *printType {
		header += fmt.Sprintf(" %-6s", "Type")
	}
	header += fmt.Sprintf(" %10s %10s %10s %5s %s", "Size", "Used", "Avail", "Use%", "Mounted on")
	fmt.Fprintln(env.Stdout, header)
	
	size := int64(1024 * 1024 * 1024) // 1GB
	used := int64(123 * 1024 * 1024) // 123MB
	avail := size - used
	usePct := "12%"

	row := fmt.Sprintf("%-15s", "shm")
	if *printType {
		row += fmt.Sprintf(" %-6s", "tmpfs")
	}

	if *si {
		row += fmt.Sprintf(" %10s %10s %10s %5s %s", "1.1G", "129M", "971M", usePct, "/")
	} else if *humanReadable {
		row += fmt.Sprintf(" %10s %10s %10s %5s %s", "1.0G", "123M", "901M", usePct, "/")
	} else if *kb {
		row += fmt.Sprintf(" %10d %10d %10d %5s %s", size/1024, used/1024, avail/1024, usePct, "/")
	} else {
		// Default is 1k blocks unless POSIXLY_CORRECT
		row += fmt.Sprintf(" %10d %10d %10d %5s %s", size/1024, used/1024, avail/1024, usePct, "/")
	}
	fmt.Fprintln(env.Stdout, row)

	if *all {
		row2 := fmt.Sprintf("%-15s", "proc")
		if *printType {
			row2 += fmt.Sprintf(" %-6s", "proc")
		}
		row2 += fmt.Sprintf(" %10s %10s %10s %5s %s", "0", "0", "0", "-", "/proc")
		fmt.Fprintln(env.Stdout, row2)
	}

	return 0
}
