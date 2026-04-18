package ls

import (
	"context"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"time"

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
	all            *bool
	almostAll      *bool
	long           *bool
	human          *bool
	classify       *bool
	sortSize       *bool
	sortTime       *bool
	reverse        *bool
	inode          *bool
	recursive      *bool
	oneLine        *bool
	numeric        *bool
	dirIndicator   *bool
	comma          *bool
	ctime          *bool
	atime          *bool
	noGroup        *bool
	directory      *bool
	doNotSort      *bool
	noOwner        *bool
	noGroupLong    *bool
	unsorted       *bool
	versionSort    *bool
	hide           *string
	ignore         *string
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
	groupDirsFirst *bool
	zero           *bool
	dereference    *bool
	dereferenceArg *bool
	format         *string
	timeStyle      *string
	color          *string
	fullTime       *bool
	vertical       *bool
	timeOpt        *string
	blockSize      *string
	sortOpt        *string
	tabSize        *int
	width          *int
	help           *bool
	version        *bool
	context        *bool
}

func (l *Ls) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flagsSet := pflag.NewFlagSet("ls", pflag.ContinueOnError)
	flagsSet.SetOutput(env.Stderr)

	f := lsFlags{
		all:            flagsSet.BoolP("all", "a", false, "do not ignore entries starting with ."),
		almostAll:      flagsSet.BoolP("almost-all", "A", false, "do not list implied . and .."),
		long:           flagsSet.BoolP("long", "l", false, "use a long listing format"),
		human:          flagsSet.BoolP("human-readable", "h", false, "with -l, print sizes like 1K 234M 2G etc."),
		classify:       flagsSet.BoolP("classify", "F", false, "append indicator (one of */=>@|) to entries"),
		sortSize:       flagsSet.BoolP("sort-size", "S", false, "sort by file size, largest first"),
		sortTime:       flagsSet.BoolP("sort-time", "t", false, "sort by modification time, newest first"),
		reverse:        flagsSet.BoolP("reverse", "r", false, "reverse order while sorting"),
		inode:          flagsSet.BoolP("inode", "i", false, "print the index number of each file"),
		recursive:      flagsSet.BoolP("recursive", "R", false, "list subdirectories recursively"),
		oneLine:        flagsSet.BoolP("format-1", "1", false, "list one file per line"),
		numeric:        flagsSet.BoolP("numeric-uid-gid", "n", false, "list numeric user and group IDs"),
		dirIndicator:   flagsSet.BoolP("directory-indicator", "p", false, "append / indicator to directories"),
		comma:          flagsSet.BoolP("comma", "m", false, "fill width with a comma separated list of entries"),
		ctime:          flagsSet.BoolP("ctime", "c", false, "with -lt: sort by, and show, ctime; with -l: show ctime and sort by name"),
		atime:          flagsSet.BoolP("atime", "u", false, "with -lt: sort by, and show, atime; with -l: show atime and sort by name"),
		noGroup:        flagsSet.BoolP("no-group", "G", false, "in a long listing, don't print group names"),
		directory:      flagsSet.BoolP("directory", "d", false, "list directories themselves, not their contents"),
		doNotSort:      flagsSet.BoolP("do-not-sort", "f", false, "do not sort, enable -aU, disable -ls --color"),
		noOwner:        flagsSet.BoolP("no-owner", "g", false, "like -l, but do not list owner"),
		noGroupLong:    flagsSet.BoolP("no-group-long", "o", false, "like -l, but do not list group"),
		unsorted:       flagsSet.BoolP("unsorted", "U", false, "do not sort; list entries in directory order"),
		versionSort:    flagsSet.BoolP("version-sort", "v", false, "natural sort of (version) numbers"),
		hide:           flagsSet.String("hide", "", "do not list implied entries matching shell PATTERN (overridden by -a or -A)"),
		ignore:         flagsSet.StringP("ignore", "I", "", "do not list implied entries matching shell PATTERN"),
		indicatorStyle: flagsSet.String("indicator-style", "none", "append indicator with style WORD to entry names: none (default), slash (-p), file-type (--file-type), classify (-F)"),
		fileType:       flagsSet.Bool("file-type", false, "likewise, except do not append '*'"),
		sortExt:        flagsSet.BoolP("sort-extension", "X", false, "sort alphabetically by entry extension"),
		sizeBlocks:     flagsSet.BoolP("size", "s", false, "print the allocated size of each file, in blocks"),
		ignoreBackups:  flagsSet.BoolP("ignore-backups", "B", false, "do not list implied entries ending with ~"),
		siUnits:        flagsSet.Bool("si", false, "likewise, but use powers of 1000 not 1024"),
		kibibytes:      flagsSet.BoolP("kibibytes", "k", false, "default to 1024-byte blocks for disk usage"),
		quotingStyle:   flagsSet.String("quoting-style", "literal", "use quoting style WORD for entry names: literal, shell, shell-always, shell-escape, shell-escape-always, c, escape"),
		escape:         flagsSet.BoolP("escape", "b", false, "print C-style escapes for nongraphic characters"),
		quoteName:      flagsSet.BoolP("quote-name", "Q", false, "enclose entry names in double quotes"),
		hideControl:    flagsSet.BoolP("hide-control-chars", "q", false, "print ? instead of nongraphic characters"),
		groupDirsFirst: flagsSet.Bool("group-directories-first", false, "group directories before files"),
		zero:           flagsSet.Bool("zero", false, "end each output line with NUL, not newline"),
		dereference:    flagsSet.BoolP("dereference", "L", false, "when showing file information for a symbolic link, show information for the file the link references rather than for the link itself"),
		dereferenceArg: flagsSet.BoolP("dereference-command-line", "H", false, "follow symbolic links listed on the command line"),
		format:         flagsSet.String("format", "vertical", "across (-x), commas (-m), horizontal (-x), long (-l), single-column (-1), verbose (-l), vertical (-C)"),
		timeStyle:      flagsSet.String("time-style", "locale", "time/date format with -l: full-iso, long-iso, iso, locale"),
		color:          flagsSet.String("color", "never", "colorize the output; WHEN can be 'always' (default if omitted), 'auto', or 'never'"),
		fullTime:       flagsSet.Bool("full-time", false, "like -l --time-style=full-iso"),
		timeOpt:        flagsSet.String("time", "", "show time as WORD instead of modification time: atime, access, use, ctime, status"),
		blockSize:      flagsSet.String("block-size", "", "scale sizes by SIZE when printing them"),
		vertical:       flagsSet.BoolP("vertical", "C", false, "list entries by columns"),
		sortOpt:        flagsSet.String("sort", "", "sort by WORD: none (-U), size (-S), time (-t), version (-v), extension (-X)"),
		tabSize:        flagsSet.IntP("tabsize", "T", 8, "assume tab stops at each COLS instead of 8"),
		width:          flagsSet.IntP("width", "w", 80, "assume screen width instead of current value"),
		help:           flagsSet.Bool("help", false, "display this help and exit"),
		version:        flagsSet.Bool("version", false, "output version information and exit"),
		context:        flagsSet.BoolP("context", "Z", false, "print any security context of each file"),
	}

	if err := flagsSet.Parse(args); err != nil {
		if err == pflag.ErrHelp {
			return 0
		}
		fmt.Fprintf(env.Stderr, "ls: %v\n", err)
		return 2
	}

	if *f.help {
		fmt.Fprintf(env.Stdout, "Usage: ls [OPTION]... [FILE]...\n")
		fmt.Fprintf(env.Stdout, "List information about the FILEs (the current directory by default).\n\n")
		flagsSet.PrintDefaults()
		return 0
	}

	if *f.version {
		commands.ShowVersion(env.Stdout, "ls")
		return 0
	}

	if *f.doNotSort {
		*f.all = true
		*f.unsorted = true
		// disable -l -s if they were set? The spec says "disable -ls".
		// But usually it just means it changes the default or behaves as if they are off unless forced?
		// Actually GNU ls -f DOES disable -l -s --color.
		*f.long = false
	}

	if *f.vertical {
		*f.format = "vertical"
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
			info, err := l.stat(env.FS, fullPath, *f.dereference || *f.dereferenceArg)
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
	var fileEntries []struct {
		info os.FileInfo
		name string
		path string
	}
	var dirTargets []struct {
		name string
		path string
	}

	for _, target := range targets {
		displayTarget := target
		fullPath := target
		if target == "." {
			displayTarget = env.Cwd
			fullPath = env.Cwd
		} else if !strings.HasPrefix(target, "/") {
			fullPath = env.Cwd + "/" + target
		}

		info, err := l.stat(env.FS, fullPath, *f.dereference || *f.dereferenceArg)
		if err != nil {
			fmt.Fprintf(env.Stderr, "ls: cannot access '%s': %v\n", target, err)
			exitCode = 2
			continue
		}

		if info.IsDir() && !*f.directory {
			dirTargets = append(dirTargets, struct {
				name string
				path string
			}{displayTarget, fullPath})
		} else {
			fileEntries = append(fileEntries, struct {
				info os.FileInfo
				name string
				path string
			}{info, displayTarget, fullPath})
		}
	}

	if len(fileEntries) > 0 {
		var results []string
		for _, fe := range fileEntries {
			results = append(results, l.formatEntry(fe.info, fe.name, fe.path, &f))
		}
		sep := "  "
		if *f.oneLine || *f.long || *f.numeric || *f.format == "long" || *f.format == "verbose" || *f.format == "single-column" {
			sep = "\n"
		}
		fmt.Fprint(env.Stdout, strings.Join(results, sep))
		fmt.Fprintln(env.Stdout)
	}

	for i, dt := range dirTargets {
		if (len(targets) > 1 || *f.recursive) && (i > 0 || len(fileEntries) > 0) {
			fmt.Fprintln(env.Stdout)
		}
		if len(targets) > 1 || *f.recursive {
			fmt.Fprintf(env.Stdout, "%s:\n", dt.name)
		}
		if res := l.listDir(ctx, env, dt.path, &f, true); res != 0 {
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
	if sortMode != "none" || *f.groupDirsFirst {
		sort.Slice(entries, func(i, j int) bool {
			if *f.groupDirsFirst {
				if entries[i].IsDir() && !entries[j].IsDir() {
					return true
				}
				if !entries[i].IsDir() && entries[j].IsDir() {
					return false
				}
			}
			if sortMode == "none" {
				return false // Should not happen with groupDirsFirst but safe
			}
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
	if *f.zero {
		sep = "\x00"
	} else if *f.oneLine || *f.long || *f.numeric || *f.format == "long" || *f.format == "verbose" || *f.format == "single-column" {
		sep = "\n"
	} else if *f.comma || *f.format == "commas" {
		sep = ", "
	}

	if len(results) > 0 {
		fmt.Fprint(env.Stdout, strings.Join(results, sep))
		if *f.zero {
			fmt.Fprint(env.Stdout, "\x00")
		} else if !(*f.comma || *f.format == "commas") || *f.long || *f.oneLine {
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
	rawName := name
	name = applyQuoting(name, f)
	if *f.color == "always" {
		name = applyColor(name, entry)
	}
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

	scale := int64(1024)
	if *f.blockSize != "" {
		fmt.Sscanf(*f.blockSize, "%d", &scale)
	}

	if *f.sizeBlocks {
		size := int64(0)
		if entry != nil {
			size = entry.Size()
		}
		blocks := (size + scale - 1) / scale
		prefix += fmt.Sprintf("%d ", blocks)
	}

	if *f.long || *f.numeric || *f.noOwner || *f.noGroupLong || *f.format == "long" || *f.format == "verbose" {
		size := int64(0)
		mode := os.FileMode(0)
		if entry != nil {
			size = entry.Size()
			mode = entry.Mode()
		} else if rawName == "." || rawName == ".." {
			mode = os.ModeDir | 0755
		}

		sizeStr := fmt.Sprintf("%10d", size)
		if *f.human || *f.siUnits {
			sizeStr = fmt.Sprintf("%10s", formatHuman(size, *f.siUnits))
		} else if *f.blockSize != "" {
			sizeStr = fmt.Sprintf("%10d", (size+scale-1)/scale)
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

		timeFormat := "Jan _2 15:04"
		if *f.fullTime || *f.timeStyle == "full-iso" {
			timeFormat = "2006-01-02 15:04:05.000000000 -0700"
		} else if *f.timeStyle == "long-iso" {
			timeFormat = "2006-01-02 15:04"
		} else if *f.timeStyle == "iso" {
			timeFormat = "01-02 15:04"
		}

		mtime := entry.ModTime().Format(timeFormat)
		if entry == nil && (rawName == "." || rawName == "..") {
			mtime = time.Now().Format(timeFormat)
		}

		fields = append(fields, sizeStr, mtime, name)

		return prefix + strings.Join(fields, "  ")
	}
	prefix += name
	return prefix
}

func applyColor(name string, entry os.FileInfo) string {
	if entry == nil {
		return "\033[1;34m" + name + "\033[0m" // Assume dir for . and ..
	}
	mode := entry.Mode()
	if mode.IsDir() {
		return "\033[1;34m" + name + "\033[0m"
	}
	if mode&os.ModeSymlink != 0 {
		return "\033[1;36m" + name + "\033[0m"
	}
	if mode&os.ModeSocket != 0 {
		return "\033[1;35m" + name + "\033[0m"
	}
	if mode&os.ModeNamedPipe != 0 {
		return "\033[33m" + name + "\033[0m"
	}
	if mode&0111 != 0 {
		return "\033[1;32m" + name + "\033[0m"
	}
	return name
}

func (l *Ls) stat(fs afero.Fs, path string, dereference bool) (os.FileInfo, error) {
	if dereference {
		return fs.Stat(path)
	}
	if lstater, ok := fs.(afero.Lstater); ok {
		info, _, err := lstater.LstatIfPossible(path)
		return info, err
	}
	return fs.Stat(path)
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
