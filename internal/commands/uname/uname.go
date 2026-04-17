package uname

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Uname struct{}

func New() *Uname {
	return &Uname{}
}

func (u *Uname) Name() string {
	return "uname"
}

func (u *Uname) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("uname", pflag.ContinueOnError)
	kernelName := flags.BoolP("kernel-name", "s", false, "print the kernel name")
	nodeName := flags.BoolP("nodename", "n", false, "print the network node hostname")
	kernelRelease := flags.BoolP("kernel-release", "r", false, "print the kernel release")
	kernelVersion := flags.BoolP("kernel-version", "v", false, "print the kernel version")
	machine := flags.BoolP("machine", "m", false, "print the machine hardware name")
	processor := flags.BoolP("processor", "p", false, "print the processor type (non-portable)")
	hardwarePlatform := flags.BoolP("hardware-platform", "i", false, "print the hardware platform (non-portable)")
	operatingSystem := flags.BoolP("operating-system", "o", false, "print the operating system")
	all := flags.BoolP("all", "a", false, "print all information")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "uname: %v\n", err)
		return 1
	}

	if *all || (!*kernelName && !*nodeName && !*kernelRelease && !*kernelVersion && !*machine && !*processor && !*hardwarePlatform && !*operatingSystem) {
		*kernelName = true
		*nodeName = true
		*kernelRelease = true
		*kernelVersion = true
		*machine = true
		*operatingSystem = true
	}

	var results []string
	if *kernelName {
		results = append(results, "BashWasm")
	}
	if *nodeName {
		results = append(results, "wasm-host")
	}
	if *kernelRelease {
		results = append(results, "0.0.1")
	}
	if *kernelVersion {
		results = append(results, "#1 SMP WebAssembly")
	}
	if *machine {
		results = append(results, runtime.GOARCH)
	}
	if *processor {
		results = append(results, "unknown")
	}
	if *hardwarePlatform {
		results = append(results, "unknown")
	}
	if *operatingSystem {
		results = append(results, "GNU/Wasm")
	}

	fmt.Fprintln(env.Stdout, strings.Join(results, " "))
	return 0
}
