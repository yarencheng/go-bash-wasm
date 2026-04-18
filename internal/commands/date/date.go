package date

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Date struct{}

func New() *Date {
	return &Date{}
}

func (d *Date) Name() string {
	return "date"
}

func (d *Date) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("date", pflag.ContinueOnError)
	utc := flags.BoolP("utc", "u", false, "print or set Coordinated Universal Time (UTC)")
	rfc2822 := flags.BoolP("rfc-email", "R", false, "output date and time in RFC 2822 format")
	iso8601 := flags.StringP("iso-8601", "I", "", "output date/time in ISO 8601 format. FMT='date' for date only (default), 'hours', 'minutes', 'seconds', or 'ns'")
	flags.Lookup("iso-8601").NoOptDefVal = "date"
	dateStr := flags.StringP("date", "d", "", "display time described by STRING, not 'now'")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "date: %v\n", err)
		return 1
	}

	now := time.Now()
	if *dateStr != "" {
		// Basic format support
		layouts := []string{
			time.RFC3339,
			"2006-01-02",
			"2006-01-02 15:04:05",
			time.UnixDate,
			time.ANSIC,
		}
		parsed := false
		for _, layout := range layouts {
			if t, err := time.Parse(layout, *dateStr); err == nil {
				now = t
				parsed = true
				break
			}
		}
		if !parsed {
			// Try parsing as duration relative to now
			if dur, err := time.ParseDuration(*dateStr); err == nil {
				now = time.Now().Add(dur)
			} else {
				fmt.Fprintf(env.Stderr, "date: invalid date '%s'\n", *dateStr)
				return 1
			}
		}
	}

	if *utc {
		now = now.UTC()
	}

	remainingArgs := flags.Args()
	format := time.UnixDate

	if *rfc2822 {
		format = time.RFC1123Z
	} else if flags.Changed("iso-8601") {
		switch *iso8601 {
		case "hours":
			format = "2006-01-02T15"
		case "minutes":
			format = "2006-01-02T15:04"
		case "seconds":
			format = "2006-01-02T15:04:05"
		case "ns":
			format = "2006-01-02T15:04:05.000000000"
		default: // date or empty
			format = "2006-01-02"
		}
		// ISO 8601 usually includes timezone if not date only
		if *iso8601 != "" && *iso8601 != "date" {
			format += "-07:00"
		}
	}

	if len(remainingArgs) > 0 && strings.HasPrefix(remainingArgs[0], "+") {
		f := remainingArgs[0][1:]
		fmt.Fprintln(env.Stdout, formatLinuxDate(now, f))
		return 0
	}

	fmt.Fprintln(env.Stdout, now.Format(format))
	return 0
}

func formatLinuxDate(t time.Time, f string) string {
	mapping := map[string]string{
		"%Y": "2006",
		"%y": "06",
		"%m": "01",
		"%d": "02",
		"%H": "15",
		"%M": "04",
		"%S": "05",
		"%b": "Jan",
		"%h": "Jan",
		"%B": "January",
		"%a": "Mon",
		"%A": "Monday",
		"%z": "-0700",
		"%Z": "MST",
		"%T": "15:04:05",
		"%R": "15:04",
		"%F": "2006-01-02",
		"%%": "%",
	}
	res := f
	for k, v := range mapping {
		res = strings.ReplaceAll(res, k, t.Format(v))
	}
	return res
}
