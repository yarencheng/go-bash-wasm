package basenc

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestBasenc_Base16(t *testing.T) {
	out := &strings.Builder{}
	in := io.NopCloser(strings.NewReader("hello"))
	env := &commands.Environment{
		Stdin:  in,
		Stdout: out,
	}

	b := New()
	status := b.Run(context.Background(), env, []string{"--base16"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "68656c6c6f\n", out.String())
}

func TestBasenc_Base32Hex(t *testing.T) {
	out := &strings.Builder{}
	in := io.NopCloser(strings.NewReader("hello"))
	env := &commands.Environment{
		Stdin:  in,
		Stdout: out,
	}

	b := New()
	status := b.Run(context.Background(), env, []string{"--base32hex"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "D1IMOR3F\n", out.String())
}

func TestBasenc_Base2msbf(t *testing.T) {
	out := &strings.Builder{}
	in := io.NopCloser(strings.NewReader("A"))
	env := &commands.Environment{
		Stdin:  in,
		Stdout: out,
	}

	b := New()
	status := b.Run(context.Background(), env, []string{"--base2msbf"})
	assert.Equal(t, 0, status)
	// 'A' is 65, which is 01000001
	assert.Equal(t, "01000001\n", out.String())
}

func TestBasenc_Base2msbf_Decode(t *testing.T) {
	out := &strings.Builder{}
	in := io.NopCloser(strings.NewReader("01000001"))
	env := &commands.Environment{
		Stdin:  in,
		Stdout: out,
	}

	b := New()
	status := b.Run(context.Background(), env, []string{"--base2msbf", "-d"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "A", out.String())
}

func TestBasenc_Decode(t *testing.T) {
	out := &strings.Builder{}
	in := io.NopCloser(strings.NewReader("68656c6c6f"))
	env := &commands.Environment{
		Stdin:  in,
		Stdout: out,
	}

	b := New()
	status := b.Run(context.Background(), env, []string{"--base16", "-d"})
	assert.Equal(t, 0, status)
	assert.Equal(t, "hello", out.String())
}
