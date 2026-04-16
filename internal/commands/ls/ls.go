package ls

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Ls struct {
}

func New() *Ls {
	return &Ls{}
}

func (l *Ls) Name() string {
	return "ls"
}

func (l *Ls) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("ls", pflag.ContinueOnError)
	flags.SetOutput(env.Stderr)

	all := flags.BoolP("all", "a", false, "do not ignore entries starting with .")
	almostAll := flags.BoolP("almost-all", "A", false, "do not list implied . and ..")
	long := flags.BoolP("long", "l", false, "use a long listing format")
	human := flags.BoolP("human-readable", "h", false, "with -l, print sizes like 1K 234M 2G etc.")
	classify := flags.BoolP("classify", "F", false, "append indicator (one of */=>@|) to entries")
	sortSize := flags.BoolP("sort-size", "S", false, "sort by file size, largest first")
	sortTime := flags.BoolP("sort-time", "t", false, "sort by modification time, newest first")
	reverse := flags.BoolP("reverse", "r", false, "reverse order while sorting")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "ls: %v\n", err)
		return 2
	}

	targets := flags.Args()
	if len(targets) == 0 {
		targets = []string{env.Cwd}
	}

	// For simplicity, we'll handle multiple targets sequentially for now
	for _, target := range targets {
		entries, err := afero.ReadDir(env.FS, target)
		if err != nil {
			fmt.Fprintf(env.Stderr, "ls: cannot access '%s': %v\n", target, err)
			return 2
		}

		// Sort entries
		if *sortSize {
			sort.Slice(entries, func(i, j int) bool {
				return entries[i].Size() > entries[j].Size()
			})
		} else if *sortTime {
			sort.Slice(entries, func(i, j int) bool {
				return entries[i].ModTime().After(entries[j].ModTime())
			})
		} else {
			sort.Slice(entries, func(i, j int) bool {
				return entries[i].Name() < entries[j].Name()
			})
		}

		if *reverse {
			for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}

		var names []string
		if *all {
			names = append(names, ".", "..")
		}

		for _, entry := range entries {
			name := entry.Name()
			if strings.HasPrefix(name, ".") && !*all && !*almostAll {
				continue
			}
			if *classify {
				if entry.IsDir() {
					name += "/"
				} else if entry.Mode()&0111 != 0 {
					name += "*"
				}
			}
			names = append(names, name)
		}

		if len(names) > 0 {
			if *long {
				for _, name := range names {
					info, err := env.FS.Stat(target + "/" + name)
					if name == "." || name == ".." {
						info, err = env.FS.Stat(target)
					}
					if err != nil {
						fmt.Fprintf(env.Stdout, "?  %s\n", name)
						continue
					}
					sizeStr := fmt.Sprintf("%d", info.Size())
					if *human {
						sizeStr = formatHuman(info.Size())
					}
					fmt.Fprintf(env.Stdout, "%s  %s  %s\n", info.Mode().String(), sizeStr, name)
				}
			} else {
				fmt.Fprintln(env.Stdout, strings.Join(names, "  "))
			}
		}
	}

	return 0
}

func formatHuman(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(size)/float64(div), "KMGTPE"[exp])
}
