//go:build js

package app

import (
	"io"

	"github.com/yarencheng/go-bash-wasm/internal/wasm"
)

// newLoggerWriter returns the browser console logger writer for JS/WASM platforms.
func newLoggerWriter() io.Writer {
	return wasm.NewBrowserConsoleWriter()
}
