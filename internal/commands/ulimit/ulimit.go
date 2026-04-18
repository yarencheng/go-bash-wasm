package ulimit

import (
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Ulimit struct{}

type resource struct {
	option      string
	description string
	units       string
	value       string
}

func New() *Ulimit {
	return &Ulimit{}
}

func (u *Ulimit) Name() string {
	return "ulimit"
}

func (u *Ulimit) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("ulimit", pflag.ContinueOnError)
	all := flags.BoolP("all", "a", false, "all current limits are reported")
	_ = flags.BoolP("hard", "H", false, "use the hard resource limit")
	_ = flags.BoolP("soft", "S", false, "use the soft resource limit")

	// Pre-define common limits
	limits := []resource{
		{"c", "core file size", "(blocks, -c)", "0"},
		{"d", "data seg size", "(kbytes, -d)", "unlimited"},
		{"e", "scheduling priority", "(-e)", "0"},
		{"f", "file size", "(blocks, -f)", "unlimited"},
		{"i", "pending signals", "(-i)", "7814"},
		{"l", "max locked memory", "(kbytes, -l)", "64"},
		{"m", "max memory size", "(kbytes, -m)", "unlimited"},
		{"n", "open files", "(-n)", "1024"},
		{"p", "pipe size", "(512 bytes, -p)", "8"},
		{"q", "POSIX message queues", "(bytes, -q)", "819200"},
		{"r", "real-time priority", "(-r)", "0"},
		{"s", "stack size", "(kbytes, -s)", "8192"},
		{"t", "cpu time", "(seconds, -t)", "unlimited"},
		{"u", "max user processes", "(-u)", "7814"},
		{"v", "virtual memory", "(kbytes, -v)", "unlimited"},
		{"x", "file locks", "(-x)", "unlimited"},
	}

	// Dynamic flags for each resource
	resourceFlags := make(map[string]*bool)
	for _, res := range limits {
		resourceFlags[res.option] = flags.BoolP(res.option, res.option, false, res.description)
	}

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "ulimit: %v\n", err)
		return 1
	}

	if *all {
		for _, res := range limits {
			fmt.Fprintf(env.Stdout, "%-25s %-15s %s\n", res.description, res.units, res.value)
		}
		return 0
	}

	// Check if any specific resource flag is set
	anySet := false
	for opt, isSet := range resourceFlags {
		if *isSet {
			anySet = true
			for _, res := range limits {
				if res.option == opt {
					fmt.Fprintf(env.Stdout, "%s\n", res.value)
					break
				}
			}
		}
	}

	if !anySet {
		// Default to file size
		if len(args) == 0 {
			fmt.Fprintf(env.Stdout, "unlimited\n")
		} else {
			// Setting limit is not supported in simulation
			fmt.Fprintf(env.Stderr, "ulimit: setting limit not supported in this environment\n")
			return 1
		}
	}

	return 0
}
