//go:build !js

package app

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// newLoggerWriter returns the default logger writer for non-JS platforms.
func newLoggerWriter() io.Writer {
	return zerolog.ConsoleWriter{Out: os.Stderr}
}
