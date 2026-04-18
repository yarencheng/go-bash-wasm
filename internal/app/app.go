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
	"github.com/yarencheng/go-bash-wasm/internal/commands/alias"
	"github.com/yarencheng/go-bash-wasm/internal/commands/arch"
	base32cmd "github.com/yarencheng/go-bash-wasm/internal/commands/base32"
	base64cmd "github.com/yarencheng/go-bash-wasm/internal/commands/base64"
	"github.com/yarencheng/go-bash-wasm/internal/commands/basename"
	"github.com/yarencheng/go-bash-wasm/internal/commands/basenc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bg"
	"github.com/yarencheng/go-bash-wasm/internal/commands/bind"
	"github.com/yarencheng/go-bash-wasm/internal/commands/boolcmd"
	breakcmd "github.com/yarencheng/go-bash-wasm/internal/commands/break"
	builtincmd "github.com/yarencheng/go-bash-wasm/internal/commands/builtin"
	"github.com/yarencheng/go-bash-wasm/internal/commands/caller"
	"github.com/yarencheng/go-bash-wasm/internal/commands/cat"
	"github.com/yarencheng/go-bash-wasm/internal/commands/cd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/chcon"
	"github.com/yarencheng/go-bash-wasm/internal/commands/chmod"
	"github.com/yarencheng/go-bash-wasm/internal/commands/chown"
	"github.com/yarencheng/go-bash-wasm/internal/commands/chroot"
	cksumcmd "github.com/yarencheng/go-bash-wasm/internal/commands/cksum"
	"github.com/yarencheng/go-bash-wasm/internal/commands/clear"
	"github.com/yarencheng/go-bash-wasm/internal/commands/colon"
	"github.com/yarencheng/go-bash-wasm/internal/commands/comm"
	commandcmd "github.com/yarencheng/go-bash-wasm/internal/commands/command"
	"github.com/yarencheng/go-bash-wasm/internal/commands/compgen"
	"github.com/yarencheng/go-bash-wasm/internal/commands/complete"
	"github.com/yarencheng/go-bash-wasm/internal/commands/compopt"
	continuecmd "github.com/yarencheng/go-bash-wasm/internal/commands/continue"
	"github.com/yarencheng/go-bash-wasm/internal/commands/cp"
	"github.com/yarencheng/go-bash-wasm/internal/commands/csplit"
	"github.com/yarencheng/go-bash-wasm/internal/commands/cut"
	"github.com/yarencheng/go-bash-wasm/internal/commands/date"
	"github.com/yarencheng/go-bash-wasm/internal/commands/dd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/declare"
	"github.com/yarencheng/go-bash-wasm/internal/commands/df"
	"github.com/yarencheng/go-bash-wasm/internal/commands/dir"
	"github.com/yarencheng/go-bash-wasm/internal/commands/dircolors"
	"github.com/yarencheng/go-bash-wasm/internal/commands/dirname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/dirs"
	"github.com/yarencheng/go-bash-wasm/internal/commands/disown"
	"github.com/yarencheng/go-bash-wasm/internal/commands/du"
	"github.com/yarencheng/go-bash-wasm/internal/commands/echo"
	"github.com/yarencheng/go-bash-wasm/internal/commands/enable"
	env "github.com/yarencheng/go-bash-wasm/internal/commands/env"
	"github.com/yarencheng/go-bash-wasm/internal/commands/eval"
	"github.com/yarencheng/go-bash-wasm/internal/commands/exec"
	"github.com/yarencheng/go-bash-wasm/internal/commands/exit"
	"github.com/yarencheng/go-bash-wasm/internal/commands/expand"
	"github.com/yarencheng/go-bash-wasm/internal/commands/export"
	"github.com/yarencheng/go-bash-wasm/internal/commands/expr"
	"github.com/yarencheng/go-bash-wasm/internal/commands/factor"
	"github.com/yarencheng/go-bash-wasm/internal/commands/fg"
	"github.com/yarencheng/go-bash-wasm/internal/commands/find"
	fmtcmd "github.com/yarencheng/go-bash-wasm/internal/commands/fmt"
	"github.com/yarencheng/go-bash-wasm/internal/commands/fold"
	"github.com/yarencheng/go-bash-wasm/internal/commands/getopts"
	"github.com/yarencheng/go-bash-wasm/internal/commands/grep"
	"github.com/yarencheng/go-bash-wasm/internal/commands/groups"
	hashcmd "github.com/yarencheng/go-bash-wasm/internal/commands/hash"
	"github.com/yarencheng/go-bash-wasm/internal/commands/getlimits"
	"github.com/yarencheng/go-bash-wasm/internal/commands/head"
	helpcmd "github.com/yarencheng/go-bash-wasm/internal/commands/help"
	historycmd "github.com/yarencheng/go-bash-wasm/internal/commands/history"
	"github.com/yarencheng/go-bash-wasm/internal/commands/hostid"
	"github.com/yarencheng/go-bash-wasm/internal/commands/hostname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/id"
	"github.com/yarencheng/go-bash-wasm/internal/commands/install"
	"github.com/yarencheng/go-bash-wasm/internal/commands/jobs"
	join "github.com/yarencheng/go-bash-wasm/internal/commands/join"
	"github.com/yarencheng/go-bash-wasm/internal/commands/kill"
	"github.com/yarencheng/go-bash-wasm/internal/commands/letcmd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/link"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ln"
	"github.com/yarencheng/go-bash-wasm/internal/commands/local"
	"github.com/yarencheng/go-bash-wasm/internal/commands/logname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/logout"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ls"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mapfile"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mkdir"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mkfifo"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mknod"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mktemp"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mv"
	"github.com/yarencheng/go-bash-wasm/internal/commands/nice"
	"github.com/yarencheng/go-bash-wasm/internal/commands/nl"
	"github.com/yarencheng/go-bash-wasm/internal/commands/nohup"
	nproccmd "github.com/yarencheng/go-bash-wasm/internal/commands/nproc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/numfmt"
	"github.com/yarencheng/go-bash-wasm/internal/commands/od"
	"github.com/yarencheng/go-bash-wasm/internal/commands/paste"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pathchk"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pinky"
	"github.com/yarencheng/go-bash-wasm/internal/commands/popd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pr"
	"github.com/yarencheng/go-bash-wasm/internal/commands/printenv"
	"github.com/yarencheng/go-bash-wasm/internal/commands/printf"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ptx"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pushd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pwd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/read"
	"github.com/yarencheng/go-bash-wasm/internal/commands/readlink"
	"github.com/yarencheng/go-bash-wasm/internal/commands/readonly"
	"github.com/yarencheng/go-bash-wasm/internal/commands/realpath"
	"github.com/yarencheng/go-bash-wasm/internal/commands/returncmd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/rm"
	"github.com/yarencheng/go-bash-wasm/internal/commands/rmdir"
	"github.com/yarencheng/go-bash-wasm/internal/commands/runcon"
	"github.com/yarencheng/go-bash-wasm/internal/commands/seq"
	"github.com/yarencheng/go-bash-wasm/internal/commands/set"
	"github.com/yarencheng/go-bash-wasm/internal/commands/shift"
	"github.com/yarencheng/go-bash-wasm/internal/commands/shred"
	"github.com/yarencheng/go-bash-wasm/internal/commands/shuf"
	"github.com/yarencheng/go-bash-wasm/internal/commands/sleep"
	"github.com/yarencheng/go-bash-wasm/internal/commands/shopt"
	sort "github.com/yarencheng/go-bash-wasm/internal/commands/sort"
	"github.com/yarencheng/go-bash-wasm/internal/commands/source"
	"github.com/yarencheng/go-bash-wasm/internal/commands/split"
	"github.com/yarencheng/go-bash-wasm/internal/commands/stat"
	"github.com/yarencheng/go-bash-wasm/internal/commands/stdbuf"
	"github.com/yarencheng/go-bash-wasm/internal/commands/stty"
	"github.com/yarencheng/go-bash-wasm/internal/commands/sum"
	"github.com/yarencheng/go-bash-wasm/internal/commands/sumlegacy"
	"github.com/yarencheng/go-bash-wasm/internal/commands/suspend"
	"github.com/yarencheng/go-bash-wasm/internal/commands/sync"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tac"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tail"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tee"
	test "github.com/yarencheng/go-bash-wasm/internal/commands/test"
	timecmd "github.com/yarencheng/go-bash-wasm/internal/commands/time"
	"github.com/yarencheng/go-bash-wasm/internal/commands/times"
	"github.com/yarencheng/go-bash-wasm/internal/commands/timeout"
	"github.com/yarencheng/go-bash-wasm/internal/commands/touch"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tr"
	"github.com/yarencheng/go-bash-wasm/internal/commands/trap"
	"github.com/yarencheng/go-bash-wasm/internal/commands/truncate"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tsort"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tty"
	typecmd "github.com/yarencheng/go-bash-wasm/internal/commands/type"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ulimit"
	"github.com/yarencheng/go-bash-wasm/internal/commands/umask"
	"github.com/yarencheng/go-bash-wasm/internal/commands/unalias"
	"github.com/yarencheng/go-bash-wasm/internal/commands/uname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/unexpand"
	"github.com/yarencheng/go-bash-wasm/internal/commands/uniq"
	"github.com/yarencheng/go-bash-wasm/internal/commands/unlink"
	"github.com/yarencheng/go-bash-wasm/internal/commands/unset"
	"github.com/yarencheng/go-bash-wasm/internal/commands/uptime"
	userscmd "github.com/yarencheng/go-bash-wasm/internal/commands/users"
	"github.com/yarencheng/go-bash-wasm/internal/commands/wait"
	"github.com/yarencheng/go-bash-wasm/internal/commands/wc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/who"
	"github.com/yarencheng/go-bash-wasm/internal/commands/whoami"
	"github.com/yarencheng/go-bash-wasm/internal/commands/yes"
	"github.com/yarencheng/go-bash-wasm/internal/commands/vdir"
	"github.com/yarencheng/go-bash-wasm/internal/shell"

	"github.com/mdp/qrterminal/v3"
)

var (
	BashVersion   = "5.3-rc"
	MachType      = "wasm32-unknown-wasi"
	BashCopyright = "Copyright (C) 2026 go-bash-wasm team"
	BashLicense   = "License Apache-2.0: Apache License, Version 2.0 <http://www.apache.org/licenses/LICENSE-2.0>"
	ProjectURL    = "https://bash.devops-playground.dev/"
	SourceURL     = "https://github.com/yarencheng/go-bash-wasm"
	BtcAddress    = "bc1qntk3pkvlkd9kg8kgjyff85rplyal7jl7t3pyxl"
	EthAddress    = "0xfC689a03D0BF58b27f1928eed952CAbCF816fAE9"
)

const Banner = `
 ██████   ██████      ██████   █████   ██████  ██   ██     ██     ██   █████   ██████  ███    ███ 
██       ██    ██     ██   ██ ██   ██ ██       ██   ██     ██     ██  ██   ██ ██       ████  ████ 
██   ███ ██    ██     ██████  ███████  █████   ███████     ██  █  ██  ███████  █████   ██ ████ ██ 
██    ██ ██    ██     ██   ██ ██   ██      ██  ██   ██     ██ ███ ██  ██   ██      ██  ██  ██  ██ 
 ██████   ██████      ██████  ██   ██ ██████   ██   ██      ███ ███   ██   ██ ██████   ██      ██ 
`

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
		complete.New(),
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
	fmt.Fprint(a.Env.Stdout, "\n")

	var btc strings.Builder
	var eth strings.Builder

	qrterminal.GenerateHalfBlock(BtcAddress, qrterminal.M, &btc)
	qrterminal.GenerateHalfBlock(EthAddress, qrterminal.M, &eth)

	var merge string
	btcLines := strings.Split(btc.String(), "\n")
	ethLines := strings.Split(eth.String(), "\n")

	for i := range btcLines {
		if i == 10 {
			merge += "BTC " + btcLines[i] + " ETH " + ethLines[i] + "\n"
		} else {
			merge += "    " + btcLines[i] + "     " + ethLines[i] + "\n"
		}
	}

	fmt.Fprint(a.Env.Stdout, merge)
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
