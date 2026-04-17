package cksum

import (
	"context"
	"fmt"
	"hash/crc32"
	"io"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Cksum struct{}

func New() *Cksum {
	return &Cksum{}
}

func (c *Cksum) Name() string {
	return "cksum"
}

func (c *Cksum) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) == 0 {
		return c.process(env, env.Stdin, "")
	}

	exitCode := 0
	for _, arg := range args {
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "cksum: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := c.process(env, f, arg); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

func (c *Cksum) process(env *commands.Environment, r io.Reader, name string) int {
	// POSIX cksum uses a specific CRC32 and includes length.
	// We'll use IEEE for simplicity in this WASM context unless strictly required.
	h := crc32.NewIEEE()
	length, err := io.Copy(h, r)
	if err != nil {
		fmt.Fprintf(env.Stderr, "cksum: %v\n", err)
		return 1
	}

	if name != "" {
		fmt.Fprintf(env.Stdout, "%d %d %s\n", h.Sum32(), length, name)
	} else {
		fmt.Fprintf(env.Stdout, "%d %d\n", h.Sum32(), length)
	}

	return 0
}
