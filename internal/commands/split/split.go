package split

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Split struct{}

func New() *Split {
	return &Split{}
}

func (s *Split) Name() string {
	return "split"
}

func (s *Split) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("split", pflag.ContinueOnError)
	lines := flags.IntP("lines", "l", 1000, "put NUMBER lines per output file")
	bytes := flags.StringP("bytes", "b", "", "put SIZE bytes per output file")
	suffixLen := flags.IntP("suffix-length", "a", 2, "use suffixes of length N (default 2)")
	numericSuffix := flags.BoolP("numeric-suffixes", "d", false, "use numeric suffixes instead of alphabetic")
	verbose := flags.Bool("verbose", false, "print a diagnostic just before each output file is opened")
	numChunks := flags.StringP("number", "n", "", "generate CHUNKS output files")
	separator := flags.StringP("separator", "t", "", "use SEP instead of newline as the record separator")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "split: %v\n", err)
		return 1
	}

	remaining := flags.Args()
	var inputFile string
	prefix := "x"

	if len(remaining) > 0 {
		inputFile = remaining[0]
	}
	if len(remaining) > 1 {
		prefix = remaining[1]
	}

	var r io.Reader
	if inputFile == "" || inputFile == "-" {
		r = env.Stdin
	} else {
		f, err := env.FS.Open(inputFile)
		if err != nil {
			fmt.Fprintf(env.Stderr, "split: %s: %v\n", inputFile, err)
			return 1
		}
		defer f.Close()
		r = f
	}

	byteLimit := int64(-1)
	if *bytes != "" {
		// Basic parser for bytes, e.g. "10k", "1m". For now just support raw numbers.
		val, err := strconv.ParseInt(*bytes, 10, 64)
		if err != nil {
			fmt.Fprintf(env.Stderr, "split: invalid number of bytes: %s\n", *bytes)
			return 1
		}
		byteLimit = val
	}

	
	
	if *numChunks != "" {
		n, err := strconv.Atoi(*numChunks)
		if err == nil && n > 0 {
			// Find total size if possible
			var size int64 = -1
			if inputFile != "" && inputFile != "-" {
				info, err := env.FS.Stat(inputFile)
				if err == nil {
					size = info.Size()
				}
			}
			if size >= 0 {
				byteLimit = (size + int64(n) - 1) / int64(n)
			}
		}
	}

	if byteLimit > 0 {
		return s.splitBytes(env, r, prefix, byteLimit, *suffixLen, *numericSuffix, *verbose)
	}
	return s.splitLines(env, r, prefix, *lines, *suffixLen, *numericSuffix, *verbose, *separator)
}

func (s *Split) splitLines(env *commands.Environment, r io.Reader, prefix string, lineLimit, suffixLen int, numeric, verbose bool, separator string) int {
	scanner := bufio.NewScanner(r)
	fileIndex := 0
	lineCount := 0
	var currentWriter io.WriteCloser

	sep := "\n"
	if separator != "" {
		sep = separator
	}

	for scanner.Scan() {
		if currentWriter == nil || lineCount >= lineLimit {
			if currentWriter != nil {
				currentWriter.Close()
			}
			suffix := s.getSuffix(fileIndex, suffixLen, numeric)
			fileName := prefix + suffix
			fullPath := fileName
			if !filepath.IsAbs(fullPath) {
				fullPath = filepath.Join(env.Cwd, fullPath)
			}
			if verbose {
				fmt.Fprintf(env.Stdout, "creating file '%s'\n", fileName)
			}
			f, err := env.FS.Create(fullPath)
			if err != nil {
				fmt.Fprintf(env.Stderr, "split: %v\n", err)
				return 1
			}
			currentWriter = f
			fileIndex++
			lineCount = 0
		}
		fmt.Fprint(currentWriter, scanner.Text()+sep)
		lineCount++
	}

	if currentWriter != nil {
		currentWriter.Close()
	}
	return 0
}

func (s *Split) splitBytes(env *commands.Environment, r io.Reader, prefix string, byteLimit int64, suffixLen int, numeric, verbose bool) int {
	fileIndex := 0
	buf := make([]byte, 32*1024)
	var currentWriter io.WriteCloser
	var currentBytes int64

	for {
		n, err := r.Read(buf)
		if n > 0 {
			toRead := int64(n)
			start := 0
			for toRead > 0 {
				if currentWriter == nil || currentBytes >= byteLimit {
					if currentWriter != nil {
						currentWriter.Close()
					}
					suffix := s.getSuffix(fileIndex, suffixLen, numeric)
					fileName := prefix + suffix
					fullPath := fileName
					if !filepath.IsAbs(fullPath) {
						fullPath = filepath.Join(env.Cwd, fullPath)
					}
					if verbose {
						fmt.Fprintf(env.Stdout, "creating file '%s'\n", fileName)
					}
					f, err := env.FS.Create(fullPath)
					if err != nil {
						fmt.Fprintf(env.Stderr, "split: %v\n", err)
						return 1
					}
					currentWriter = f
					fileIndex++
					currentBytes = 0
				}

				canWrite := byteLimit - currentBytes
				if canWrite > toRead {
					canWrite = toRead
				}
				currentWriter.Write(buf[start : start+int(canWrite)])
				currentBytes += canWrite
				toRead -= canWrite
				start += int(canWrite)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(env.Stderr, "split: %v\n", err)
			return 1
		}
	}

	if currentWriter != nil {
		currentWriter.Close()
	}
	return 0
}

func (s *Split) getSuffix(index, length int, numeric bool) string {
	if numeric {
		return fmt.Sprintf("%0*d", length, index)
	}
	
	res := make([]byte, length)
	temp := index
	for i := length - 1; i >= 0; i-- {
		res[i] = byte('a' + (temp % 26))
		temp /= 26
	}
	return string(res)
}
