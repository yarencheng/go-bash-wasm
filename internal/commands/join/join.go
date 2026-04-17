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

	// Simple implementation: load all of file 2 into a map
	f2Data := make(map[string][]string)
	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {
		line := scanner2.Text()
		parts := strings.Fields(line)
		if len(parts) >= *field2 {
			key := parts[*field2-1]
			// Store remaining parts
			var others []string
			for i, p := range parts {
				if i != *field2-1 {
					others = append(others, p)
				}
			}
			f2Data[key] = others
		}
	}

	scanner1 := bufio.NewScanner(f1)
	for scanner1.Scan() {
		line := scanner1.Text()
		parts := strings.Fields(line)
		if len(parts) >= *field1 {
			key := parts[*field1-1]
			if others2, ok := f2Data[key]; ok {
				// Output key, then others from f1, then others from f2
				var out []string
				out = append(out, key)
				for i, p := range parts {
					if i != *field1-1 {
						out = append(out, p)
					}
				}
				out = append(out, others2...)
				fmt.Fprintln(env.Stdout, strings.Join(out, " "))
			}
		}
	}

	return 0
}
