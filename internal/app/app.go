package app

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/alias"
	base32cmd "github.com/yarencheng/go-bash-wasm/internal/commands/base32"
	base64cmd "github.com/yarencheng/go-bash-wasm/internal/commands/base64"
	"github.com/yarencheng/go-bash-wasm/internal/commands/basenc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/basename"
	"github.com/yarencheng/go-bash-wasm/internal/commands/boolcmd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/cat"
	"github.com/yarencheng/go-bash-wasm/internal/commands/csplit"
	"github.com/yarencheng/go-bash-wasm/internal/commands/cd"
	cksumcmd "github.com/yarencheng/go-bash-wasm/internal/commands/cksum"
	"github.com/yarencheng/go-bash-wasm/internal/commands/clear"
	"github.com/yarencheng/go-bash-wasm/internal/commands/colon"
	commandcmd "github.com/yarencheng/go-bash-wasm/internal/commands/command"
	"github.com/yarencheng/go-bash-wasm/internal/commands/comm"
	"github.com/yarencheng/go-bash-wasm/internal/commands/cp"
	"github.com/yarencheng/go-bash-wasm/internal/commands/cut"
	"github.com/yarencheng/go-bash-wasm/internal/commands/date"
	"github.com/yarencheng/go-bash-wasm/internal/commands/dd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/declare"
	"github.com/yarencheng/go-bash-wasm/internal/commands/getopts"
	"github.com/yarencheng/go-bash-wasm/internal/commands/df"
	"github.com/yarencheng/go-bash-wasm/internal/commands/dirname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/du"
	"github.com/yarencheng/go-bash-wasm/internal/commands/echo"
	env "github.com/yarencheng/go-bash-wasm/internal/commands/env"
	"github.com/yarencheng/go-bash-wasm/internal/commands/exit"
	"github.com/yarencheng/go-bash-wasm/internal/commands/export"
	"github.com/yarencheng/go-bash-wasm/internal/commands/factor"
	"github.com/yarencheng/go-bash-wasm/internal/commands/expr"
	"github.com/yarencheng/go-bash-wasm/internal/commands/find"
	"github.com/yarencheng/go-bash-wasm/internal/commands/grep"
	"github.com/yarencheng/go-bash-wasm/internal/commands/groups"
	hashcmd "github.com/yarencheng/go-bash-wasm/internal/commands/hash"
	helpcmd "github.com/yarencheng/go-bash-wasm/internal/commands/help"
	historycmd "github.com/yarencheng/go-bash-wasm/internal/commands/history"
	"github.com/yarencheng/go-bash-wasm/internal/commands/head"
	"github.com/yarencheng/go-bash-wasm/internal/commands/hostname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/hostid"
	"github.com/yarencheng/go-bash-wasm/internal/commands/id"
	"github.com/yarencheng/go-bash-wasm/internal/commands/install"
	"github.com/yarencheng/go-bash-wasm/internal/commands/link"
	join "github.com/yarencheng/go-bash-wasm/internal/commands/join"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ln"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ls"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mapfile"
	"github.com/yarencheng/go-bash-wasm/internal/commands/logout"
	"github.com/yarencheng/go-bash-wasm/internal/commands/logname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/letcmd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mkdir"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mktemp"
	"github.com/yarencheng/go-bash-wasm/internal/commands/mv"
	"github.com/yarencheng/go-bash-wasm/internal/commands/nice"
	"github.com/yarencheng/go-bash-wasm/internal/commands/nl"
	"github.com/yarencheng/go-bash-wasm/internal/commands/od"
	"github.com/yarencheng/go-bash-wasm/internal/commands/paste"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pr"
	"github.com/yarencheng/go-bash-wasm/internal/commands/nohup"
	nproccmd "github.com/yarencheng/go-bash-wasm/internal/commands/nproc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/kill"
	"github.com/yarencheng/go-bash-wasm/internal/commands/printenv"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pushd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/popd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/dirs"
	"github.com/yarencheng/go-bash-wasm/internal/commands/printf"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pwd"
	"github.com/yarencheng/go-bash-wasm/internal/commands/pathchk"
	"github.com/yarencheng/go-bash-wasm/internal/commands/read"
	"github.com/yarencheng/go-bash-wasm/internal/commands/readlink"
	"github.com/yarencheng/go-bash-wasm/internal/commands/rm"
	"github.com/yarencheng/go-bash-wasm/internal/commands/rmdir"
	"github.com/yarencheng/go-bash-wasm/internal/commands/seq"
	"github.com/yarencheng/go-bash-wasm/internal/commands/shuf"
	"github.com/yarencheng/go-bash-wasm/internal/commands/split"
	"github.com/yarencheng/go-bash-wasm/internal/commands/sleep"
	sort "github.com/yarencheng/go-bash-wasm/internal/commands/sort"
	"github.com/yarencheng/go-bash-wasm/internal/commands/stat"
	"github.com/yarencheng/go-bash-wasm/internal/commands/sum"
	"github.com/yarencheng/go-bash-wasm/internal/commands/sync"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tac"
	"github.com/yarencheng/go-bash-wasm/internal/commands/timeout"
	"github.com/yarencheng/go-bash-wasm/internal/commands/truncate"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tsort"
	"github.com/yarencheng/go-bash-wasm/internal/commands/unexpand"
	"github.com/yarencheng/go-bash-wasm/internal/commands/expand"
	"github.com/yarencheng/go-bash-wasm/internal/commands/fmt"
	"github.com/yarencheng/go-bash-wasm/internal/commands/fold"
	"github.com/yarencheng/go-bash-wasm/internal/commands/numfmt"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tail"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tee"
	test "github.com/yarencheng/go-bash-wasm/internal/commands/test"
	"github.com/yarencheng/go-bash-wasm/internal/commands/touch"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tty"
	"github.com/yarencheng/go-bash-wasm/internal/commands/tr"
	"github.com/yarencheng/go-bash-wasm/internal/commands/uname"
	"github.com/yarencheng/go-bash-wasm/internal/commands/unalias"
	"github.com/yarencheng/go-bash-wasm/internal/commands/uniq"
	userscmd "github.com/yarencheng/go-bash-wasm/internal/commands/users"
	"github.com/yarencheng/go-bash-wasm/internal/commands/unset"
	"github.com/yarencheng/go-bash-wasm/internal/commands/uptime"
	"github.com/yarencheng/go-bash-wasm/internal/commands/wc"
	"github.com/yarencheng/go-bash-wasm/internal/commands/who"
	"github.com/yarencheng/go-bash-wasm/internal/commands/whoami"
	"github.com/yarencheng/go-bash-wasm/internal/commands/yes"
	"github.com/yarencheng/go-bash-wasm/internal/commands/arch"
	"github.com/yarencheng/go-bash-wasm/internal/commands/chmod"
	"github.com/yarencheng/go-bash-wasm/internal/commands/chown"
	"github.com/yarencheng/go-bash-wasm/internal/commands/realpath"
	"github.com/yarencheng/go-bash-wasm/internal/commands/unlink"
	"github.com/yarencheng/go-bash-wasm/internal/shell"
	builtincmd "github.com/yarencheng/go-bash-wasm/internal/commands/builtin"
)

// App encapsulates the bash simulator application.
type App struct {
	Registry *commands.Registry
	Env      *commands.Environment
	Logger   zerolog.Logger
}

// New creates a new bash simulator application.
func New(stdin io.ReadCloser, stdout, stderr io.Writer) *App {
	// Setup standard logger
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
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
		boolcmd.NewTrue(),
		boolcmd.NewFalse(),
		echo.New(),
		pwd.New(),
		pathchk.New(),
		read.New(),
		readlink.New(),
		cat.New(),
		cksumcmd.New(),
		cd.New(),
		nl.New(),
		od.New(),
		paste.New(),
		pr.New(),
		basename.New(),
		dirname.New(),
		mkdir.New(),
		mktemp.New(),
		rmdir.New(),
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
		id.New(),
		install.New(),
		ln.New(),
		link.New(),
		whoami.New(),
		uname.New(),
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
		helpcmd.New(),
		historycmd.New(),
		logname.New(),
		pushd.New(),
		popd.New(),
		dirs.New(),
		base32cmd.New(),
		base64cmd.New(),
		basenc.New(),
		sum.New("md5sum"),
		sum.New("sha1sum"),
		sum.New("sha224sum"),
		sum.New("sha256sum"),
		sum.New("sha384sum"),
		sum.New("sha512sum"),
		shuf.New(),
		split.New(),
		tac.New(),
		sync.New(),
		timeout.New(),
		truncate.New(),
		tsort.New(),
		unexpand.New(),
		expand.New(),
		fmt.New(),
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
		arch.New(),
		chmod.New(),
		chown.New(),
		chown.NewChgrp(),
		unlink.New(),
		realpath.New(),
	}

	for _, cmd := range cmds {
		if err := registry.Register(cmd); err != nil {
			logger.Error().Err(err).Str("command", cmd.Name()).Msg("failed to register command")
		}
	}


	// Setup environment
	env := &commands.Environment{
		FS:     fs,
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
		Cwd:    "/",
		User:   "wasm",
		Uid:    1000,
		Gid:       1000,
		Groups:    []int{1000},
		StartTime: time.Now(),
		EnvVars: map[string]string{
			"PATH": "/usr/local/bin:/usr/bin:/bin",
			"USER": "wasm",
			"HOME": "/home/wasm",
			"PWD":  "/",
		},
		Aliases: make(map[string]string),
		Arrays:  make(map[string][]string),
		DirStack: make([]string, 0),
		Hash:    make(map[string]string),
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
	a.Logger.Info().Msg("starting interactive shell")
	shellObj := shell.New(a.Registry, a.Env)
	
	if err := shellObj.RunInteractive(); err != nil {
		a.Logger.Error().Err(err).Msg("shell session ended with error")
		return err
	}

	a.Logger.Info().Msg("shell session ended successfully")
	return nil
}

func setupMockFiles(fs afero.Fs, logger zerolog.Logger) {
	// Mock some files for initial testing
	_ = afero.WriteFile(fs, "/demo.txt", []byte("hello go-bash-wasm"), 0644)
	_ = fs.Mkdir("/configs", 0755)
	_ = afero.WriteFile(fs, "/configs/app.yaml", []byte("version: 0.1\nenv: development"), 0644)
	_ = afero.WriteFile(fs, "/welcome.log", []byte("bash simulator started"), 0644)
	
	logger.Debug().Msg("mock filesystem populated")
}
