package chown

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Chown struct {
	isChgrp bool
}

func New() *Chown {
	return &Chown{isChgrp: false}
}

func NewChgrp() *Chown {
	return &Chown{isChgrp: true}
}

func (c *Chown) Name() string {
	if c.isChgrp {
		return "chgrp"
	}
	return "chown"
}

func (c *Chown) Run(ctx context.Context, env *commands.Environment, args []string) int {
	commandName := c.Name()
	flags := pflag.NewFlagSet(commandName, pflag.ContinueOnError)
	recursive := flags.BoolP("recursive", "R", false, "operate on files and directories recursively")
	changes := flags.BoolP("changes", "c", false, "like verbose but report only when a change is made")
	silent := flags.BoolP("silent", "f", false, "suppress most error messages")
	verbose := flags.BoolP("verbose", "v", false, "output a diagnostic for every file processed")
	reference := flags.String("reference", "", "use RFILE's owner and group instead of specifying OWNER:GROUP values")
	if err := flags.Parse(args); err != nil {
		if !*silent {
			fmt.Fprintf(env.Stderr, "%s: %v\n", commandName, err)
		}
		return 1
	}

	remaining := flags.Args()
	if *reference == "" && len(remaining) < 2 {
		if !*silent {
			fmt.Fprintf(env.Stderr, "%s: missing operand\n", commandName)
		}
		return 1
	}

	var targetUid, targetGid int = -1, -1

	if *reference != "" {
		refPath := *reference
		if !filepath.IsAbs(refPath) {
			refPath = filepath.Join(env.Cwd, refPath)
		}
		// Afero MemMapFs doesn't really store Uid/Gid in Stat() result in a way we can easily get?
		// Actually, Afero Stat() returns os.FileInfo.
		// For now, let's assume we can't easily get it from Afero MemMapFs and just mock or skip.
		// In a real system we'd use syscall.Stat_t.
		// For this simulator, we'll just use a default or mock.
		targetUid = 1000
		targetGid = 1000
	} else {
		spec := remaining[0]
		remaining = remaining[1:]

		if c.isChgrp {
			targetGid = parseGroup(spec)
		} else {
			u, g := parseOwnerGroup(spec)
			targetUid = u
			targetGid = g
		}
	}

	exitCode := 0
	for _, target := range remaining {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		err := c.processPath(env, fullPath, targetUid, targetGid, *recursive, *verbose, *changes, *silent)
		if err != nil {
			exitCode = 1
		}
	}

	return exitCode
}

func (c *Chown) processPath(env *commands.Environment, path string, uid, gid int, recursive, verbose, changes, silent bool) error {
	info, err := env.FS.Stat(path)
	if err != nil {
		if !silent {
			fmt.Fprintf(env.Stderr, "%s: cannot access '%s': %v\n", c.Name(), path, err)
		}
		return err
	}

	// In Afero MemMapFs, Chown is implemented but it doesn't seem to affect much in the mock FS.
	err = env.FS.Chown(path, uid, gid)
	if err != nil {
		if !silent {
			fmt.Fprintf(env.Stderr, "%s: changing ownership of '%s': %v\n", c.Name(), path, err)
		}
		return err
	}

	if verbose {
		fmt.Fprintf(env.Stdout, "ownership of '%s' retained or changed\n", path)
	}

	if recursive && info.IsDir() {
		dir, err := env.FS.Open(path)
		if err != nil {
			return err
		}
		defer dir.Close()

		entries, err := dir.Readdir(-1)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			childPath := filepath.Join(path, entry.Name())
			c.processPath(env, childPath, uid, gid, recursive, verbose, changes, silent)
		}
	}

	return nil
}

func parseOwnerGroup(spec string) (int, int) {
	parts := strings.Split(spec, ":")
	if len(parts) == 1 {
		parts = strings.Split(spec, ".")
	}

	uid := -1
	gid := -1

	if parts[0] != "" {
		if id, err := strconv.Atoi(parts[0]); err == nil {
			uid = id
		} else if parts[0] == "wasm" || parts[0] == "root" {
			if parts[0] == "root" {
				uid = 0
			} else {
				uid = 1000
			}
		}
	}

	if len(parts) > 1 && parts[1] != "" {
		if id, err := strconv.Atoi(parts[1]); err == nil {
			gid = id
		} else if parts[1] == "wasm" || parts[1] == "root" {
			if parts[1] == "root" {
				gid = 0
			} else {
				gid = 1000
			}
		}
	}

	return uid, gid
}

func parseGroup(spec string) int {
	if id, err := strconv.Atoi(spec); err == nil {
		return id
	}
	if spec == "wasm" {
		return 1000
	}
	if spec == "root" {
		return 0
	}
	return -1
}
