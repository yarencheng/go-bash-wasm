package ls

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Ls struct{}

func New() *Ls {
	return &Ls{}
}

func (l *Ls) Name() string {
	return "ls"
}

type lsFlags struct {
	all          *bool
	almostAll    *bool
	long         *bool
	human        *bool
	classify     *bool
	sortSize     *bool
	sortTime     *bool
	reverse      *bool
	inode        *bool
	recursive    *bool
	oneLine      *bool
	numeric      *bool
	dirIndicator *bool
	comma        *bool
	ctime        *bool
	atime        *bool
	noGroup      *bool
	sortOpt      *string
}

func (l *Ls) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flagsSet := pflag.NewFlagSet("ls", pflag.ContinueOnError)
	flagsSet.SetOutput(env.Stderr)

	f := lsFlags{
		all:          flagsSet.BoolP("all", "a", false, "do not ignore entries starting with ."),
		almostAll:    flagsSet.BoolP("almost-all", "A", false, "do not list implied . and .."),
		long:         flagsSet.BoolP("long", "l", false, "use a long listing format"),
		human:        flagsSet.BoolP("human-readable", "h", false, "with -l, print sizes like 1K 234M 2G etc."),
		classify:     flagsSet.BoolP("classify", "F", false, "append indicator (one of */=>@|) to entries"),
		sortSize:     flagsSet.BoolP("sort-size", "S", false, "sort by file size, largest first"),
		sortTime:     flagsSet.BoolP("sort-time", "t", false, "sort by modification time, newest first"),
		reverse:      flagsSet.BoolP("reverse", "r", false, "reverse order while sorting"),
		inode:        flagsSet.BoolP("inode", "i", false, "print the index number of each file"),
		recursive:    flagsSet.BoolP("recursive", "R", false, "list subdirectories recursively"),
		oneLine:      flagsSet.BoolP("format-1", "1", false, "list one file per line"),
		numeric:      flagsSet.BoolP("numeric-uid-gid", "n", false, "list numeric user and group IDs"),
		dirIndicator: flagsSet.BoolP("directory-indicator", "p", false, "append / indicator to directories"),
		comma:        flagsSet.BoolP("comma", "m", false, "fill width with a comma separated list of entries"),
		ctime:        flagsSet.BoolP("ctime", "c", false, "with -lt: sort by, and show, ctime; with -l: show ctime and sort by name"),
		atime:        flagsSet.BoolP("atime", "u", false, "with -lt: sort by, and show, atime; with -l: show atime and sort by name"),
		noGroup:      flagsSet.BoolP("no-group", "G", false, "in a long listing, don't print group names"),
		sortOpt:      flagsSet.String("sort", "", "sort by WORD: none (-U), size (-S), time (-t), version (-v), extension (-X)"),
	}

	if err := flagsSet.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "ls: %v\n", err)
		return 2
	}

	targets := flagsSet.Args()
	if len(targets) == 0 {
		targets = []string{env.Cwd}
	}

	exitCode := 0
	for i, target := range targets {
		if (len(targets) > 1 || *f.recursive) && i > 0 {
			fmt.Fprintln(env.Stdout)
		}
		if len(targets) > 1 || *f.recursive {
			fmt.Fprintf(env.Stdout, "%s:\n", target)
		}
		if res := l.listDir(ctx, env, target, &f, true); res != 0 {
			exitCode = res
		}
	}

	return exitCode
}

func (l *Ls) listDir(ctx context.Context, env *commands.Environment, target string, f *lsFlags, firstLevel bool) int {
	entries, err := afero.ReadDir(env.FS, target)
	if err != nil {
		fmt.Fprintf(env.Stderr, "ls: cannot access '%s': %v\n", target, err)
		return 2
	}

	// Determine sort mode
	sortMode := "name"
	if *f.sortSize || *f.sortOpt == "size" {
		sortMode = "size"
	} else if *f.sortTime || *f.sortOpt == "time" {
		sortMode = "time"
	} else if *f.ctime || *f.sortOpt == "ctime" {
		sortMode = "ctime"
	} else if *f.atime || *f.sortOpt == "atime" {
		sortMode = "atime"
	} else if *f.sortOpt == "none" {
		sortMode = "none"
	}

	// Sort entries
	if sortMode != "none" {
		sort.Slice(entries, func(i, j int) bool {
			var cmp bool
			switch sortMode {
			case "size":
				cmp = entries[i].Size() > entries[j].Size()
			case "time":
				cmp = entries[i].ModTime().After(entries[j].ModTime())
			case "ctime", "atime":
				// Afero MemMapFs doesn't distinguish well, so we use ModTime for now but extensible
				cmp = entries[i].ModTime().After(entries[j].ModTime())
			default:
				cmp = entries[i].Name() < entries[j].Name()
			}
			return cmp
		})
	}

	if *f.reverse {
		for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
			entries[i], entries[j] = entries[j], entries[i]
		}
	}

	var subDirs []string
	var results []string

	if *f.all {
		results = append(results, l.formatName(".", target, f))
		results = append(results, l.formatName("..", target, f))
	}

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, ".") && !*f.all && !*f.almostAll {
			continue
		}
		if *f.recursive && entry.IsDir() {
			subDirs = append(subDirs, target+"/"+name)
		}
		results = append(results, l.formatEntry(entry, target, f))
	}

	sep := "  "
	if *f.oneLine || *f.long || *f.numeric {
		sep = "\n"
	} else if *f.comma {
		sep = ", "
	}

	if len(results) > 0 {
		fmt.Fprint(env.Stdout, strings.Join(results, sep))
		if !*f.comma || *f.long || *f.oneLine {
			fmt.Fprintln(env.Stdout)
		} else {
			fmt.Fprintln(env.Stdout)
		}
	}

	if *f.recursive {
		for _, subDir := range subDirs {
			fmt.Fprintf(env.Stdout, "\n%s:\n", subDir)
			l.listDir(ctx, env, subDir, f, false)
		}
	}

	return 0
}

func (l *Ls) formatName(name string, target string, f *lsFlags) string {
	// For implied . and ..
	if *f.inode {
		return fmt.Sprintf("? %s", name)
	}
	return name
}

func (l *Ls) formatEntry(entry os.FileInfo, target string, f *lsFlags) string {
	name := entry.Name()
	if *f.classify {
		if entry.IsDir() {
			name += "/"
		} else if entry.Mode()&0111 != 0 {
			name += "*"
		}
	} else if *f.dirIndicator && entry.IsDir() {
		name += "/"
	}

	prefix := ""
	if *f.inode {
		prefix = "0 "
	}

	if *f.long || *f.numeric {
		sizeStr := fmt.Sprintf("%10d", entry.Size())
		if *f.human {
			sizeStr = fmt.Sprintf("%10s", formatHuman(entry.Size()))
		}
		owner := "root"
		group := "  root"
		if *f.numeric {
			owner = "0"
			group = "  0"
		}
		if *f.noGroup {
			group = ""
		}
		return fmt.Sprintf("%s%s  %s%s  %s  %s", prefix, entry.Mode().String(), owner, group, sizeStr, name)
	}

	return prefix + name
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
