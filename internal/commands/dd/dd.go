package dd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type DD struct {
	name string
}

func New(name string) *DD {
	return &DD{name: name}
}

func (d *DD) Name() string {
	return d.name
}

func (d *DD) Run(ctx context.Context, env *commands.Environment, args []string) int {
	var inputPath, outputPath string
	bs := int64(512)
	count := int64(-1)
	skip := int64(0)
	seek := int64(0)
	notrunc := false

	for _, arg := range args {
		if arg == "--help" {
			fmt.Fprintf(env.Stdout, "Usage: dd [OPERAND]...\n")
			fmt.Fprintf(env.Stdout, "  or:  dd OPTION\n")
			fmt.Fprintf(env.Stdout, "Copy a file, converting and formatting according to the operands.\n\n")
			fmt.Fprintf(env.Stdout, "  bs=BYTES        read and write up to BYTES bytes at a time (default: 512)\n")
			fmt.Fprintf(env.Stdout, "  count=N         copy only N input blocks\n")
			fmt.Fprintf(env.Stdout, "  if=FILE         read from FILE instead of stdin\n")
			fmt.Fprintf(env.Stdout, "  of=FILE         write to FILE instead of stdout\n")
			fmt.Fprintf(env.Stdout, "  seek=N          skip N obs-sized blocks at start of output\n")
			fmt.Fprintf(env.Stdout, "  skip=N          skip N ibs-sized blocks at start of input\n")
			fmt.Fprintf(env.Stdout, "\nOperands that are currently ignored (stubs for parity):\n")
			fmt.Fprintf(env.Stdout, "  cbs=BYTES       convert BYTES bytes at a time\n")
			fmt.Fprintf(env.Stdout, "  conv=CONVS      convert the file as per the comma separated symbol list\n")
			fmt.Fprintf(env.Stdout, "  ibs=BYTES       read up to BYTES bytes at a time (default: 512)\n")
			fmt.Fprintf(env.Stdout, "  obs=BYTES       write up to BYTES bytes at a time (default: 512)\n")
			fmt.Fprintf(env.Stdout, "  iflag=FLAGS     read as per the comma separated symbol list\n")
			fmt.Fprintf(env.Stdout, "  oflag=FLAGS     write as per the comma separated symbol list\n")
			fmt.Fprintf(env.Stdout, "  status=LEVEL    The LEVEL of information to print to stderr\n")
			return 0
		}
		if arg == "--version" {
			commands.ShowVersion(env.Stdout, "dd")
			return 0
		}

		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key, val := parts[0], parts[1]
		switch key {
		case "if":
			inputPath = val
		case "of":
			outputPath = val
		case "bs":
			bs = d.parseSize(val)
		case "count":
			count = d.parseSize(val)
		case "skip":
			skip = d.parseSize(val)
		case "seek":
			seek = d.parseSize(val)
		case "conv":
			if val == "notrunc" {
				notrunc = true
			}
		case "ibs", "obs", "cbs", "iflag", "oflag", "status":
			// Ignored for now
		}
	}

	var in io.Reader = env.Stdin
	if inputPath != "" {
		f, err := env.FS.Open(d.absPath(env, inputPath))
		if err != nil {
			fmt.Fprintf(env.Stderr, "dd: failed to open %s: %v\n", inputPath, err)
			return 1
		}
		defer f.Close()
		if skip > 0 {
			if seeker, ok := f.(io.Seeker); ok {
				seeker.Seek(skip*bs, io.SeekStart)
			} else {
				io.CopyN(io.Discard, f, skip*bs)
			}
		}
		in = f
	} else if skip > 0 {
		io.CopyN(io.Discard, in, skip*bs)
	}

	var out io.Writer = env.Stdout
	if outputPath != "" {
		var f afero.File
		var err error
		path := d.absPath(env, outputPath)
		if notrunc {
			f, err = env.FS.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
		} else {
			f, err = env.FS.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		}
		if err != nil {
			fmt.Fprintf(env.Stderr, "dd: failed to open %s: %v\n", outputPath, err)
			return 1
		}
		defer f.Close()
		if seek > 0 {
			if seeker, ok := f.(io.Seeker); ok {
				seeker.Seek(seek*bs, io.SeekStart)
			}
		}
		out = f
	}

	buf := make([]byte, bs)
	var totalRead int64
	for count < 0 || totalRead < count {
		n, err := io.ReadFull(in, buf)
		if n > 0 {
			_, werr := out.Write(buf[:n])
			if werr != nil {
				fmt.Fprintf(env.Stderr, "dd: write error: %v\n", werr)
				return 1
			}
			totalRead++
		}
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			fmt.Fprintf(env.Stderr, "dd: read error: %v\n", err)
			return 1
		}
	}

	return 0
}

func (d *DD) parseSize(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func (d *DD) absPath(env *commands.Environment, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(env.Cwd, path)
}
