package commands

import (
	"fmt"
	"io"
)

// ShowVersion prints the version information for a given command.
func ShowVersion(stdout io.Writer, name string) {
	fmt.Fprintf(stdout, "%s (go-bash-wasm) %s\n", name, BashVersion)
	fmt.Fprintf(stdout, "%s\n", BashCopyright)
	fmt.Fprintf(stdout, "%s\n", BashLicense)
	fmt.Fprintf(stdout, "Written by go-bash-wasm team.\n")
}
