package joincmd

import (
	"bufio"
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Join struct{}

func New() *Join {
	return &Join{}
}

func (j *Join) Name() string {
	return "join"
}

func (j *Join) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("join", pflag.ContinueOnError)
	field1 := flags.IntP("field1", "1", 1, "join on this FIELD of file 1")
	field2 := flags.IntP("field2", "2", 1, "join on this FIELD of file 2")
	ignoreCase := flags.BoolP("ignore-case", "i", false, "ignore differences in case when comparing fields")
	delim := flags.StringP("tab", "t", "", "use CHAR as input and output field separator")
	auto1 := flags.Bool("a1", false, "also print unpairable lines from file 1")
	auto2 := flags.Bool("a2", false, "also print unpairable lines from file 2")
	unpair1 := flags.Bool("v1", false, "like -a1, but suppress joined output lines")
	unpair2 := flags.Bool("v2", false, "like -a2, but suppress joined output lines")
	// Legacy -a N and -v N support? GNU join uses -a 1, -a 2. pflag doesn't handle -a 1 easily with Int if we want -a alone to mean something else.
	// We'll support --a=1 and --v=1 via manual check or flags.
	var aFilter, vFilter int
	flags.IntVarP(&aFilter, "unpairable", "a", 0, "also print unpairable lines from file FILENUM, where FILENUM is 1 or 2")
	flags.IntVarP(&vFilter, "suppress", "v", 0, "like -a FILENUM, but suppress joined output lines")
	empty := flags.StringP("empty", "e", "", "replace empty fields with EMPTY")
	format := flags.StringP("format", "o", "", "obey FORMAT while constructing output line")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "join: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) < 2 {
		fmt.Fprintf(env.Stderr, "join: missing operand\n")
		return 1
	}

	f1Path := targets[0]
	if !filepath.IsAbs(f1Path) {
		f1Path = filepath.Join(env.Cwd, f1Path)
	}
	f1, err := env.FS.Open(f1Path)
	if err != nil {
		fmt.Fprintf(env.Stderr, "join: %v\n", err)
		return 1
	}
	defer f1.Close()

	f2Path := targets[1]
	if !filepath.IsAbs(f2Path) {
		f2Path = filepath.Join(env.Cwd, f2Path)
	}
	f2, err := env.FS.Open(f2Path)
	if err != nil {
		fmt.Fprintf(env.Stderr, "join: %v\n", err)
		return 1
	}
	defer f2.Close()

	if *auto1 || aFilter == 1 {
		aFilter = 1
	}
	if *auto2 || aFilter == 2 {
		aFilter = 2
	}
	if *unpair1 || vFilter == 1 {
		vFilter = 1
	}
	if *unpair2 || vFilter == 2 {
		vFilter = 2
	}

	splitFunc := strings.Fields
	if *delim != "" {
		splitFunc = func(s string) []string {
			return strings.Split(s, *delim)
		}
	}

	// Simple implementation: load all of file 2 into a slice to support ordered join
	type lineInfo struct {
		key   string
		parts []string
		used  bool
	}
	var f2Lines []lineInfo
	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {
		line := scanner2.Text()
		parts := splitFunc(line)
		if len(parts) >= *field2 {
			key := parts[*field2-1]
			if *ignoreCase {
				key = strings.ToLower(key)
			}
			f2Lines = append(f2Lines, lineInfo{key: key, parts: parts, used: false})
		}
	}

	scanner1 := bufio.NewScanner(f1)
	for scanner1.Scan() {
		line := scanner1.Text()
		parts := splitFunc(line)
		if len(parts) >= *field1 {
			key := parts[*field1-1]
			matchKey := key
			if *ignoreCase {
				matchKey = strings.ToLower(key)
			}

			matched := false
			for i := range f2Lines {
				if f2Lines[i].key == matchKey {
					matched = true
					f2Lines[i].used = true
					if vFilter == 0 {
						if *format != "" {
							fmt.Fprintln(env.Stdout, j.formatOutput(*format, key, parts, f2Lines[i].parts, *field1, *field2, *empty, *delim))
						} else {
							// Default output
							var out []string
							out = append(out, key)
							for j, p := range parts {
								if j != *field1-1 {
									out = append(out, p)
								}
							}
							for j, p := range f2Lines[i].parts {
								if j != *field2-1 {
									out = append(out, p)
								}
							}
							fmt.Fprintln(env.Stdout, strings.Join(out, getJoinDelim(*delim)))
						}
					}
				}
			}

			if !matched && (aFilter == 1 || vFilter == 1) {
				if *format != "" {
					fmt.Fprintln(env.Stdout, j.formatOutput(*format, key, parts, nil, *field1, *field2, *empty, *delim))
				} else {
					fmt.Fprintln(env.Stdout, line)
				}
			}
		} else if aFilter == 1 || vFilter == 1 {
			fmt.Fprintln(env.Stdout, line)
		}
	}

	if aFilter == 2 || vFilter == 2 {
		for _, info := range f2Lines {
			if !info.used {
				if *format != "" {
					fmt.Fprintln(env.Stdout, j.formatOutput(*format, info.key, nil, info.parts, *field1, *field2, *empty, *delim))
				} else {
					fmt.Fprintln(env.Stdout, strings.Join(info.parts, getJoinDelim(*delim)))
				}
			}
		}
	}

	return 0
}

func (j *Join) formatOutput(format, key string, p1, p2 []string, f1, f2 int, empty, delim string) string {
	d := getJoinDelim(delim)
	parts := strings.Fields(format)
	var res []string
	for _, p := range parts {
		if p == "0" {
			res = append(res, key)
			continue
		}
		numField := strings.Split(p, ".")
		if len(numField) < 2 {
			res = append(res, p)
			continue
		}
		fileNum := numField[0]
		fieldIdx := 0
		fmt.Sscanf(numField[1], "%d", &fieldIdx)

		val := empty
		if fileNum == "1" && p1 != nil && fieldIdx <= len(p1) {
			val = p1[fieldIdx-1]
		} else if fileNum == "2" && p2 != nil && fieldIdx <= len(p2) {
			val = p2[fieldIdx-1]
		}
		res = append(res, val)
	}
	return strings.Join(res, d)
}

func getJoinDelim(d string) string {
	if d == "" {
		return " "
	}
	return d
}
