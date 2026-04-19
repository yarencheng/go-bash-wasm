package shred

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Shred struct{}

func New() *Shred {
	return &Shred{}
}

func (s *Shred) Name() string {
	return "shred"
}

func (s *Shred) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("shred", pflag.ContinueOnError)
	_ = flags.BoolP("force", "f", false, "change permissions to allow writing if necessary")
	iterations := flags.IntP("iterations", "n", 3, "overwrite N times instead of the default (3)")
	remove := flags.BoolP("remove", "u", false, "truncate and remove file after overwriting")
	verbose := flags.BoolP("verbose", "v", false, "show progress")
	zero := flags.BoolP("zero", "z", false, "add a final overwrite with zeros to hide shredding")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "shred: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		fmt.Fprintf(env.Stderr, "shred: missing file operand\n")
		return 1
	}

	exitCode := 0
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		stat, err := env.FS.Stat(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "shred: %s: %v\n", target, err)
			exitCode = 1
			continue
		}

		size := stat.Size()

		// Overwrite iterations
		for i := 0; i < *iterations; i++ {
			if *verbose {
				fmt.Fprintf(env.Stdout, "shred: %s: pass %d/%d (random)... \n", target, i+1, *iterations)
			}
			if err := s.overwrite(env, fullPath, size, false); err != nil {
				fmt.Fprintf(env.Stderr, "shred: %s: %v\n", target, err)
				exitCode = 1
				break
			}
		}

		if exitCode != 0 {
			continue
		}

		if *zero {
			if *verbose {
				fmt.Fprintf(env.Stdout, "shred: %s: pass %d/%d (000000)... \n", target, *iterations+1, *iterations+1)
			}
			if err := s.overwrite(env, fullPath, size, true); err != nil {
				fmt.Fprintf(env.Stderr, "shred: %s: %v\n", target, err)
				exitCode = 1
				continue
			}
		}

		if *remove {
			if *verbose {
				fmt.Fprintf(env.Stdout, "shred: %s: removing\n", target)
			}
			if err := env.FS.Remove(fullPath); err != nil {
				fmt.Fprintf(env.Stderr, "shred: %s: %v\n", target, err)
				exitCode = 1
			}
		}
	}

	return exitCode
}

func (s *Shred) overwrite(env *commands.Environment, path string, size int64, zero bool) error {
	f, err := env.FS.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	buffer := make([]byte, 4096)
	if zero {
		// all zeros already
	} else {
		for i := range buffer {
			buffer[i] = 0xAA // simple fixed pattern for "random"
		}
	}

	for written := int64(0); written < size; {
		toWrite := int64(len(buffer))
		if size-written < toWrite {
			toWrite = size - written
		}
		n, err := f.Write(buffer[:toWrite])
		if err != nil {
			return err
		}
		written += int64(n)
	}
	return nil
}
