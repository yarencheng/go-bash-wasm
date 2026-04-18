//go:build js
package wasm

import (
	"encoding/json"
	"io"
	"syscall/js"
)

// BrowserConsoleWriter implements io.Writer and redirects output to the browser's console.
type BrowserConsoleWriter struct{}

// NewBrowserConsoleWriter creates a new BrowserConsoleWriter.
func NewBrowserConsoleWriter() io.Writer {
	return &BrowserConsoleWriter{}
}

func (w *BrowserConsoleWriter) Write(p []byte) (n int, err error) {
	console := js.Global().Get("console")
	if !console.Truthy() {
		return len(p), nil
	}

	// Try to parse as JSON to provide a better experience in the console
	var entry map[string]any
	if err := json.Unmarshal(p, &entry); err == nil {
		level, _ := entry["level"].(string)
		msg, _ := entry["message"].(string)

		// Map zerolog levels to browser console methods
		var args []any
		if msg != "" {
			args = append(args, msg)
		}
		// Include the full entry as an object so it's expandable in the console
		args = append(args, entry)

		jsArgs := make([]any, len(args))
		for i, arg := range args {
			jsArgs[i] = arg
		}

		switch level {
		case "debug", "trace":
			console.Call("debug", jsArgs...)
		case "info":
			console.Call("info", jsArgs...)
		case "warn":
			console.Call("warn", jsArgs...)
		case "error", "fatal", "panic":
			console.Call("error", jsArgs...)
		default:
			console.Call("log", jsArgs...)
		}
	} else {
		// Fallback to plain text if not JSON
		console.Call("log", string(p))
	}

	return len(p), nil
}
