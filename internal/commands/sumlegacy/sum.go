package sumlegacy

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type SumLegacy struct{}

func New() *SumLegacy {
	return &SumLegacy{}
}

func (s *SumLegacy) Name() string {
	return "sum"
}

func (s *SumLegacy) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("sum", pflag.ContinueOnError)
	sysv := flags.BoolP("sysv", "s", false, "use System V sum algorithm")
	_ = flags.BoolP("bsd", "r", false, "use BSD sum algorithm (default)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "sum: %v\n", err)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		return s.process(env, env.Stdin, "", *sysv)
	}

	exitCode := 0
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		f, err := env.FS.Open(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "sum: %s: %v\n", target, err)
			exitCode = 1
			continue
		}

		if res := s.process(env, f, target, *sysv); res != 0 {
			exitCode = res
		}
		f.Close()
	}

	return exitCode
}

func (s *SumLegacy) process(env *commands.Environment, r io.Reader, name string, sysv bool) int {
	if sysv {
		return s.sysv(env, r, name)
	}
	return s.bsd(env, r, name)
}

func (s *SumLegacy) bsd(env *commands.Environment, r io.Reader, name string) int {
	var checksum uint16
	var length int64
	buf := make([]byte, 8192)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			for i := 0; i < n; i++ {
				length++
				checksum = (checksum >> 1) + (checksum << 15)
				checksum += uint16(buf[i])
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(env.Stderr, "sum: %v\n", err)
			return 1
		}
	}

	blocks := (length + 1023) / 1024
	if name != "" {
		fmt.Fprintf(env.Stdout, "%05d %5d %s\n", checksum, blocks, name)
	} else {
		fmt.Fprintf(env.Stdout, "%05d %5d\n", checksum, blocks)
	}
	return 0
}

func (s *SumLegacy) sysv(env *commands.Environment, r io.Reader, name string) int {
	var checksum uint32
	var length int64
	buf := make([]byte, 8192)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			for i := 0; i < n; i++ {
				checksum += uint32(buf[i])
				length++
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(env.Stderr, "sum: %v\n", err)
			return 1
		}
	}

	rVal := (checksum & 0xffff) + (checksum >> 16)
	checksum = (rVal & 0xffff) + (rVal >> 16)

	blocks := (length + 511) / 512
	if name != "" {
		fmt.Fprintf(env.Stdout, "%d %d %s\n", checksum, blocks, name)
	} else {
		fmt.Fprintf(env.Stdout, "%d %d\n", checksum, blocks)
	}
	return 0
}
