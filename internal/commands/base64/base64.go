package base64cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Base64 struct{}

func New() *Base64 {
	return &Base64{}
}

func (b *Base64) Name() string {
	return "base64"
}

func (b *Base64) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("base64", pflag.ContinueOnError)
	decode := flags.BoolP("decode", "d", false, "decode data")
	ignoreGarbage := flags.BoolP("ignore-garbage", "i", false, "when decoding, ignore non-alphabet characters")
	wrap := flags.IntP("wrap", "w", 76, "wrap encoded lines after COLS character (default 76)")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "base64: %v\n", err)
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
			fmt.Fprintf(env.Stderr, "base64: %v\n", err)
			return 1
		}
		input = f
	}
	defer func() {
		if input != env.Stdin {
			input.Close()
		}
	}()

	if *decode {
		decoder := base64.NewDecoder(base64.StdEncoding, input)
		if *ignoreGarbage {
			// Basic implementation of ignore garbage would require a custom reader.
			// Skipping for now as standard decoder is strict.
		}
		_, err := io.Copy(env.Stdout, decoder)
		if err != nil {
			fmt.Fprintf(env.Stderr, "base64: decode error: %v\n", err)
			return 1
		}
	} else {
		data, err := io.ReadAll(input)
		if err != nil {
			fmt.Fprintf(env.Stderr, "base64: read error: %v\n", err)
			return 1
		}
		encoded := base64.StdEncoding.EncodeToString(data)

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
