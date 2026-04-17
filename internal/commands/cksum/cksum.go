package cksum

import (
	"bufio"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Cksum struct{}

func New() *Cksum {
	return &Cksum{}
}

func (c *Cksum) Name() string {
	return "cksum"
}

type CksumOptions struct {
	Algorithm string
	Check     bool
	Zero      bool
	Quiet     bool
	Status    bool
	Strict    bool
	Warn      bool
}

func (c *Cksum) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("cksum", pflag.ContinueOnError)
	algorithm := flags.StringP("algorithm", "a", "crc", "select hashing algorithm (crc, md5, sha1, sha256, sha512)")
	check := flags.BoolP("check", "c", false, "read checksums from the FILEs and check them")
	zero := flags.BoolP("zero", "z", false, "end each output line with NUL, not newline")
	quiet := flags.Bool("quiet", false, "don't print OK for each successfully verified file")
	status := flags.Bool("status", false, "don't output anything, status code shows success")
	strict := flags.Bool("strict", false, "exit non-zero for improperly formatted checksum lines")
	warn := flags.BoolP("warn", "w", false, "warn about improperly formatted checksum lines")
	flags.Parse(args)

	// Combine flags into a struct or pass them down
	opts := CksumOptions{
		Algorithm: *algorithm,
		Check:     *check,
		Zero:      *zero,
		Quiet:     *quiet,
		Status:    *status,
		Strict:    *strict,
		Warn:      *warn,
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return c.Process(env, env.Stdin, "", opts)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			if !opts.Status {
				fmt.Fprintf(env.Stderr, "cksum: %s: %v\n", arg, err)
			}
			exitCode = 1
			continue
		}
		if status := c.Process(env, f, arg, opts); status != 0 {
			exitCode = status
		}
		f.Close()
	}

	return exitCode
}

func (c *Cksum) Process(env *commands.Environment, r io.Reader, name string, opts CksumOptions) int {
	if opts.Check {
		return c.RunCheck(env, r, opts)
	}

	var h hash.Hash
	switch opts.Algorithm {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha224":
		h = sha256.New224()
	case "sha256":
		h = sha256.New()
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	default:
		// Default to CRC
		return c.ProcessCRC(env, r, name, opts)
	}

	_, err := io.Copy(h, r)
	if err != nil {
		if !opts.Status {
			fmt.Fprintf(env.Stderr, "cksum: %v\n", err)
		}
		return 1
	}

	if !opts.Status {
		terminator := "\n"
		if opts.Zero {
			terminator = "\x00"
		}
		if name != "" {
			fmt.Fprintf(env.Stdout, "%x  %s%s", h.Sum(nil), name, terminator)
		} else {
			fmt.Fprintf(env.Stdout, "%x%s", h.Sum(nil), terminator)
		}
	}

	return 0
}

func (c *Cksum) ProcessCRC(env *commands.Environment, r io.Reader, name string, opts CksumOptions) int {
	// Standard posix cksum uses CRC-32 with a specific table and includes length
	table := crc32.MakeTable(0x04C11DB7)
	h := crc32.New(table)
	
	buf := make([]byte, 32*1024)
	var length int64
	for {
		n, err := r.Read(buf)
		if n > 0 {
			h.Write(buf[:n])
			length += int64(n)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(env.Stderr, "cksum: %v\n", err)
			return 1
		}
	}

	// Append length in little-endian like some POSIX implementations
	lenBuf := make([]byte, 8)
	tmpLen := length
	for i := 0; tmpLen > 0; i++ {
		lenBuf[i] = byte(tmpLen & 0xFF)
		tmpLen >>= 8
		h.Write(lenBuf[i : i+1])
	}

	sum := h.Sum32()
	if !opts.Status {
		terminator := "\n"
		if opts.Zero {
			terminator = "\x00"
		}
		if name != "" {
			fmt.Fprintf(env.Stdout, "%d %d %s%s", sum, length, name, terminator)
		} else {
			fmt.Fprintf(env.Stdout, "%d %d%s", sum, length, terminator)
		}
	}
	return 0
}

func (c *Cksum) RunCheck(env *commands.Environment, r io.Reader, opts CksumOptions) int {
	scanner := bufio.NewScanner(r)
	exitCode := 0
	totalCount := 0
	failCount := 0
	improperCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			improperCount++
			if opts.Warn {
				fmt.Fprintf(env.Stderr, "cksum: improperly formatted line: %s\n", line)
			}
			continue
		}

		totalCount++
		expected := parts[0]
		fileName := parts[1]

		f, err := env.FS.Open(fileName)
		if err != nil {
			if !opts.Status {
				fmt.Fprintf(env.Stderr, "cksum: %s: %v\n", fileName, err)
			}
			exitCode = 1
			failCount++
			continue
		}

		var h hash.Hash
		switch len(expected) {
		case 32:
			h = md5.New()
		case 40:
			h = sha1.New()
		case 56:
			h = sha256.New224()
		case 64:
			h = sha256.New()
		case 96:
			h = sha512.New384()
		case 128:
			h = sha512.New()
		default:
			improperCount++
			if opts.Warn {
				fmt.Fprintf(env.Stderr, "cksum: %s: unknown checksum format\n", fileName)
			}
			f.Close()
			continue
		}

		io.Copy(h, f)
		actual := fmt.Sprintf("%x", h.Sum(nil))
		f.Close()

		if actual == expected {
			if !opts.Status && !opts.Quiet {
				fmt.Fprintf(env.Stdout, "%s: OK\n", fileName)
			}
		} else {
			if !opts.Status {
				fmt.Fprintf(env.Stdout, "%s: FAILED\n", fileName)
			}
			failCount++
			exitCode = 1
		}
	}

	if opts.Strict && improperCount > 0 {
		exitCode = 1
	}

	if failCount > 0 && !opts.Status && !opts.Quiet {
		fmt.Fprintf(env.Stderr, "cksum: WARNING: %d computed checksum did NOT match\n", failCount)
	}

	return exitCode
}
