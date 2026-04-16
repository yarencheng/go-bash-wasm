package ls

import (
	"context"
	"fmt"
	"os"
	"path"
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
	directory    *bool
	doNotSort    *bool
	noOwner      *bool
	noGroupLong  *bool
	unsorted     *bool
	versionSort  *bool
	hide         *string
	ignore       *string
	indicatorStyle *string
	fileType       *bool
	sortExt        *bool
	sizeBlocks     *bool
	ignoreBackups  *bool
	siUnits        *bool
	kibibytes      *bool
	quotingStyle   *string
	escape         *bool
	quoteName      *bool
	hideControl    *bool
	sortOpt        *string
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
		directory:    flagsSet.BoolP("directory", "d", false, "list directories themselves, not their contents"),
		doNotSort:    flagsSet.BoolP("do-not-sort", "f", false, "do not sort, enable -aU, disable -ls --color"),
		noOwner:      flagsSet.BoolP("no-owner", "g", false, "like -l, but do not list owner"),
		noGroupLong:  flagsSet.BoolP("no-group-long", "o", false, "like -l, but do not list group"),
		unsorted:     flagsSet.BoolP("unsorted", "U", false, "do not sort; list entries in directory order"),
		versionSort:  flagsSet.BoolP("version-sort", "v", false, "natural sort of (version) numbers"),
		hide:         flagsSet.String("hide", "", "do not list implied entries matching shell PATTERN (overridden by -a or -A)"),
		ignore:       flagsSet.StringP("ignore", "I", "", "do not list implied entries matching shell PATTERN"),
		indicatorStyle: flagsSet.String("indicator-style", "none", "append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type), classify (-F)"),
		fileType:     flagsSet.Bool("file-type", false, "likewise, except do not append '*'"),
		sortExt:      flagsSet.BoolP("sort-extension", "X", false, "sort alphabetically by entry extension"),
		sizeBlocks:   flagsSet.BoolP("size", "s", false, "print the allocated size of each file, in blocks"),
		ignoreBackups: flagsSet.BoolP("ignore-backups", "B", false, "do not list implied entries ending with ~"),
		siUnits:      flagsSet.Bool("si", false, "likewise, but use powers of 1000 not 1024"),
		kibibytes:    flagsSet.BoolP("kibibytes", "k", false, "default to 1024-byte blocks for disk usage"),
		quotingStyle: flagsSet.String("quoting-style", "literal", "use quoting style WORD for entry names: literal, shell, shell-always, shell-escape, shell-escape-always, c, escape"),
		escape:       flagsSet.BoolP("escape", "b", false, "print C-style escapes for nongraphic characters"),
		quoteName:    flagsSet.BoolP("quote-name", "Q", false, "enclose entry names in double quotes"),
		hideControl:  flagsSet.BoolP("hide-control-chars", "q", false, "print ? instead of nongraphic characters"),
		sortOpt:      flagsSet.String("sort", "", "sort by WORD: none (-U), size (-S), time (-t), version (-v), extension (-X)"),
	}

	if err := flagsSet.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "ls: %v\n", err)
		return 2
	}

	if *f.doNotSort {
		*f.all = true
		*f.unsorted = true
		// disable -l -s if they were set? The spec says "disable -ls".
		// But usually it just means it changes the default or behaves as if they are off unless forced?
		// Actually GNU ls -f DOES disable -l -s --color.
		*f.long = false
	}

	targets := flagsSet.Args()
	if len(targets) == 0 {
		targets = []string{"."}
	}

	if *f.directory {
		for i, target := range targets {
			if i > 0 {
				fmt.Fprint(env.Stdout, "  ")
			}
			fullPath := target
			if !strings.HasPrefix(target, "/") {
				fullPath = env.Cwd + "/" + target
			}
			info, err := env.FS.Stat(fullPath)
			if err != nil {
				fmt.Fprintf(env.Stderr, "ls: cannot access '%s': %v\n", target, err)
				continue
			}
			fmt.Fprint(env.Stdout, l.formatEntry(info, target, target, &f))
		}
		fmt.Fprintln(env.Stdout)
		return 0
	}

	exitCode := 0
	for i, target := range targets {
		displayTarget := target
		fullPath := target
		if target == "." {
			displayTarget = env.Cwd
			fullPath = env.Cwd
		} else if !strings.HasPrefix(target, "/") {
			fullPath = env.Cwd + "/" + target
		}

		if (len(targets) > 1 || *f.recursive) && i > 0 {
			fmt.Fprintln(env.Stdout)
		}
		if len(targets) > 1 || *f.recursive {
			fmt.Fprintf(env.Stdout, "%s:\n", displayTarget)
		}
		if res := l.listDir(ctx, env, fullPath, &f, true); res != 0 {
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
	if *f.unsorted || *f.sortOpt == "none" {
		sortMode = "none"
	} else if *f.sortSize || *f.sortOpt == "size" {
		sortMode = "size"
	} else if *f.sortTime || *f.sortOpt == "time" {
		sortMode = "time"
	} else if *f.ctime || *f.sortOpt == "ctime" {
		sortMode = "ctime"
	} else if *f.atime || *f.sortOpt == "atime" {
		sortMode = "atime"
	} else if *f.versionSort || *f.sortOpt == "version" || *f.sortOpt == "v" {
		sortMode = "version"
	} else if *f.sortExt || *f.sortOpt == "extension" || *f.sortOpt == "X" {
		sortMode = "extension"
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
			case "version":
				cmp = naturalLess(entries[i].Name(), entries[j].Name())
			case "extension":
				extI := path.Ext(entries[i].Name())
				extJ := path.Ext(entries[j].Name())
				if extI != extJ {
					cmp = extI < extJ
				} else {
					cmp = entries[i].Name() < entries[j].Name()
				}
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
		results = append(results, l.formatEntry(nil, ".", target, f))
		results = append(results, l.formatEntry(nil, "..", target, f))
	}

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, ".") && !*f.all && !*f.almostAll {
			continue
		}

		if *f.ignore != "" {
			matched, _ := path.Match(*f.ignore, name)
			if matched {
				continue
			}
		}

		if *f.hide != "" && !*f.all && !*f.almostAll {
			matched, _ := path.Match(*f.hide, name)
			if matched {
				continue
			}
		}

		if *f.ignoreBackups && strings.HasSuffix(name, "~") {
			continue
		}

		if *f.recursive && entry.IsDir() {
			subDirs = append(subDirs, target+"/"+name)
		}
		results = append(results, l.formatEntry(entry, name, target, f))
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

func (l *Ls) formatEntry(entry os.FileInfo, name string, target string, f *lsFlags) string {
	name = applyQuoting(name, f)
	style := *f.indicatorStyle
	if *f.classify {
		style = "classify"
	} else if *f.fileType {
		style = "file-type"
	} else if *f.dirIndicator {
		style = "slash"
	}

	switch style {
	case "slash":
		if entry != nil && entry.IsDir() {
			name += "/"
		} else if entry == nil && (name == "." || name == "..") {
			name += "/"
		}
	case "file-type":
		if entry != nil && entry.IsDir() {
			name += "/"
		} else if entry == nil && (name == "." || name == "..") {
			name += "/"
		} else if entry != nil && entry.Mode()&os.ModeSymlink != 0 {
			name += "@"
		} else if entry != nil && entry.Mode()&os.ModeSocket != 0 {
			name += "="
		} else if entry != nil && entry.Mode()&os.ModeNamedPipe != 0 {
			name += "|"
		}
	case "classify":
		if entry != nil && entry.IsDir() {
			name += "/"
		} else if entry == nil && (name == "." || name == "..") {
			name += "/"
		} else if entry != nil && entry.Mode()&os.ModeSymlink != 0 {
			name += "@"
		} else if entry != nil && entry.Mode()&os.ModeSocket != 0 {
			name += "="
		} else if entry != nil && entry.Mode()&os.ModeNamedPipe != 0 {
			name += "|"
		} else if entry != nil && entry.Mode()&0111 != 0 {
			name += "*"
		}
	}

	prefix := ""
	if *f.inode {
		prefix = "0 "
	}

	if *f.sizeBlocks {
		size := int64(0)
		if entry != nil {
			size = entry.Size()
		}
		blocks := (size + 1023) / 1024
		prefix += fmt.Sprintf("%d ", blocks)
	}

	if *f.long || *f.numeric || *f.noOwner || *f.noGroupLong {
		size := int64(0)
		mode := os.FileMode(0)
		if entry != nil {
			size = entry.Size()
			mode = entry.Mode()
		} else if name == "." || name == ".." {
			mode = os.ModeDir | 0755
		}

		sizeStr := fmt.Sprintf("%10d", size)
		if *f.human || *f.siUnits {
			sizeStr = fmt.Sprintf("%10s", formatHuman(size, *f.siUnits))
		}
		owner := "root"
		group := "  root"
		if *f.numeric {
			owner = "0"
			group = "  0"
		}
		if *f.noOwner {
			owner = ""
		}
		if *f.noGroup || *f.noGroupLong {
			group = ""
		}

		fields := []string{mode.String()}
		if owner != "" {
			fields = append(fields, owner)
		}
		if group != "" {
			fields = append(fields, group)
		}
		fields = append(fields, sizeStr, name)

		return prefix + strings.Join(fields, "  ")
	}
	prefix += name
	return prefix
}

func applyQuoting(name string, f *lsFlags) string {
	style := *f.quotingStyle
	if *f.quoteName {
		style = "c"
	} else if *f.escape {
		style = "escape"
	}

	switch style {
	case "c":
		return fmt.Sprintf("%q", name)
	case "escape":
		s := fmt.Sprintf("%q", name)
		return s[1 : len(s)-1]
	case "shell":
		if strings.ContainsAny(name, " \t\n'\"") {
			return fmt.Sprintf("'%s'", strings.ReplaceAll(name, "'", "'\\''"))
		}
		return name
	default:
		if *f.hideControl {
			return strings.Map(func(r rune) rune {
				if r < 32 || r == 127 {
					return '?'
				}
				return r
			}, name)
		}
		return name
	}
}

func formatHuman(size int64, si bool) string {
	unit := int64(1024)
	if si {
		unit = 1000
	}
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

func naturalLess(s1, s2 string) bool {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		c1, c2 := s1[i], s2[j]
		if isDigit(c1) && isDigit(c2) {
			n1, nextI := parseDigits(s1, i)
			n2, nextJ := parseDigits(s2, j)
			if n1 != n2 {
				return n1 < n2
			}
			i, j = nextI, nextJ
		} else {
			if c1 != c2 {
				return c1 < c2
			}
			i++
			j++
		}
	}
	return len(s1) < len(s2)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseDigits(s string, start int) (int, int) {
	val := 0
	i := start
	for i < len(s) && isDigit(s[i]) {
		val = val*10 + int(s[i]-'0')
		i++
	}
	return val, i
}
