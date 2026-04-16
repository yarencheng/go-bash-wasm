package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/ls"
)

func main() {
	fmt.Println("go-bash-wasm initialized")

	// Setup environment
	fs := afero.NewMemMapFs()
	// Mock some files
	_ = afero.WriteFile(fs, "/demo.txt", []byte("hello"), 0644)
	_ = fs.Mkdir("/configs", 0755)

	registry := commands.New()
	lsCmd := ls.New()
	_ = registry.Register(lsCmd)

	env := &commands.Environment{
		FS:     fs,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Cwd:    "/",
	}

	fmt.Println("\nRunning 'ls -F':")
	if cmd, ok := registry.Get("ls"); ok {
		cmd.Run(context.Background(), env, []string{"-F"})
	}
}
