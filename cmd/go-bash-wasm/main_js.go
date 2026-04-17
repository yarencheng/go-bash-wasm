//go:build js

package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"syscall/js"

	"github.com/rs/zerolog"
	"github.com/yarencheng/go-bash-wasm/internal/app"
)

func main() {
	// Suppress zerolog noise from appearing in the browser terminal.
	zerolog.SetGlobalLevel(zerolog.Disabled)

	ctx := context.Background()

	// Create a synchronous pipe. The JS side writes lines in; the Go
	// shell reads them out via bufio.Scanner (wasm goroutine-safe).
	pr, pw := io.Pipe()

	// Expose writeStdin(line string) to JS so xterm.js can feed input.
	writeStdinFn := js.FuncOf(func(_ js.Value, args []js.Value) any {
		if len(args) == 0 {
			return nil
		}
		data := args[0].String()
		// Write in a goroutine so we don't block the JS callback.
		go func() {
			if _, err := pw.Write([]byte(data)); err != nil {
				fmt.Fprintf(os.Stderr, "writeStdin: %v\n", err)
			}
		}()
		return nil
	})
	js.Global().Set("writeStdin", writeStdinFn)
	defer writeStdinFn.Release()

	// Expose closeStdin() to JS so Ctrl+D can signal EOF.
	closeStdinFn := js.FuncOf(func(_ js.Value, _ []js.Value) any {
		pw.Close()
		return nil
	})
	js.Global().Set("closeStdin", closeStdinFn)
	defer closeStdinFn.Release()

	a := app.New(pr, os.Stdout, os.Stderr)
	if err := a.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "bash simulator failed: %v\n", err)
	}

	// Keep the WASM module alive after the shell exits so the page
	// doesn't throw "Go program has already exited" on callbacks.
	select {}
}
