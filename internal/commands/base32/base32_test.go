package base32cmd

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestBase32_Run_Encode(t *testing.T) {
	out := &strings.Builder{}
	in := io.NopCloser(strings.NewReader("hello"))
	env := &commands.Environment{
		Stdin:  in,
		Stdout: out,
	}

	b := New()
	status := b.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	assert.Equal(t, "NBSWY3DP\n", out.String())
}

func TestBase32_Run_Decode(t *testing.T) {
	out := &strings.Builder{}
	in := io.NopCloser(strings.NewReader("NBSWY3DP"))
	env := &commands.Environment{
		Stdin:  in,
		Stdout: out,
	}

	b := New()
	status := b.Run(context.Background(), env, []string{"-d"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello", out.String())
}

func TestBase32_Run_Wrap(t *testing.T) {
	out := &strings.Builder{}
	in := io.NopCloser(strings.NewReader("hello world this is a test"))
	env := &commands.Environment{
		Stdin:  in,
		Stdout: out,
	}

	b := New()
	// "hello world this is a test" -> "NBSWY3DPEB3W64TMMQQHI2DJOMQGS4ZAMEQHIZLTOQ======"
	status := b.Run(context.Background(), env, []string{"-w", "10"})
	assert.Equal(t, 0, status)
	expected := "NBSWY3DPEB\n3W64TMMQQH\nI2DJOMQGS4\nZAMEQHIZLT\nOQ======\n"
	assert.Equal(t, expected, out.String())
}

func TestBase32_Run_File(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "/test.txt", []byte("hello"), 0644)
	out := &strings.Builder{}
	env := &commands.Environment{
		FS:     fs,
		Stdout: out,
		Cwd:    "/",
	}

	b := New()
	status := b.Run(context.Background(), env, []string{"test.txt"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "NBSWY3DP\n", out.String())
}
