package sum

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Sum struct {
	name string
}

func New(name string) *Sum {
	return &Sum{name: name}
}

func (s *Sum) Name() string {
	return s.name
}

func (s *Sum) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet(s.name, pflag.ContinueOnError)
	check := flags.BoolP("check", "c", false, "read checksums from the FILEs and check them")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "%s: %v\n", s.name, err)
		return 1
	}

	if *check {
		fmt.Fprintf(env.Stderr, "%s: --check not supported yet\n", s.name)
		return 1
	}

	targets := flags.Args()
	if len(targets) == 0 {
		h := s.getHash()
		if _, err := io.Copy(h, env.Stdin); err != nil {
			fmt.Fprintf(env.Stderr, "%s: %v\n", s.name, err)
			return 1
		}
		fmt.Fprintf(env.Stdout, "%x  -\n", h.Sum(nil))
		return 0
	}

	exitCode := 0
	for _, target := range targets {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		f, err := env.FS.Open(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "%s: %s: %v\n", s.name, target, err)
			exitCode = 1
			continue
		}

		h := s.getHash()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Fprintf(env.Stderr, "%s: %s: %v\n", s.name, target, err)
			f.Close()
			exitCode = 1
			continue
		}
		f.Close()
		fmt.Fprintf(env.Stdout, "%x  %s\n", h.Sum(nil), target)
	}

	return exitCode
}

func (s *Sum) getHash() hash.Hash {
	switch s.name {
	case "md5sum":
		return md5.New()
	case "sha1sum":
		return sha1.New()
	case "sha256sum":
		return sha256.New()
	case "sha512sum":
		return sha512.New()
	default:
		return sha256.New()
	}
}
