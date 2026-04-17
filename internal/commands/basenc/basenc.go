package basenc

import (
	"context"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Basenc struct{}

func New() *Basenc {
	return &Basenc{}
}

func (b *Basenc) Name() string {
	return "basenc"
}

func (b *Basenc) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("basenc", pflag.ContinueOnError)
	decode := flags.BoolP("decode", "d", false, "decode data")
	ignoreGarbage := flags.BoolP("ignore-garbage", "i", false, "when decoding, ignore non-alphabet characters")
	wrap := flags.IntP("wrap", "w", 76, "wrap encoded lines after COLS character (default 76)")

	fBase16 := flags.Bool("base16", false, "base16 (hex) encoding/decoding")
	fBase32 := flags.Bool("base32", false, "base32 encoding/decoding")
	fBase32hex := flags.Bool("base32hex", false, "base32hex encoding/decoding")
	fBase64 := flags.Bool("base64", false, "base64 encoding/decoding")
	fBase64url := flags.Bool("base64url", false, "base64url encoding/decoding")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "basenc: %v\n", err)
		return 1
	}

	targets := flags.Args()
	var input io.ReadCloser
	if len(targets) == 0 {
		input = env.Stdin
	} else {
		fullPath := targets[0]
		if !filepath.IsAbs(fullPath) {
			fullPath = filepath.Join(env.Cwd, fullPath)
		}
		f, err := env.FS.Open(fullPath)
		if err != nil {
			fmt.Fprintf(env.Stderr, "basenc: %v\n", err)
			return 1
		}
		input = f
	}
	defer func() {
		if input != nil && input != env.Stdin {
			input.Close()
		}
	}()

	var encoder func([]byte) string
	var decoder func(io.Reader) io.Reader

	if *fBase16 {
		encoder = hex.EncodeToString
		decoder = func(r io.Reader) io.Reader { return hex.NewDecoder(r) }
	} else if *fBase32 {
		enc := base32.StdEncoding
		encoder = enc.EncodeToString
		decoder = func(r io.Reader) io.Reader { return base32.NewDecoder(enc, r) }
	} else if *fBase32hex {
		enc := base32.HexEncoding
		encoder = enc.EncodeToString
		decoder = func(r io.Reader) io.Reader { return base32.NewDecoder(enc, r) }
	} else if *fBase64 {
		enc := base64.StdEncoding
		encoder = enc.EncodeToString
		decoder = func(r io.Reader) io.Reader { return base64.NewDecoder(enc, r) }
	} else if *fBase64url {
		enc := base64.URLEncoding
		encoder = enc.EncodeToString
		decoder = func(r io.Reader) io.Reader { return base64.NewDecoder(enc, r) }
	} else {
		fmt.Fprintf(env.Stderr, "basenc: missing encoding flag\n")
		return 1
	}

	if *decode {
		var r io.Reader = input
		if *ignoreGarbage {
			r = &garbageFilter{r, *fBase16, *fBase32 || *fBase32hex, *fBase64 || *fBase64url}
		}
		_, err := io.Copy(env.Stdout, decoder(r))
		if err != nil {
			fmt.Fprintf(env.Stderr, "basenc: decode error: %v\n", err)
			return 1
		}
	} else {
		data, err := io.ReadAll(input)
		if err != nil {
			fmt.Fprintf(env.Stderr, "basenc: read error: %v\n", err)
			return 1
		}
		encoded := encoder(data)
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

type garbageFilter struct {
	r      io.Reader
	hex    bool
	base32 bool
	base64 bool
}

func (gf *garbageFilter) Read(p []byte) (n int, err error) {
	temp := make([]byte, len(p))
	tn, terr := gf.r.Read(temp)
	if tn == 0 {
		return 0, terr
	}

	for i := 0; i < tn; i++ {
		c := temp[i]
		valid := false
		if gf.hex {
			valid = (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
		} else if gf.base32 {
			valid = (c >= 'A' && c <= 'Z') || (c >= '2' && c <= '7') || c == '='
		} else if gf.base64 {
			valid = (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '+' || c == '/' || c == '_' || c == '-' || c == '='
		}
		if valid {
			p[n] = c
			n++
		}
	}
	if n == 0 && terr == nil {
		// if we filtered everything, try reading more
		return gf.Read(p)
	}
	return n, terr
}
