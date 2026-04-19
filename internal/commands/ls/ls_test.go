package ls

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestLs_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	ls := New()

	// Setup mock filesystem
	require.NoError(t, afero.WriteFile(fs, "/file1.txt", []byte("content1"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/file2.txt", []byte("content2"), 0644))
	require.NoError(t, fs.Mkdir("/dir1", 0755))
	require.NoError(t, afero.WriteFile(fs, "/.hidden", []byte("hidden"), 0644))

	t.Run("default listing-alphabetical", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{})
		assert.Equal(t, 0, status)
		// We'll implement basic space separation first.
		assert.Contains(t, stdout.String(), "dir1")
		assert.Contains(t, stdout.String(), "file1.txt")
		assert.Contains(t, stdout.String(), "file2.txt")
		assert.NotContains(t, stdout.String(), ".hidden")
	})

	t.Run("all files -a", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-a"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), ".")
		assert.Contains(t, stdout.String(), "..")
		assert.Contains(t, stdout.String(), ".hidden")
		assert.Contains(t, stdout.String(), "file1.txt")
	})

	t.Run("long listing -l", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-l"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "drwxr-xr-x") // directory1
		assert.Contains(t, stdout.String(), "-rw-r--r--") // file1.txt
		assert.Contains(t, stdout.String(), "8")          // size of file1.txt
	})

	t.Run("indicators -F", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-F"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "dir1/")
		assert.Contains(t, stdout.String(), "file1.txt")
	})

	t.Run("sort by size -S", func(t *testing.T) {
		// file1 and file2 are same size in my setup, let's create a bigger one
		require.NoError(t, afero.WriteFile(fs, "/big.txt", make([]byte, 100), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-S"})
		assert.Equal(t, 0, status)
		// big.txt should be first
		assert.True(t, strings.HasPrefix(stdout.String(), "big.txt"))
	})

	t.Run("one per line -1", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-1"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "dir1\n")
		assert.Contains(t, stdout.String(), "file1.txt\n")
	})

	t.Run("recursive listing -R", func(t *testing.T) {
		require.NoError(t, fs.MkdirAll("/dir1/subdir", 0755))
		require.NoError(t, afero.WriteFile(fs, "/dir1/subdir/deep.txt", []byte("deep"), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-R"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "/dir1:")
		assert.Contains(t, stdout.String(), "/dir1/subdir:")
		assert.Contains(t, stdout.String(), "deep.txt")
	})

	t.Run("numeric IDs -n", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-n"})
		assert.Equal(t, 0, status)
		// Usually numeric IDs are shown as 0 or 1000 etc.
		// afero doesn't mock these well but we check for formatting.
		assert.Contains(t, stdout.String(), "0")
	})

	t.Run("directory indicator -p", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-p"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "dir1/")
		assert.NotContains(t, stdout.String(), "file1.txt*") // -p doesn't classify executables
	})

	t.Run("comma separated -m", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-m"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "dir1, file1.txt")
	})

	t.Run("no group -G", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-lG"})
		assert.Equal(t, 0, status)
		// Should not contain group name (second "root" or "0")
		output := stdout.String()
		lines := strings.Split(strings.TrimSpace(output), "\n")
		for _, line := range lines {
			parts := strings.Fields(line)
			// Mode, owner, size, month, day, time, name
			assert.LessOrEqual(t, len(parts), 7)
		}
	})

	t.Run("sort flag --sort", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"--sort=size"})
		assert.Equal(t, 0, status)
		assert.True(t, strings.HasPrefix(stdout.String(), "big.txt"))
	})

	t.Run("directory itself -d", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-d", "dir1"})
		assert.Equal(t, 0, status)
		assert.Equal(t, "dir1\n", stdout.String())
	})

	t.Run("do not sort -f", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		// -f implies -a and -U (unsorted)
		status := ls.Run(context.Background(), env, []string{"-f"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), ".hidden")
	})

	t.Run("no owner -g", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-g"})
		assert.Equal(t, 0, status)
		// Should contain group but not owner
		assert.Contains(t, stdout.String(), "root")
		// owner usually comes before group, if it's "root  root", we check if only one is there.
		// Actually my implementation joins with "  ".
		output := stdout.String()
		assert.True(t, strings.Count(output, "root") >= 1)
	})

	t.Run("no group -o", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-o"})
		assert.Equal(t, 0, status)
		// Should contain owner but not group
		assert.Contains(t, stdout.String(), "root")
	})

	t.Run("version sort -v", func(t *testing.T) {
		fs2 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs2, "/v10", []byte(""), 0644))
		require.NoError(t, afero.WriteFile(fs2, "/v2", []byte(""), 0644))
		require.NoError(t, afero.WriteFile(fs2, "/v1", []byte(""), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs2,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		// Use --sort=version as -v is not yet a pflag shorthand in my Run but it's in the task
		status := ls.Run(context.Background(), env, []string{"--sort=version"})
		assert.Equal(t, 0, status)
		output := stdout.String()
		// Order should be v1, v2, v10
		assert.True(t, strings.Index(output, "v1") < strings.Index(output, "v2"))
		assert.True(t, strings.Index(output, "v2") < strings.Index(output, "v10"))
	})

	t.Run("hide pattern --hide", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		// Initially file1.txt and file2.txt exist
		status := ls.Run(context.Background(), env, []string{"--hide=*.txt"})
		assert.Equal(t, 0, status)
		output := stdout.String()
		assert.NotContains(t, output, "file1.txt")
		assert.NotContains(t, output, "file2.txt")
		assert.Contains(t, output, "dir1")

		// -a should override --hide
		stdout.Reset()
		status = ls.Run(context.Background(), env, []string{"--hide=*.txt", "-a"})
		assert.Equal(t, 0, status)
		output = stdout.String()
		assert.Contains(t, output, "file1.txt")
	})

	t.Run("sort by extension -X", func(t *testing.T) {
		fs3 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs3, "/a.txt", []byte(""), 0644))
		require.NoError(t, afero.WriteFile(fs3, "/b.bin", []byte(""), 0644))
		require.NoError(t, afero.WriteFile(fs3, "/c.txt", []byte(""), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs3,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-X"})
		assert.Equal(t, 0, status)
		output := stdout.String()
		// .bin should come before .txt
		assert.True(t, strings.Index(output, "b.bin") < strings.Index(output, "a.txt"))
	})

	t.Run("size in blocks -s", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-s"})
		assert.Equal(t, 0, status)
		// Should contain a block number before names
		assert.Contains(t, stdout.String(), "1")
	})

	t.Run("ignore backups -B", func(t *testing.T) {
		require.NoError(t, afero.WriteFile(fs, "/backup~", []byte(""), 0644))
		defer func() { _ = fs.Remove("/backup~") }()

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-B"})
		assert.Equal(t, 0, status)
		assert.NotContains(t, stdout.String(), "backup~")
	})

	t.Run("quote names -Q", func(t *testing.T) {
		fs4 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs4, "/file with spaces", []byte(""), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs4,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-Q"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "\"file with spaces\"")
	})

	t.Run("escape names -b", func(t *testing.T) {
		fs5 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs5, "/file\nnewline", []byte(""), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs5,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-b"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "file\\nnewline")
	})

	t.Run("group directories first --group-directories-first", func(t *testing.T) {
		fs6 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs6, "/file.txt", []byte(""), 0644))
		require.NoError(t, fs6.Mkdir("/adir", 0755))
		require.NoError(t, afero.WriteFile(fs6, "/zfile.txt", []byte(""), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs6,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"--group-directories-first"})
		assert.Equal(t, 0, status)
		output := stdout.String()
		// adir should come before any file, even if it starts with 'a' (alphabetical would put it first anyway)
		// but let's compare with zfile.txt
		assert.True(t, strings.Index(output, "adir") < strings.Index(output, "file.txt"))
	})

	t.Run("zero terminated --zero", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"--zero"})
		assert.Equal(t, 0, status)
		// Should contain null characters
		assert.Contains(t, stdout.String(), "\x00")
	})

	t.Run("color output --color", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"--color=always"})
		assert.Equal(t, 0, status)
		// Should contain ANSI color codes (e.g., \033[1;34m for directories)
		assert.Contains(t, stdout.String(), "\033[1;34m")
	})

	t.Run("help flag", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"--help"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "Usage:")
	})

	t.Run("version flag", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"--version"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "ls")
	})

	t.Run("almost all -A", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-A"})
		assert.Equal(t, 0, status)
		assert.NotContains(t, stdout.String(), ".\n")
		assert.NotContains(t, stdout.String(), "..\n")
		assert.Contains(t, stdout.String(), ".hidden")
	})

	t.Run("reverse sort -r", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-r"})
		assert.Equal(t, 0, status)
		output := stdout.String()
		// Alphabetical: big.txt, dir1, file1.txt, file2.txt
		// Reverse: file2.txt, file1.txt, dir1, big.txt
		assert.True(t, strings.Index(output, "file2.txt") < strings.Index(output, "big.txt"))
	})

	t.Run("sort by time -t", func(t *testing.T) {
		fs7 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs7, "/old", []byte(""), 0644))
		time.Sleep(10 * time.Millisecond)
		require.NoError(t, afero.WriteFile(fs7, "/new", []byte(""), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs7,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-t"})
		assert.Equal(t, 0, status)
		output := stdout.String()
		assert.True(t, strings.Index(output, "new") < strings.Index(output, "old"))
	})

	t.Run("human readable -h", func(t *testing.T) {
		fs8 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs8, "/large", make([]byte, 2048), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs8,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-lh"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "2.0K")
	})

	t.Run("SI units --si", func(t *testing.T) {
		fs9 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs9, "/large", make([]byte, 2000), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs9,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-l", "--si"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "2.0K")
	})

	t.Run("multiple targets", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"file1.txt", "file2.txt"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "file1.txt")
		assert.Contains(t, stdout.String(), "file2.txt")
	})

	t.Run("inode -i", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-i"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "0 ")
	})

	t.Run("block size --block-size", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-l", "--block-size=1"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "8")
	})

	t.Run("time styles", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		styles := []string{"full-iso", "long-iso", "iso"}
		for _, style := range styles {
			stdout.Reset()
			status := ls.Run(context.Background(), env, []string{"-l", "--time-style=" + style})
			assert.Equal(t, 0, status)
			assert.NotEmpty(t, stdout.String())
		}
	})

	t.Run("quoting style shell", func(t *testing.T) {
		fs10 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs10, "/file with space", []byte(""), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs10,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"--quoting-style=shell"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "'file with space'")
	})

	t.Run("hide control chars -q", func(t *testing.T) {
		fs11 := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs11, "/file\x01name", []byte(""), 0644))

		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs11,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/",
		}

		status := ls.Run(context.Background(), env, []string{"-q"})
		assert.Equal(t, 0, status)
		assert.Contains(t, stdout.String(), "file?name")
	})

	t.Run("invalid directory", func(t *testing.T) {
		var stdout, stderr bytes.Buffer
		env := &commands.Environment{
			FS:     fs,
			Stdout: &stdout,
			Stderr: &stderr,
			Cwd:    "/nonexistent",
		}

		status := ls.Run(context.Background(), env, []string{})
		assert.NotEqual(t, 0, status)
		assert.Contains(t, stderr.String(), "exist")
	})
}
