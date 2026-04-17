package hostname

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Hostname struct{}

func New() *Hostname {
	return &Hostname{}
}

func (h *Hostname) Name() string {
	return "hostname"
}

func (h *Hostname) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("hostname", pflag.ContinueOnError)
	help := flags.BoolP("help", "", false, "display this help and exit")
	version := flags.BoolP("version", "", false, "output version information and exit")

	if err := flags.Parse(args); err != nil {
		if env.Stderr != nil {
			fmt.Fprintf(env.Stderr, "hostname: %v\n", err)
		}
		return 1
	}

	if *help {
		fmt.Fprintln(env.Stdout, "Usage: hostname [NAME]")
		fmt.Fprintln(env.Stdout, "Print or set the system's host name.")
		return 0
	}

	if *version {
		fmt.Fprintln(env.Stdout, "hostname (go-bash-wasm) 1.0")
		return 0
	}

	remaining := flags.Args()
	if len(remaining) > 0 {
		env.EnvVars["HOSTNAME"] = remaining[0]
		return 0
	}

	hostname, ok := env.EnvVars["HOSTNAME"]
	if !ok {
		hostname = "wasm-host"
	}
	fmt.Fprintln(env.Stdout, hostname)
	return 0
}
