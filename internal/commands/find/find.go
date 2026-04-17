package find

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Find struct{}

func New() *Find {
	return &Find{}
}

func (f *Find) Name() string {
	return "find"
}

func (f *Find) Run(ctx context.Context, env *commands.Environment, args []string) int {
	var paths []string
	var namePattern string
	var typePattern string

	i := 0
	for i < len(args) {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-name":
				if i+1 < len(args) {
					namePattern = args[i+1]
					i += 2
					continue
				}
			case "-type":
				if i+1 < len(args) {
					typePattern = args[i+1]
					i += 2
					continue
				}
			}
		} else {
			paths = append(paths, arg)
			i++
			continue
		}
		i++
	}

	if len(paths) == 0 {
		paths = []string{"."}
	}

	for _, startPath := range paths {
		fullStartPath := startPath
		if !filepath.IsAbs(startPath) {
			fullStartPath = filepath.Join(env.Cwd, startPath)
		}

		err := afero.Walk(env.FS, fullStartPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Fprintf(env.Stderr, "find: %s: %v\n", path, err)
				return nil
			}

			// Filter by type
			if typePattern != "" {
				switch typePattern {
				case "d":
					if !info.IsDir() {
						return nil
					}
				case "f":
					if info.IsDir() {
						return nil
					}
				}
			}

			// Filter by name
			if namePattern != "" {
				matched, err := filepath.Match(namePattern, info.Name())
				if err != nil || !matched {
					return nil
				}
			}

			// Output relative or absolute as requested
			rel, err := filepath.Rel(env.Cwd, path)
			if err == nil && !filepath.IsAbs(startPath) {
				if !strings.HasPrefix(rel, ".") && !strings.HasPrefix(rel, "/") {
					rel = "./" + rel
				}
				fmt.Fprintln(env.Stdout, rel)
			} else {
				fmt.Fprintln(env.Stdout, path)
			}

			return nil
		})

		if err != nil {
			fmt.Fprintf(env.Stderr, "find: %v\n", err)
			return 1
		}
	}

	return 0
}
