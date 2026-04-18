package touch

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Touch struct{}

func New() *Touch {
	return &Touch{}
}

func (t *Touch) Name() string {
	return "touch"
}

func (t *Touch) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("touch", pflag.ContinueOnError)
	noCreate := flags.BoolP("no-create", "c", false, "do not create any files")
	access := flags.BoolP("access", "a", false, "change only the access time")
	modification := flags.BoolP("modification", "m", false, "change only the modification time")
	reference := flags.StringP("reference", "r", "", "use this file's times instead of current time")
	dateStr := flags.StringP("date", "d", "", "parse STRING and use it instead of current time")
	timeStr := flags.StringP("time", "t", "", "use [[CC]YY]MMDDhhmm[.ss] instead of current time")
	
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "touch: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "touch: missing file operand\n")
		return 1
	}

	exitCode := 0
	atime := time.Now()
	mtime := time.Now()

	if *reference != "" {
		fullRefPath := *reference
		if !filepath.IsAbs(fullRefPath) {
			fullRefPath = filepath.Join(env.Cwd, fullRefPath)
		}
		info, err := env.FS.Stat(fullRefPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "touch: cannot stat '%s': %v\n", *reference, err)
			return 1
		}
		atime = info.ModTime() // In many systems ModTime is same for both if not specified
		mtime = info.ModTime()
	} else if *dateStr != "" {
		t, err := dateparse.ParseLocal(*dateStr)
		if err != nil {
			fmt.Fprintf(env.Stderr, "touch: invalid date format '%s'\n", *dateStr)
			return 1
		}
		atime = t
		mtime = t
	} else if *timeStr != "" {
		t, err := parseTouchTime(*timeStr)
		if err != nil {
			fmt.Fprintf(env.Stderr, "touch: invalid time format '%s'\n", *timeStr)
			return 1
		}
		atime = t
		mtime = t
	}

	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		info, err := env.FS.Stat(fullPath)
		exists := err == nil
		
		if !exists {
			if *noCreate {
				continue
			}
			f, err := env.FS.Create(fullPath)
			if err != nil {
				fmt.Fprintf(env.Stderr, "touch: cannot create '%s': %v\n", target, err)
				exitCode = 1
				continue
			}
			f.Close()
			info, _ = env.FS.Stat(fullPath)
		}

		targetAtime := atime
		targetMtime := mtime

		if *access && !*modification {
			targetMtime = info.ModTime()
		} else if *modification && !*access {
			targetAtime = info.ModTime() // Afero doesn't give us Atime easily, so we use mtime if we must
		}

		err = env.FS.Chtimes(fullPath, targetAtime, targetMtime)
		if err != nil {
			// Some filesystems might not support Chtimes
		}
	}

	return exitCode
}

func parseTouchTime(s string) (time.Time, error) {
	// [[CC]YY]MMDDhhmm[.ss]
	// Possible lengths: 8, 10, 12. With .ss: 11, 13, 15
	dot := ""
	if idx := strings.Index(s, "."); idx != -1 {
		dot = s[idx+1:]
		s = s[:idx]
	}

	var year, month, day, hour, min, sec int
	var err error

	if dot != "" {
		sec, err = strconv.Atoi(dot)
		if err != nil {
			return time.Time{}, err
		}
	}

	now := time.Now()
	switch len(s) {
	case 8: // MMDDhhmm
		month, _ = strconv.Atoi(s[0:2])
		day, _ = strconv.Atoi(s[2:4])
		hour, _ = strconv.Atoi(s[4:6])
		min, _ = strconv.Atoi(s[6:8])
		year = now.Year()
	case 10: // YYMMDDhhmm
		year2, _ := strconv.Atoi(s[0:2])
		if year2 < 69 {
			year = 2000 + year2
		} else {
			year = 1900 + year2
		}
		month, _ = strconv.Atoi(s[2:4])
		day, _ = strconv.Atoi(s[4:6])
		hour, _ = strconv.Atoi(s[6:8])
		min, _ = strconv.Atoi(s[8:10])
	case 12: // CCYYMMDDhhmm
		year, _ = strconv.Atoi(s[0:4])
		month, _ = strconv.Atoi(s[4:6])
		day, _ = strconv.Atoi(s[6:8])
		hour, _ = strconv.Atoi(s[8:10])
		min, _ = strconv.Atoi(s[10:12])
	default:
		return time.Time{}, fmt.Errorf("invalid length")
	}

	return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local), nil
}
