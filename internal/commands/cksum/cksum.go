package cksum

import (
	"bufio"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
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
	Algorithm     string
	Check         bool
	Zero          bool
	Quiet         bool
	Status        bool
	Strict        bool
	Warn          bool
	Base64        bool
	Raw           bool
	Tag           bool
	Untagged      bool
	IgnoreMissing bool
	Length        int
}

func (c *Cksum) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("cksum", pflag.ContinueOnError)
	algorithm := flags.StringP("algorithm", "a", "crc", "select hashing algorithm (crc, md5, sha1, sha224, sha256, sha384, sha512)")
	check := flags.BoolP("check", "c", false, "read checksums from the FILEs and check them")
	zero := flags.BoolP("zero", "z", false, "end each output line with NUL, not newline")
	quiet := flags.Bool("quiet", false, "don't print OK for each successfully verified file")
	status := flags.Bool("status", false, "don't output anything, status code shows success")
	strict := flags.Bool("strict", false, "exit non-zero for improperly formatted checksum lines")
	warn := flags.BoolP("warn", "w", false, "warn about improperly formatted checksum lines")
	base64 := flags.Bool("base64", false, "emit base64-encoded checksums")
	raw := flags.Bool("raw", false, "emit raw binary checksums")
	tag := flags.Bool("tag", false, "create a BSD-style checksum")
	untagged := flags.Bool("untagged", false, "create a reverse-style checksum, without algorithm name")
	ignoreMissing := flags.Bool("ignore-missing", false, "don't fail or report status for missing files")
	length := flags.IntP("length", "l", 0, "digest length in bits; must not exceed the maximum for the blake2 algorithm")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "cksum: %v\n", err)
		return 1
	}

	opts := CksumOptions{
		Algorithm:     *algorithm,
		Check:         *check,
		Zero:          *zero,
		Quiet:         *quiet,
		Status:        *status,
		Strict:        *strict,
		Warn:          *warn,
		Base64:        *base64,
		Raw:           *raw,
		Tag:           *tag,
		Untagged:      *untagged,
		IgnoreMissing: *ignoreMissing,
		Length:        *length,
	}

	remaining := flags.Args()
	if len(remaining) == 0 {
		return c.Process(env, env.Stdin, "", opts)
	}

	exitCode := 0
	for _, arg := range remaining {
		f, err := env.FS.Open(arg)
		if err != nil {
			if opts.Check && opts.IgnoreMissing {
				continue
			}
			if !opts.Status {
				fmt.Fprintf(env.Stderr, "cksum: %s: %v\n", arg, err)
			}
			exitCode = 1
			continue
		}
		if res := c.Process(env, f, arg, opts); res != 0 {
			exitCode = res
		}
		f.Close()
	}

	return exitCode
}

func (c *Cksum) Process(env *commands.Environment, r io.Reader, name string, opts CksumOptions) int {
	if opts.Check {
		return c.RunCheck(env, r, opts)
	}

	if opts.Algorithm == "crc" || opts.Algorithm == "" {
		return c.ProcessCRC(env, r, name, opts)
	}

	var h hash.Hash
	algoName := strings.ToUpper(opts.Algorithm)
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
		fmt.Fprintf(env.Stderr, "cksum: %s: unknown algorithm\n", opts.Algorithm)
		return 1
	}

	_, err := io.Copy(h, r)
	if err != nil {
		if !opts.Status {
			fmt.Fprintf(env.Stderr, "cksum: %v\n", err)
		}
		return 1
	}

	if opts.Status {
		return 0
	}

	sum := h.Sum(nil)
	if opts.Length > 0 && opts.Length/8 < len(sum) {
		sum = sum[:opts.Length/8]
	}

	terminator := "\n"
	if opts.Zero {
		terminator = "\x00"
	}

	var output string
	if opts.Raw {
		env.Stdout.Write(sum)
		return 0
	}

	encoded := ""
	if opts.Base64 {
		encoded = base64.StdEncoding.EncodeToString(sum)
	} else {
		encoded = hex.EncodeToString(sum)
	}

	if opts.Tag {
		if name != "" {
			output = fmt.Sprintf("%s (%s) = %s", algoName, name, encoded)
		} else {
			output = fmt.Sprintf("%s = %s", algoName, encoded)
		}
	} else if opts.Untagged {
		output = encoded
	} else {
		if name != "" {
			output = fmt.Sprintf("%s  %s", encoded, name)
		} else {
			output = encoded
		}
	}

	fmt.Fprintf(env.Stdout, "%s%s", output, terminator)
	return 0
}

func (c *Cksum) ProcessCRC(env *commands.Environment, r io.Reader, name string, opts CksumOptions) int {
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

	// POSIX cksum length handling
	tmpLen := length
	for tmpLen > 0 {
		h.Write([]byte{byte(tmpLen & 0xFF)})
		tmpLen >>= 8
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
	failCount := 0
	improperCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		
		var expected, fileName string
		if strings.Contains(line, "(") && strings.Contains(line, ") = ") {
			// BSD style: ALGO (FILE) = SUM
			start := strings.Index(line, "(")
			end := strings.Index(line, ") = ")
			if start != -1 && end != -1 {
				fileName = line[start+1 : end]
				expected = line[end+4:]
			}
		} else if len(parts) >= 2 {
			expected = parts[0]
			fileName = parts[1]
		}

		if fileName == "" || expected == "" {
			improperCount++
			if opts.Warn {
				fmt.Fprintf(env.Stderr, "cksum: improperly formatted line: %s\n", line)
			}
			continue
		}

		f, err := env.FS.Open(fileName)
		if err != nil {
			if opts.IgnoreMissing {
				continue
			}
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
			// Try base64
			if opts.Base64 {
				// Base64 lengths are different
			}
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

