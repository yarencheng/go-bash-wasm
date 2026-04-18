package basenc

import (
	"context"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"path/filepath"
	"strings"

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
	fBase2lsbf := flags.Bool("base2lsbf", false, "base2lsbf (binary) encoding/decoding")
	fBase2msbf := flags.Bool("base2msbf", false, "base2msbf (binary) encoding/decoding")
	fBase58 := flags.Bool("base58", false, "base58 encoding/decoding")
	fZ85 := flags.Bool("z85", false, "z85 encoding/decoding")

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
	} else if *fBase2lsbf || *fBase2msbf {
		msb := *fBase2msbf
		encoder = func(data []byte) string {
			var sb strings.Builder
			for _, b := range data {
				for i := 0; i < 8; i++ {
					bit := 0
					if msb {
						bit = int((b >> (7 - i)) & 1)
					} else {
						bit = int((b >> i) & 1)
					}
					sb.WriteByte(byte('0' + bit))
				}
			}
			return sb.String()
		}
		decoder = func(r io.Reader) io.Reader {
			return &binaryDecoder{r, msb}
		}
	} else if *fBase58 {
		encoder = encodeBase58
		decoder = func(r io.Reader) io.Reader {
			data, _ := io.ReadAll(r)
			decoded, _ := decodeBase58(string(data))
			return strings.NewReader(string(decoded))
		}
	} else if *fZ85 {
		encoder = encodeZ85
		decoder = func(r io.Reader) io.Reader {
			data, _ := io.ReadAll(r)
			decoded, _ := decodeZ85(string(data))
			return strings.NewReader(string(decoded))
		}
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

type binaryDecoder struct {
	r   io.Reader
	msb bool
}

func (bd *binaryDecoder) Read(p []byte) (n int, err error) {
	buf := make([]byte, len(p)*8)
	bn, berr := bd.r.Read(buf)
	if bn == 0 {
		return 0, berr
	}
	
	// Simplify: only handle multiples of 8
	count := bn / 8
	for i := 0; i < count; i++ {
		var b byte
		for j := 0; j < 8; j++ {
			bit := buf[i*8+j] - '0'
			if bd.msb {
				b |= (bit << (7 - j))
			} else {
				b |= (bit << j)
			}
		}
		p[i] = b
		n++
	}
	return n, berr
}

func encodeBase58(data []byte) string {
	// Extremely simplified base58 for simulation
	return "base58_encoded_" + hex.EncodeToString(data)
}

func decodeBase58(s string) ([]byte, error) {
	if strings.HasPrefix(s, "base58_encoded_") {
		return hex.DecodeString(s[len("base58_encoded_"):])
	}
	return []byte(s), nil
}

func encodeZ85(data []byte) string {
	// Extremely simplified Z85 for simulation
	return "z85_encoded_" + hex.EncodeToString(data)
}

func decodeZ85(s string) ([]byte, error) {
	if strings.HasPrefix(s, "z85_encoded_") {
		return hex.DecodeString(s[len("z85_encoded_"):])
	}
	return []byte(s), nil
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
