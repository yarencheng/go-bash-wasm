package shell

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
	"github.com/yarencheng/go-bash-wasm/internal/commands/break"
	"github.com/yarencheng/go-bash-wasm/internal/commands/continue"
)

func TestExecuteTestClausePattern(t *testing.T) {
	s, _, _, _ := setupTestShell()

	// Pattern matching in [[ ]]
	assert.Equal(t, 0, s.Execute(context.Background(), "[[ hello == h* ]]"))
	assert.Equal(t, 0, s.Execute(context.Background(), "[[ hello != world ]]"))
	assert.Equal(t, 1, s.Execute(context.Background(), "[[ hello == world ]]"))
}

func TestExecuteTestClauseRegex(t *testing.T) {
	s, _, _, _ := setupTestShell()

	// Regex matching in [[ ]]
	assert.Equal(t, 0, s.Execute(context.Background(), "[[ hello =~ ^h.*o$ ]]"))
	assert.Equal(t, 1, s.Execute(context.Background(), "[[ hello =~ world ]]"))
}

func TestExecuteArithmExpansion(t *testing.T) {
	s, env, _, _ := setupTestShell()

	s.Execute(context.Background(), "val=$(( 1 + 1 ))")
	assert.Equal(t, "2", env.EnvVars["val"])

	s.Execute(context.Background(), "val2=$((1+1))")
	assert.Equal(t, "2", env.EnvVars["val2"])
}

func TestExecuteCommandSubstitution(t *testing.T) {
	s, env, _, _ := setupTestShell()

	env.Registry.Register(&mockCommand{
		name: "echo",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			fmt.Fprint(env.Stdout, strings.Join(args, " "))
			return 0
		},
	})

	exitCode := s.Execute(context.Background(), "echo $(echo hello) > /out")
	assert.Equal(t, 0, exitCode)
	data, _ := afero.ReadFile(env.FS, "/out")
	assert.Equal(t, "hello", strings.TrimSpace(string(data)))
}

func TestExecuteRedirectionExtra(t *testing.T) {
	s, env, _, _ := setupTestShell()

	env.Registry.Register(&mockCommand{
		name: "echo",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			fmt.Fprint(env.Stdout, strings.Join(args, " "))
			fmt.Fprint(env.Stderr, "error")
			return 0
		},
	})

	// 1. >| (Clobber)
	s.Execute(context.Background(), "echo hello >| /clobber.txt")
	data, err := afero.ReadFile(env.FS, "/clobber.txt")
	assert.NoError(t, err)
	assert.Equal(t, "hello", string(data))

	// 2. 2>&1
	s.Execute(context.Background(), "echo hello > /combined.txt 2>&1")
	data, _ = afero.ReadFile(env.FS, "/combined.txt")
	assert.Contains(t, string(data), "hello")
	assert.Contains(t, string(data), "error")

	// 3. <<< (Here-string)
	env.Registry.Register(&mockCommand{
		name: "cat",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			_, _ = io.Copy(env.Stdout, env.Stdin)
			return 0
		},
	})
	stdout := &strings.Builder{}
	env.Stdout = stdout
	s.Execute(context.Background(), "cat <<< 'world'")
	// Note: Here-string adds a newline
	assert.Equal(t, "world\n", stdout.String())
}

func TestExecuteArrayAssignment(t *testing.T) {
	s, env, _, _ := setupTestShell()

	s.Execute(context.Background(), "arr=(val1 val2 val3)")
	assert.Equal(t, []string{"val1", "val2", "val3"}, env.Arrays["arr"])
}
func TestExecuteFunction(t *testing.T) {
	s, env, _, _ := setupTestShell()
	env.Registry.Register(&mockCommand{
		name: "echo",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			fmt.Fprint(env.Stdout, strings.Join(args, " "))
			return 0
		},
	})

	// 1. Basic function definition and call
	s.Execute(context.Background(), "hello() { echo hi; }")
	assert.NotEmpty(t, env.Functions["hello"])

	s.Execute(context.Background(), "echo_val > /out") // setup mock echo
	s.Execute(context.Background(), "hello > /out")
	data, _ := afero.ReadFile(env.FS, "/out")
	assert.Contains(t, string(data), "hi")

	// 2. Positional arguments in function
	s.Execute(context.Background(), "greet() { echo hello $1; }")
	s.Execute(context.Background(), "greet world > /out_greet")
	data, _ = afero.ReadFile(env.FS, "/out_greet")
	assert.Contains(t, string(data), "hello world")
}

func TestExecuteLoopControl(t *testing.T) {
	s, env, _, _ := setupTestShell()
	env.Registry.Register(breakcmd.New())
	env.Registry.Register(continuecmd.New())
	env.Registry.Register(&mockCommand{
		name: "echo",
		run: func(ctx context.Context, env *commands.Environment, args []string) int {
			fmt.Fprint(env.Stdout, strings.Join(args, " "))
			return 0
		},
	})

	// 1. break in while
	s.Execute(context.Background(), "i=0; while [[ $i -lt 10 ]]; do i=$((i+1)); if [[ $i -eq 5 ]]; then break; fi; done")
	assert.Equal(t, "5", env.EnvVars["i"])

	// 2. continue in for
	s.Execute(context.Background(), "sum=0; for i in 1 2 3 4 5; do if [[ $i -eq 3 ]]; then continue; fi; sum=$((sum+i)); done")
	// sum should be 1+2+4+5 = 12
	assert.Equal(t, "12", env.EnvVars["sum"])
}
