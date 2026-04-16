package app

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ls"
	"github.com/yarencheng/go-bash-wasm/internal/shell"
)

// App encapsulates the bash simulator application.
type App struct {
	Registry *commands.Registry
	Env      *commands.Environment
	Logger   zerolog.Logger
}

// New creates a new bash simulator application.
func New(stdout, stderr io.Writer) *App {
	// Setup standard logger
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	log.Logger = logger

	logger.Info().Msg("initializing go-bash-wasm application")

	// Setup virtual filesystem
	fs := afero.NewMemMapFs()
	setupMockFiles(fs, logger)

	// Setup command registry
	registry := commands.New()
	lsCmd := ls.New()
	if err := registry.Register(lsCmd); err != nil {
		logger.Error().Err(err).Msg("failed to register ls command")
	}

	// Setup environment
	env := &commands.Environment{
		FS:     fs,
		Stdout: stdout,
		Stderr: stderr,
		Cwd:    "/",
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
