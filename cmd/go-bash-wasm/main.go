//go:build !js

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/yarencheng/go-bash-wasm/internal/app"
)

func main() {
	ctx := context.Background()

	// Initialize the application with standard input, output and error.
	a := app.New(os.Stdin, os.Stdout, os.Stderr)

	// Check for --version flag
	for _, arg := range os.Args {
		if arg == "--version" {
			a.ShowVersion()
			os.Exit(0)
		}
	}

	// Run the interactive shell.
	if err := a.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "bash simulator failed: %v\n", err)
		os.Exit(1)
	}
}
