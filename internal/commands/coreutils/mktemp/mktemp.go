package mktemp

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Mktemp struct{}

func New() *Mktemp {
	return &Mktemp{}
}

func (m *Mktemp) Name() string {
	return "mktemp"
}

func (m *Mktemp) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("mktemp", pflag.ContinueOnError)
	directory := flags.BoolP("directory", "d", false, "create a directory, not a file")
	dryRun := flags.BoolP("dry-run", "u", false, "do not create anything; merely print a name")
	_ = flags.BoolP("quiet", "q", false, "suppress diagnostics about file/dir-creation failure")
	tmpdir := flags.String("tmpdir", "", "interpret TEMPLATE relative to DIR; if DIR is not specified, use $TMPDIR if set, else /tmp")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "mktemp: %v\n", err)
		return 1
	}

	targets := flags.Args()
	var template string
	if len(targets) > 0 {
		template = targets[0]
	} else {
		template = "tmp.XXXXXXXXXX"
	}

	dir := *tmpdir
	if dir == "" {
		dir = env.EnvVars["TMPDIR"]
	}
	if dir == "" {
		dir = "/tmp"
	}

	// Basic implementation of mktemp template
	// GNU mktemp template suffix is XXXXXX (at least 3)
	// afero.TempFile/TempDir use 'pattern' differently.

	pattern := template
	if strings.HasSuffix(template, "XXXXXX") {
		// afero uses '*' as a placeholder for random string
		pattern = strings.ReplaceAll(template, "XXXXXX", "*")
	}

	if *dryRun {
		// Very simplified dry run
		fmt.Fprintln(env.Stdout, filepath.Join(dir, strings.ReplaceAll(pattern, "*", "123456")))
		return 0
	}

	if *directory {
		name, err := afero.TempDir(env.FS, dir, pattern)
		if err != nil {
			fmt.Fprintf(env.Stderr, "mktemp: failed to create directory: %v\n", err)
			return 1
		}
		fmt.Fprintln(env.Stdout, name)
	} else {
		f, err := afero.TempFile(env.FS, dir, pattern)
		if err != nil {
			fmt.Fprintf(env.Stderr, "mktemp: failed to create file: %v\n", err)
			return 1
		}
		name := f.Name()
		f.Close()
		fmt.Fprintln(env.Stdout, name)
	}

	return 0
}
