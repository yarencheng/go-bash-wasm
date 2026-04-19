package cd

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Cd struct{}

func New() *Cd {
	return &Cd{}
}

func (c *Cd) Name() string {
	return "cd"
}

func (c *Cd) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("cd", pflag.ContinueOnError)
	usePhysical := flags.BoolP("physical", "P", false, "use the physical directory structure without following symbolic links")
	useLogical := flags.BoolP("logical", "L", false, "force symbolic links to be followed (default)")
	exitOnError := flags.BoolP("exit-on-error", "e", false, "if the -P option is supplied, and the current working directory cannot be determined successfully, exit with a non-zero status")

	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(env.Stderr, "cd: %v\n", err)
		return 1
	}

	targets := flags.Args()
	target := ""
	if len(targets) == 0 {
		target = env.EnvVars["HOME"]
		if target == "" {
			fmt.Fprintf(env.Stderr, "cd: HOME not set\n")
			return 1
		}
	} else {
		target = targets[0]
	}

	if target == "-" {
		target = env.EnvVars["OLDPWD"]
		if target == "" {
			fmt.Fprintf(env.Stderr, "cd: OLDPWD not set\n")
			return 1
		}
		// Special case: print new directory when using cd -
		fmt.Fprintln(env.Stdout, target)
	}

	var newPath string
	if path.IsAbs(target) || strings.HasPrefix(target, "./") || strings.HasPrefix(target, "../") || target == "." || target == ".." {
		newPath = resolvePath(env.Cwd, target)
	} else {
		// Search in CDPATH
		cdPath := env.EnvVars["CDPATH"]
		if cdPath == "" {
			newPath = resolvePath(env.Cwd, target)
		} else {
			paths := strings.Split(cdPath, ":")
			found := false
			for _, p := range paths {
				if p == "" {
					p = "."
				}
				candidate := resolvePath(env.Cwd, path.Join(p, target))
				if isDir(env, candidate) {
					newPath = candidate
					found = true
					// If we found it via CDPATH (and it's not .), print the name
					if p != "." {
						fmt.Fprintln(env.Stdout, candidate)
					}
					break
				}
			}
			if !found {
				newPath = resolvePath(env.Cwd, target)
			}
		}
	}

	newPath = path.Clean(newPath)

	// Handle -P (physical)
	if *usePhysical && !*useLogical {
		// In Afero/WASM, we don't have real symlinks in MemMapFs usually,
		// but let's assume we want the "real" path.
		// For now, path.Clean is the best we can do without more FS support.
	}

	info, err := env.FS.Stat(newPath)
	if err != nil {
		fmt.Fprintf(env.Stderr, "cd: %s: No such file or directory\n", target)
		if *exitOnError && *usePhysical {
			return 1
		}
		return 1
	}

	if !info.IsDir() {
		fmt.Fprintf(env.Stderr, "cd: %s: Not a directory\n", target)
		return 1
	}

	if env.EnvVars == nil {
		env.EnvVars = make(map[string]string)
	}
	env.EnvVars["OLDPWD"] = env.Cwd
	env.Cwd = newPath
	env.EnvVars["PWD"] = env.Cwd
	return 0
}

func resolvePath(cwd, target string) string {
	if path.IsAbs(target) {
		return target
	}
	return path.Join(cwd, target)
}

func isDir(env *commands.Environment, p string) bool {
	info, err := env.FS.Stat(p)
	if err != nil {
		return false
	}
	return info.IsDir()
}
