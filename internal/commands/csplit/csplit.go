package csplit

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Csplit struct{}

func New() *Csplit {
	return &Csplit{}
}

func (c *Csplit) Name() string {
	return "csplit"
}

func (c *Csplit) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) < 2 {
		fmt.Fprintf(env.Stderr, "csplit: missing operand\n")
		return 1
	}

	inputPath := args[0]
	patterns := args[1:]

	f, err := env.FS.Open(c.absPath(env, inputPath))
	if err != nil {
		fmt.Fprintf(env.Stderr, "csplit: %v\n", err)
		return 1
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	prefix := "xx"
	fileIdx := 0
	currentLine := 1

	for _, pattern := range patterns {
		splitLine, err := strconv.Atoi(pattern)
		if err != nil {
			// Regex not supported yet
			continue
		}

		if splitLine <= currentLine {
			continue
		}

		filename := fmt.Sprintf("%s%02d", prefix, fileIdx)
		err = c.writeLines(env.FS, c.absPath(env, filename), lines[currentLine-1:splitLine-1])
		if err != nil {
			fmt.Fprintf(env.Stderr, "csplit: %v\n", err)
			return 1
		}

		fileIdx++
		currentLine = splitLine
	}

	// Write remaining lines
	filename := fmt.Sprintf("%s%02d", prefix, fileIdx)
	err = c.writeLines(env.FS, c.absPath(env, filename), lines[currentLine-1:])
	if err != nil {
		fmt.Fprintf(env.Stderr, "csplit: %v\n", err)
		return 1
	}

	return 0
}

func (c *Csplit) writeLines(fs afero.Fs, path string, lines []string) error {
	f, err := fs.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range lines {
		fmt.Fprintln(f, line)
	}
	return nil
}

func (c *Csplit) absPath(env *commands.Environment, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(env.Cwd, path)
}
