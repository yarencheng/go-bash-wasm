package od

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Od struct{}

func New() *Od {
	return &Od{}
}

func (o *Od) Name() string {
	return "od"
}

func (o *Od) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("od", pflag.ContinueOnError)
	addressRadix := flags.StringP("address-radix", "A", "o", "output radix for file offsets (d, o, x, n)")
	skip := flags.IntP("skip-bytes", "j", 0, "skip bytes from the beginning of input")
	readLimit := flags.IntP("read-bytes", "N", -1, "limit dump to BYTES input bytes")
	width := flags.IntP("width", "w", 16, "output BYTES bytes per line")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "od: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return o.process(env, env.Stdin, *addressRadix, *skip, *readLimit, *width)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "od: %s: %v\n", arg, err)
			exitCode = 1
			continue
		}
		if status := o.process(env, f, *addressRadix, *skip, *readLimit, *width); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

func (o *Od) process(env *commands.Environment, r io.Reader, addrRadix string, skip, limit, width int) int {
	if skip > 0 {
		if seeker, ok := r.(io.Seeker); ok {
			seeker.Seek(int64(skip), io.SeekStart)
		} else {
			io.CopyN(io.Discard, r, int64(skip))
		}
	}

	var reader io.Reader = r
	if limit >= 0 {
		reader = io.LimitReader(r, int64(limit))
	}

	buf := make([]byte, width)
	offset := int64(skip)

	for {
		n, err := io.ReadFull(reader, buf)
		if n > 0 {
			o.printLine(env, offset, buf[:n], addrRadix, width)
			offset += int64(n)
		}
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			fmt.Fprintf(env.Stderr, "od: %v\n", err)
			return 1
		}
	}

	// Print final offset
	if addrRadix != "n" {
		o.printAddress(env, offset, addrRadix)
		fmt.Fprintln(env.Stdout)
	}

	return 0
}

func (o *Od) printAddress(env *commands.Environment, offset int64, radix string) {
	switch radix {
	case "d":
		fmt.Fprintf(env.Stdout, "%07d", offset)
	case "o":
		fmt.Fprintf(env.Stdout, "%07o", offset)
	case "x":
		fmt.Fprintf(env.Stdout, "%07x", offset)
	case "n":
		// no address
	default:
		fmt.Fprintf(env.Stdout, "%07o", offset)
	}
}

func (o *Od) printLine(env *commands.Environment, offset int64, data []byte, radix string, width int) {
	if radix != "n" {
		o.printAddress(env, offset, radix)
	}

	// Default behavior is 2-byte octal words
	for i := 0; i < len(data); i += 2 {
		if i+1 < len(data) {
			val := uint16(data[i]) | uint16(data[i+1])<<8
			fmt.Fprintf(env.Stdout, " %06o", val)
		} else {
			fmt.Fprintf(env.Stdout, " %03o", data[i])
		}
	}
	fmt.Fprintln(env.Stdout)
}
