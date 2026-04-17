package truncate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Truncate struct{}

func New() *Truncate {
	return &Truncate{}
}

func (t *Truncate) Name() string {
	return "truncate"
}

func (t *Truncate) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("truncate", pflag.ContinueOnError)
	noCreate := flags.BoolP("no-create", "c", false, "do not create any files")
	_ = flags.BoolP("io-blocks", "o", false, "interpret SIZE as number of IO blocks")
	reference := flags.StringP("reference", "r", "", "use this file's size")
	sizeStr := flags.StringP("size", "s", "", "set or adjust the file size by SIZE bytes")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "truncate: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		fmt.Fprintf(env.Stderr, "truncate: missing file operand\n")
		return 1
	}

	var targetSize int64
	if *reference != "" {
		stat, err := env.FS.Stat(*reference)
		if err != nil {
			fmt.Fprintf(env.Stderr, "truncate: cannot stat '%s': %v\n", *reference, err)
			return 1
		}
		targetSize = stat.Size()
	}

	if *sizeStr != "" {
		// Basic parser for size. Support +, -.
		rel := 0 // 0: set, 1: add, -1: sub
		valStr := *sizeStr
		if strings.HasPrefix(*sizeStr, "+") {
			rel = 1
			valStr = (*sizeStr)[1:]
		} else if strings.HasPrefix(*sizeStr, "-") {
			rel = -1
			valStr = (*sizeStr)[1:]
		}

		// Support simple multipliers like K, M, G
		val, multiplier, err := parseSize(valStr)
		if err != nil {
			fmt.Fprintf(env.Stderr, "truncate: invalid size '%s': %v\n", *sizeStr, err)
			return 1
		}
		
		valSize := val * multiplier
		if rel == 0 {
			targetSize = valSize
		} else {
			// Relative adjustment happens per file. We'll handle it in the loop.
		}
	}

	exitCode := 0
	for _, arg := range remaining {
		var finalSize int64 = targetSize
		
		if strings.HasPrefix(*sizeStr, "+") || strings.HasPrefix(*sizeStr, "-") {
			stat, err := env.FS.Stat(arg)
			if err != nil {
				if !*noCreate {
					// If it doesn't exist and we would create it, size is 0.
					stat = nil
				} else {
					fmt.Fprintf(env.Stderr, "truncate: cannot stat '%s': %v\n", arg, err)
					exitCode = 1
					continue
				}
			}
			
			currentSize := int64(0)
			if stat != nil {
				currentSize = stat.Size()
			}
			
			val, multiplier, _ := parseSize((*sizeStr)[1:])
			adj := val * multiplier
			if strings.HasPrefix(*sizeStr, "+") {
				finalSize = currentSize + adj
			} else {
				finalSize = currentSize - adj
				if finalSize < 0 {
					finalSize = 0
				}
			}
		}

		mode := os.O_WRONLY
		if !*noCreate {
			mode |= os.O_CREATE
		}
		
		f, err := env.FS.OpenFile(arg, mode, 0666)
		if err != nil {
			if !(*noCreate && os.IsNotExist(err)) {
				fmt.Fprintf(env.Stderr, "truncate: cannot open '%s' for writing: %v\n", arg, err)
				exitCode = 1
			}
			continue
		}
		
		if err := f.Truncate(finalSize); err != nil {
			fmt.Fprintf(env.Stderr, "truncate: failed to truncate '%s' to %d bytes: %v\n", arg, finalSize, err)
			exitCode = 1
		}
		f.Close()
	}

	return exitCode
}

func parseSize(s string) (int64, int64, error) {
	if s == "" {
		return 0, 0, fmt.Errorf("empty size")
	}
	
	multiplier := int64(1)
	valStr := s
	
	last := s[len(s)-1]
	switch last {
	case 'K', 'k':
		multiplier = 1024
		valStr = s[:len(s)-1]
	case 'M', 'm':
		multiplier = 1024 * 1024
		valStr = s[:len(s)-1]
	case 'G', 'g':
		multiplier = 1024 * 1024 * 1024
		valStr = s[:len(s)-1]
	}

	val, err := strconv.ParseInt(valStr, 10, 64)
	return val, multiplier, err
}
