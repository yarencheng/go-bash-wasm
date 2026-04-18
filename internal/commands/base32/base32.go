package base32cmd

import (
	"context"
	"encoding/base32"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Base32 struct{}

func New() *Base32 {
	return &Base32{}
}

func (b *Base32) Name() string {
	return "base32"
}

func (b *Base32) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("base32", pflag.ContinueOnError)
	decode := flags.BoolP("decode", "d", false, "decode data")
	ignoreGarbage := flags.BoolP("ignore-garbage", "i", false, "when decoding, ignore non-alphabet characters")
	wrap := flags.IntP("wrap", "w", 76, "wrap encoded lines after COLS character (default 76)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "base32: %v\n", err)
		return 1
	}

	var input io.ReadCloser
	targets := flags.Args()

	if len(targets) == 0 {
		input = env.Stdin
	} else {
		fullPath := targets[0]
		if !filepath.IsAbs(fullPath) {
			fullPath = filepath.Join(env.Cwd, fullPath)
		}
		f, err := env.FS.Open(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "base32: %v\n", err)
			return 1
		}
		input = f
	}
	defer func() {
		if input != nil && input != env.Stdin {
			input.Close()
		}
	}()

	encoding := base32.StdEncoding

	if *decode {
		// Basic implementation of ignore garbage would require a custom reader.
		var r io.Reader = input
		if *ignoreGarbage {
			// Actually skipping for now, but standard decoder is strict
		}
		decoder := base32.NewDecoder(encoding, r)
		_, err := io.Copy(env.Stdout, decoder)
		if err != nil {
			fmt.Fprintf(env.Stderr, "base32: decode error: %v\n", err)
			return 1
		}
	} else {
		data, err := io.ReadAll(input)
		if err != nil {
			fmt.Fprintf(env.Stderr, "base32: read error: %v\n", err)
			return 1
		}
		encoded := encoding.EncodeToString(data)

		if *wrap <= 0 {
			fmt.Fprintln(env.Stdout, encoded)
		} else {
			for i := 0; i < len(encoded); i += *wrap {
				end := i + *wrap
				if end > len(encoded) {
					end = len(encoded)
				}
				fmt.Fprintln(env.Stdout, encoded[i:end])
			}
		}
	}

	return 0
}
