package df

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type fsInfo struct {
	fs, fstype    string
	size, used    int64
	itotal, iused int64
	mount         string
	pseudo        bool
}

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
	inodes := flags.BoolP("inodes", "i", false, "list inode information instead of block usage")
	local := flags.BoolP("local", "l", false, "limit listing to local file systems")
	portability := flags.BoolP("portability", "P", false, "use the POSIX output format")
	total := flags.Bool("total", false, "elide all entries insignificant to available space, and produce a grand total")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "df: %v\n", err)
		return 1
	}

	data := []fsInfo{
		{"shm", "tmpfs", 1024 * 1024 * 1024, 123 * 1024 * 1024, 1000000, 100, "/", false},
		{"proc", "proc", 0, 0, 0, 0, "/proc", true},
		{"sys", "sysfs", 0, 0, 0, 0, "/sys", true},
		{"dev", "devtmpfs", 512 * 1024 * 1024, 10 * 1024 * 1024, 100000, 50, "/dev", true},
	}

	// Filter local
	if *local {
		// All simulated ones are local for now
	}

	// Filter all
	if !*all {
		var filtered []fsInfo
		for _, d := range data {
			if !d.pseudo || d.size > 0 {
				filtered = append(filtered, d)
			}
		}
		data = filtered
	}

	// Header
	var header strings.Builder
	if *portability {
		fmt.Fprintf(&header, "%-15s %-10s %-10s %-10s %-5s %s", "Filesystem", "1024-blocks", "Used", "Available", "Capacity", "Mounted on")
	} else if *inodes {
		fmt.Fprintf(&header, "%-15s", "Filesystem")
		if *printType {
			fmt.Fprintf(&header, " %-6s", "Type")
		}
		fmt.Fprintf(&header, " %10s %10s %10s %5s %s", "Inodes", "IUsed", "IFree", "IUse%", "Mounted on")
	} else {
		fmt.Fprintf(&header, "%-15s", "Filesystem")
		if *printType {
			fmt.Fprintf(&header, " %-6s", "Type")
		}
		fmt.Fprintf(&header, " %10s %10s %10s %5s %s", "1K-blocks", "Used", "Available", "Use%", "Mounted on")
	}
	fmt.Fprintln(env.Stdout, header.String())

	var totalSize, totalUsed, totalAvail int64
	var totalItotal, totalIused int64

	for _, d := range data {
		totalSize += d.size
		totalUsed += d.used
		totalAvail += (d.size - d.used)
		totalItotal += d.itotal
		totalIused += d.iused

		fmt.Fprintln(env.Stdout, formatRow(d, *inodes, *printType, *humanReadable, *si, *kb, *portability))
	}

	if *total {
		t := fsInfo{"total", "-", totalSize, totalUsed, totalItotal, totalIused, "-", false}
		fmt.Fprintln(env.Stdout, formatRow(t, *inodes, *printType, *humanReadable, *si, *kb, *portability))
	}

	return 0
}

func formatRow(d fsInfo, inodes, printType, h, si, kb, p bool) string {
	var row strings.Builder
	fmt.Fprintf(&row, "%-15s", d.fs)
	if printType && !p {
		fmt.Fprintf(&row, " %-6s", d.fstype)
	}

	if inodes {
		if d.itotal == 0 && d.pseudo {
			fmt.Fprintf(&row, " %10s %10s %10s %5s %s", "-", "-", "-", "-", d.mount)
		} else {
			free := d.itotal - d.iused
			pct := "0%"
			if d.itotal > 0 {
				pct = fmt.Sprintf("%d%%", d.iused*100/d.itotal)
			}
			fmt.Fprintf(&row, " %10d %10d %10d %5s %s", d.itotal, d.iused, free, pct, d.mount)
		}
	} else {
		s, u, a := d.size, d.used, d.size-d.used
		pct := "0%"
		if s > 0 {
			pct = fmt.Sprintf("%d%%", u*100/s)
		}

		if h || si {
			row.WriteString(fmt.Sprintf(" %10s %10s %10s %5s %s", formatSize(s, si), formatSize(u, si), formatSize(a, si), pct, d.mount))
		} else {
			div := int64(1024)
			row.WriteString(fmt.Sprintf(" %10d %10d %10d %5s %s", s/div, u/div, a/div, pct, d.mount))
		}
	}
	return row.String()
}

func formatSize(n int64, si bool) string {
	div := float64(1024)
	if si {
		div = 1000
	}
	if n < int64(div) {
		return fmt.Sprintf("%dB", n)
	}
	units := []string{"K", "M", "G", "T"}
	f := float64(n)
	u := ""
	for _, unit := range units {
		f /= div
		u = unit
		if f < div {
			break
		}
	}
	return fmt.Sprintf("%.1f%s", f, u)
}
