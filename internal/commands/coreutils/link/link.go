package link

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Link struct{}

func New() *Link {
	return &Link{}
}

func (l *Link) Name() string {
	return "link"
}

func (l *Link) Run(ctx context.Context, env *commands.Environment, args []string) int {
	if len(args) != 2 {
		fmt.Fprintf(env.Stderr, "link: call properly: link FILE1 FILE2\n")
		return 1
	}

	src := args[0]
	dst := args[1]

	if !filepath.IsAbs(src) {
		src = filepath.Join(env.Cwd, src)
	}
	if !filepath.IsAbs(dst) {
		dst = filepath.Join(env.Cwd, dst)
	}

	// Many filesystems don't support hard links.
	// afero.MemMapFs doesn't support Symlink or Link?
	// Actually afero has Linker interface.

	linker, ok := env.FS.(interface {
		Link(oldname, newname string) error
	})
	if !ok {
		// Fallback: copy for simulator if Link not supported?
		// No, better to report error or just copy if we want to be "fake" but usable.
		// GNU link is strictly for hard links.
		fmt.Fprintf(env.Stderr, "link: hard links not supported by filesystem\n")
		return 1
	}

	err := linker.Link(src, dst)
	if err != nil {
		fmt.Fprintf(env.Stderr, "link: cannot create link '%s' to '%s': %v\n", dst, src, err)
		return 1
	}

	return 0
}
