package app

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/alias"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/arch"
	base32cmd "github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/base32"
	base64cmd "github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/base64"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/basename"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/basenc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/bg"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/bind"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/boolcmd"
	breakcmd "github.com/yarencheng/go-bash-wasm/internal/commands/bash/break"
	builtincmd "github.com/yarencheng/go-bash-wasm/internal/commands/bash/builtin"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/caller"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/cat"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/cd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/chcon"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/chmod"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/chown"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/chroot"
	cksumcmd "github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/cksum"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/clear"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/colon"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/comm"
	commandcmd "github.com/yarencheng/go-bash-wasm/internal/commands/bash/command"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/compgen"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/complete"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/compopt"
	continuecmd "github.com/yarencheng/go-bash-wasm/internal/commands/bash/continue"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/cp"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/csplit"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/cut"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/date"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/dd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/declare"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/df"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/dir"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/dircolors"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/dirname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/dirs"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/disown"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/du"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/echo"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/enable"
	env "github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/env"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/eval"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/exec"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/exit"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/expand"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/export"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/expr"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/factor"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/fc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/fg"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/find"
	fmtcmd "github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/fmt"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/fold"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/getlimits"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/getopts"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/grep"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/groups"
	hashcmd "github.com/yarencheng/go-bash-wasm/internal/commands/bash/hash"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/head"
	helpcmd "github.com/yarencheng/go-bash-wasm/internal/commands/bash/help"
	historycmd "github.com/yarencheng/go-bash-wasm/internal/commands/bash/history"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/hostid"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/hostname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/id"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/install"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/jobs"
	join "github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/join"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/kill"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/letcmd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/link"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/ln"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/local"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/logname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/logout"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/ls"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/mapfile"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/mkdir"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/mkfifo"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/mknod"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/mktemp"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/mv"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/nice"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/nl"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/nohup"
	nproccmd "github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/nproc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/numfmt"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/od"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/paste"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/pathchk"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/pinky"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/popd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/pr"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/printenv"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/printf"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/ptx"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/pushd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/pwd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/read"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/readlink"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/readonly"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/realpath"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/returncmd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/rm"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/rmdir"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/runcon"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/seq"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/set"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/shift"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/shopt"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/shred"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/shuf"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/sleep"
	sort "github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/sort"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/source"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/split"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/stat"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/stdbuf"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/stty"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/sum"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/sumlegacy"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/suspend"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/sync"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/tac"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/tail"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/tee"
	test "github.com/yarencheng/go-bash-wasm/internal/commands/bash/test"
	timecmd "github.com/yarencheng/go-bash-wasm/internal/commands/bash/time"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/timeout"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/times"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/touch"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/tr"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/trap"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/truncate"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/tsort"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/tty"
	typecmd "github.com/yarencheng/go-bash-wasm/internal/commands/bash/type"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/ulimit"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/umask"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/unalias"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/uname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/unexpand"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/uniq"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/unlink"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/unset"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/uptime"
	userscmd "github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/users"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/vdir"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bash/wait"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/wc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/who"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/whoami"
	"github.com/yarencheng/go-bash-wasm/internal/commands/coreutils/yes"
	"github.com/yarencheng/go-bash-wasm/internal/shell"

	"github.com/mdp/qrterminal/v3"
)

var (
	BashVersion   = commands.BashVersion
	MachType      = commands.MachType
	BashCopyright = commands.BashCopyright
	BashLicense   = commands.BashLicense
	ProjectURL    = commands.ProjectURL
	SourceURL     = commands.SourceURL
	BtcAddress    = commands.BtcAddress
	EthAddress    = commands.EthAddress
)

const (
	ansiReset     = "\x1b[0m"
	ansiBold      = "\x1b[1m"
	ansiUnderline = "\x1b[4m"
	ansiDim       = "\x1b[2m"
	ansiCyan      = "\x1b[36m"
	ansiWhite     = "\x1b[37m"
	ansiYellow    = "\x1b[33m"
	ansiGray      = "\x1b[90m"
)

const Banner = ansiBold + ansiCyan + `
 ██████   ██████      ██████   █████   ██████  ██   ██     ██     ██   █████   ██████  ███    ███ 
██       ██    ██     ██   ██ ██   ██ ██       ██   ██     ██     ██  ██   ██ ██       ████  ████ 
██   ███ ██    ██     ██████  ███████  █████   ███████     ██  █  ██  ███████  █████   ██ ████ ██ 
██    ██ ██    ██     ██   ██ ██   ██      ██  ██   ██     ██ ███ ██  ██   ██      ██  ██  ██  ██ 
 ██████   ██████      ██████  ██   ██ ██████   ██   ██      ███ ███   ██   ██ ██████   ██      ██ 
` + ansiReset

// App encapsulates the bash simulator application.
type App struct {
	Registry *commands.Registry
	Env      *commands.Environment
	Logger   zerolog.Logger
}

// New creates a new bash simulator application.
func New(stdin io.ReadCloser, stdout, stderr io.Writer) *App {
	// Setup standard logger
	logger := zerolog.New(newLoggerWriter()).With().Timestamp().Logger()
	log.Logger = logger

	logger.Info().Msg("initializing go-bash-wasm application")

	// Setup virtual filesystem
	fs := afero.NewMemMapFs()
	setupMockFiles(fs, logger)

	// Setup command registry
	registry := commands.New()

	// Register commands
	cmds := []commands.Command{
		builtincmd.New(),
		ls.New(),
		colon.New(),
		commandcmd.New(),
		basename.New(),
		bind.New(),
		boolcmd.NewTrue(),
		boolcmd.NewFalse(),
		enable.New(),
		echo.New(),
		pwd.New(),
		pathchk.New(),
		read.New(),
		readonly.New(),
		readlink.New(),
		cat.New(),
		cksumcmd.New(),
		cd.New(),
		nl.New(),
		od.New(),
		paste.New(),
		pr.New(),
		ptx.New(),
		dirname.New(),
		mkdir.New(),
		mkfifo.New(),
		mknod.New(),
		mktemp.New(),
		rmdir.New(),
		runcon.New(),
		rm.New(),
		cp.New(),
		mv.New(),
		csplit.New(),
		head.New(),
		tail.New(),
		fc.New(),
		wc.New(),
		tty.New(),
		touch.New(),
		stat.New(),
		stdbuf.New(),
		stty.New(),
		id.New(),
		install.New(),
		ln.New(),
		link.New(),
		whoami.New(),
		uname.New(),
		ulimit.New(),
		userscmd.New(),
		nproccmd.New(),
		uptime.New(),
		hostid.New(),
		hostname.New(),
		nohup.New(),
		nice.New(),
		kill.New(),
		letcmd.New(),
		exit.New(),
		logout.New(),
		mapfile.New(),
		mapfile.NewWithName("readarray"),
		grep.New(),
		find.New(),
		getopts.New(),
		env.New(),
		export.New(),
		unset.New(),
		alias.New(),
		unalias.New(),
		declare.New(),
		declare.NewWithName("typeset"),
		groups.New(),
		hashcmd.New(),
		set.New(),
		shopt.New(),
		helpcmd.New(),
		historycmd.New(),
		logname.New(),
		getlimits.New(),
		pushd.New(),
		popd.New(),
		dirs.New(),
		base32cmd.New(),
		base64cmd.New(),
		basenc.New(),
		sumlegacy.New(),
		sum.New("md5sum"),
		sum.New("sha1sum"),
		sum.New("sha224sum"),
		sum.New("sha256sum"),
		sum.New("sha384sum"),
		sum.New("sha512sum"),
		shuf.New(),
		shred.New(),
		split.New(),
		tac.New(),
		suspend.New(),
		sync.New(),
		timeout.New(),
		trap.New(),
		truncate.New(),
		times.New(),
		tsort.New(),
		unexpand.New(),
		expand.New(),
		fmtcmd.New(),
		fold.New(),
		numfmt.New(),
		yes.New(),
		sleep.New(),
		printenv.New(),
		tr.New(),
		uniq.New(),
		sort.New(),
		cut.New(),
		join.New(),
		comm.New(),
		dir.New(),
		dircolors.New(),
		seq.New(),
		expr.New(),
		factor.New(),
		printf.New(),
		clear.New(),
		dd.New("dd"),
		test.New("test"),
		test.New("["),
		tee.New(),
		du.New(),
		df.New(),
		date.New(),
		who.New(),
		pinky.New(),
		arch.New(),
		chmod.New(),
		chown.New(),
		chcon.New(),
		chown.NewChgrp(),
		unlink.New(),
		realpath.New(),
		returncmd.New(),
		typecmd.New(),
		timecmd.New(),
		umask.New(),
		wait.New(),
		eval.New(),
		exec.New(),
		chroot.New(),
		jobs.New(),
		source.New(),
		source.NewDot(),
		shift.New(),
		bg.New(),
		fg.New(),
		breakcmd.New(),
		continuecmd.New(),
		disown.New(),
		complete.New(),
		compgen.New(),
		compopt.New(),
		caller.New(),
		local.New(),
		vdir.New(),
	}

	for _, cmd := range cmds {
		if err := registry.Register(cmd); err != nil {
			logger.Error().Err(err).Str("command", cmd.Name()).Msg("failed to register command")
		}
	}

	// Setup environment
	env := &commands.Environment{
		FS:        fs,
		Stdin:     stdin,
		Stdout:    stdout,
		Stderr:    stderr,
		Cwd:       "/",
		User:      "wasm",
		Uid:       1000,
		Gid:       1000,
		Umask:     022,
		Groups:    []int{1000},
		StartTime: time.Now(),
		EnvVars: map[string]string{
			"PATH":         "/usr/local/bin:/usr/bin:/bin",
			"USER":         "wasm",
			"HOME":         "/home/wasm",
			"PWD":          "/",
			"BASH_VERSION": BashVersion,
			"HOSTNAME":     "wasm-simulator",
			"HOSTTYPE":     "wasm32",
			"MACHTYPE":     MachType,
			"OSTYPE":       "linux-gnu",
			"TERM":         "xterm-256color",
			"SHELL":        "/bin/bash",
			"SHLVL":        "1",
			"UID":          "1000",
			"EUID":         "1000",
			"GID":          "1000",
			"BASH_SOURCE":  "(unknown)",
			"FUNCNAME":     "",
			"HISTSIZE":     "500",
			"HISTFILE":     "/home/wasm/.bash_history",
			"HISTFILESIZE": "500",
			"PS1":          "\\u@\\h:\\w\\$ ",
		},
		Aliases:     make(map[string]string),
		Arrays:      make(map[string][]string),
		DirStack:    make([]string, 0),
		Hash:        make(map[string]string),
		Jobs:        make([]*commands.Job, 0),
		Completions: make(map[string]*commands.CompSpec),
		Shopts: map[string]bool{
			"autocd":               false,
			"cdspell":              false,
			"checkwinsize":         true,
			"cmdhist":              true,
			"dirspell":             false,
			"dotglob":              false,
			"expand_aliases":       true,
			"extglob":              false,
			"globstar":             false,
			"histappend":           true,
			"interactive_comments": true,
			"lastpipe":             false,
			"nocaseglob":           false,
			"nullglob":             false,
			"progcomp":             true,
			"promptvars":           true,
			"sourcepath":           true,
		},
		Traps:    make(map[string]string),
		Registry: registry,
	}

	return &App{
		Registry: registry,
		Env:      env,
		Logger:   logger,
	}
}

// Run starts the interactive shell.
func (a *App) Run(ctx context.Context) error {
	a.ShowBanner()
	a.Logger.Info().Msg("starting interactive shell")
	shellObj := shell.New(a.Registry, a.Env)

	if err := shellObj.RunInteractive(); err != nil {
		a.Logger.Error().Err(err).Msg("shell session ended with error")
		return err
	}

	a.Logger.Info().Msg("shell session ended successfully")
	return nil
}

// ShowBanner prints the startup banner.
func (a *App) ShowBanner() {
	fmt.Fprint(a.Env.Stdout, Banner)

	divider := ansiGray + "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" + ansiReset + "\n"
	fmt.Fprint(a.Env.Stdout, divider)

	fmt.Fprintf(a.Env.Stdout, " %-12s %s\n", ansiGray+"Version:"+ansiReset, ansiBold+ansiWhite+BashVersion+" ("+MachType+")"+ansiReset)
	fmt.Fprintf(a.Env.Stdout, " %-12s %s\n", ansiGray+"Website:"+ansiReset, ansiUnderline+ansiCyan+ProjectURL+ansiReset)
	fmt.Fprintf(a.Env.Stdout, " %-12s %s\n", ansiGray+"GitHub: "+ansiReset, ansiUnderline+ansiCyan+SourceURL+ansiReset)

	fmt.Fprint(a.Env.Stdout, divider)
	fmt.Fprintln(a.Env.Stdout, " "+ansiBold+ansiYellow+"Support the Project (Donations):"+ansiReset)
	fmt.Fprint(a.Env.Stdout, "\n")

	var btc strings.Builder
	var eth strings.Builder

	qrterminal.GenerateHalfBlock(BtcAddress, qrterminal.M, &btc)
	qrterminal.GenerateHalfBlock(EthAddress, qrterminal.M, &eth)

	btcLines := strings.Split(strings.TrimSpace(btc.String()), "\n")
	ethLines := strings.Split(strings.TrimSpace(eth.String()), "\n")

	for i := range btcLines {
		prefix := "    "
		middle := "      "
		if i == len(btcLines)/2 {
			prefix = ansiBold + ansiYellow + "BTC " + ansiReset
			middle = "  " + ansiBold + ansiYellow + "ETH " + ansiReset
		}
		fmt.Fprintf(a.Env.Stdout, "%s%s%s%s\n", prefix, btcLines[i], middle, ethLines[i])
	}

	fmt.Fprint(a.Env.Stdout, "\n")
	fmt.Fprintf(a.Env.Stdout, " %s %s\n", ansiDim+"BTC Address:"+ansiReset, ansiBold+ansiWhite+BtcAddress+ansiReset)
	fmt.Fprintf(a.Env.Stdout, " %s %s\n", ansiDim+"ETH Address:"+ansiReset, ansiBold+ansiWhite+EthAddress+ansiReset)
	fmt.Fprint(a.Env.Stdout, divider)
}

// ShowVersion prints the version information of the bash simulator.
func (a *App) ShowVersion() {
	fmt.Fprintf(a.Env.Stdout, "go-bash-wasm, version %s (%s)\n", BashVersion, MachType)
	fmt.Fprintf(a.Env.Stdout, "%s\n", BashCopyright)
	fmt.Fprintf(a.Env.Stdout, "%s\n", BashLicense)
	fmt.Fprintf(a.Env.Stdout, "Home page:      <%s>\n", ProjectURL)
	fmt.Fprintf(a.Env.Stdout, "Source code:    <%s>\n", SourceURL)
	fmt.Fprintf(a.Env.Stdout, "Donation (BTC): <%s>\n", BtcAddress)
	fmt.Fprintf(a.Env.Stdout, "Donation (ETH): <%s>\n", EthAddress)
}

func setupMockFiles(fs afero.Fs, logger zerolog.Logger) {
	// Mock some files for initial testing
	_ = afero.WriteFile(fs, "/demo.txt", []byte("hello go-bash-wasm"), 0644)
	_ = fs.Mkdir("/configs", 0755)
	_ = afero.WriteFile(fs, "/configs/app.yaml", []byte("version: 0.1\nenv: development"), 0644)
	_ = afero.WriteFile(fs, "/welcome.log", []byte("bash simulator started"), 0644)

	logger.Debug().Msg("mock filesystem populated")
}
