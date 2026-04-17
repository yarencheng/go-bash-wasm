package chmod

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

type Chmod struct{}

func New() *Chmod {
	return &Chmod{}
}

func (c *Chmod) Name() string {
	return "chmod"
}

func (c *Chmod) Run(ctx context.Context, env *commands.Environment, args []string) int {
	flags := pflag.NewFlagSet("chmod", pflag.ContinueOnError)
	recursive := flags.BoolP("recursive", "R", false, "change files and directories recursively")
	changes := flags.BoolP("changes", "c", false, "like verbose but report only when a change is made")
	silent := flags.BoolP("silent", "f", false, "suppress most error messages")
	verbose := flags.BoolP("verbose", "v", false, "output a diagnostic for every file processed")
	reference := flags.String("reference", "", "use RFILE's mode instead of MODE values")

	if err := flags.Parse(args); err != nil {
		if !*silent {
			fmt.Fprintf(env.Stderr, "chmod: %v\n", err)
		}
		return 1
	}

	remaining := flags.Args()
	if *reference == "" && len(remaining) < 2 {
		if !*silent {
			fmt.Fprintf(env.Stderr, "chmod: missing operand after '%s'\n", args[len(args)-1])
		}
		return 1
	}

	var mode os.FileMode
	var isSymbolic bool
	var modeStr string

	if *reference != "" {
		refPath := *reference
		if !filepath.IsAbs(refPath) {
			refPath = filepath.Join(env.Cwd, refPath)
		}
		info, err := env.FS.Stat(refPath)
		if err != nil {
			if !*silent {
				fmt.Fprintf(env.Stderr, "chmod: failed to get attributes of '%s': %v\n", *reference, err)
			}
			return 1
		}
		mode = info.Mode().Perm()
	} else {
		modeStr = remaining[0]
		remaining = remaining[1:]
		if m, err := strconv.ParseUint(modeStr, 8, 32); err == nil {
			mode = os.FileMode(m)
		} else {
			isSymbolic = true
		}
	}

	exitCode := 0
	for _, target := range remaining {
		fullPath := target
		if !filepath.IsAbs(target) {
			fullPath = filepath.Join(env.Cwd, target)
		}

		err := c.processPath(env, fullPath, mode, isSymbolic, modeStr, *recursive, *verbose, *changes, *silent)
		if err != nil {
			exitCode = 1
		}
	}

	return exitCode
}

func (c *Chmod) processPath(env *commands.Environment, path string, targetMode os.FileMode, isSymbolic bool, modeStr string, recursive, verbose, changes, silent bool) error {
	info, err := env.FS.Stat(path)
	if err != nil {
		if !silent {
			fmt.Fprintf(env.Stderr, "chmod: cannot access '%s': %v\n", path, err)
		}
		return err
	}

	oldMode := info.Mode().Perm()
	var newMode os.FileMode

	if isSymbolic {
		m, err := applySymbolicMode(oldMode, modeStr)
		if err != nil {
			if !silent {
				fmt.Fprintf(env.Stderr, "chmod: invalid mode: '%s'\n", modeStr)
			}
			return err
		}
		newMode = m
	} else {
		newMode = targetMode
	}

	if oldMode != newMode {
		err = env.FS.Chmod(path, newMode)
		if err != nil {
			if !silent {
				fmt.Fprintf(env.Stderr, "chmod: changing permissions of '%s': %v\n", path, err)
			}
			return err
		}
		if verbose || changes {
			fmt.Fprintf(env.Stdout, "mode of '%s' changed from %04o (%s) to %04o (%s)\n", path, oldMode, oldMode.String(), newMode, newMode.String())
		}
	} else {
		if verbose {
			fmt.Fprintf(env.Stdout, "mode of '%s' retained as %04o (%s)\n", path, oldMode, oldMode.String())
		}
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
			c.processPath(env, childPath, targetMode, isSymbolic, modeStr, recursive, verbose, changes, silent)
		}
	}

	return nil
}

// applySymbolicMode applies a symbolic mode string to an existing mode.
// This is a simplified implementation.
func applySymbolicMode(current os.FileMode, symMode string) (os.FileMode, error) {
	// [ugoa...][[-+=][perms...]...]
	// perms: rwx
	
	// Split by comma if multiple
	parts := strings.Split(symMode, ",")
	newMode := current

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		
		// Find operator
		opIdx := strings.IndexAny(part, "+-=")
		if opIdx == -1 {
			// If no operator, maybe it's just a numeric mode? No, handled elsewhere.
			return 0, fmt.Errorf("invalid symbolic mode")
		}

		who := part[:opIdx]
		op := part[opIdx]
		permsStr := part[opIdx+1:]

		var mask os.FileMode
		for _, p := range permsStr {
			switch p {
			case 'r':
				mask |= 0444
			case 'w':
				mask |= 0222
			case 'x':
				mask |= 0111
			default:
				// ignore unsupported for now
			}
		}

		// Apply "who" filters
		var whoMask os.FileMode
		if who == "" || strings.Contains(who, "a") {
			whoMask = 0777
		} else {
			if strings.Contains(who, "u") {
				whoMask |= 0700
			}
			if strings.Contains(who, "g") {
				whoMask |= 0070
			}
			if strings.Contains(who, "o") {
				whoMask |= 0007
			}
		}
		
		mask &= whoMask

		switch op {
		case '+':
			newMode |= mask
		case '-':
			newMode &= ^mask
		case '=':
			newMode = (newMode & ^whoMask) | mask
		}
	}

	return newMode, nil
}
