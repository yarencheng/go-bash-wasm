package df

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDf_Flags(t *testing.T) {
	d := New()

	t.Run("Inodes", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{Stdout: &out, Stderr: io.Discard}
		status := d.Run(context.Background(), env, []string{"-i"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "Inodes")
	})

	t.Run("Portability", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{Stdout: &out, Stderr: io.Discard}
		status := d.Run(context.Background(), env, []string{"-P"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "1024-blocks")
	})

	t.Run("Total", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{Stdout: &out, Stderr: io.Discard}
		status := d.Run(context.Background(), env, []string{"--total"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "total")
	})

	t.Run("HumanReadable", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{Stdout: &out, Stderr: io.Discard}
		status := d.Run(context.Background(), env, []string{"-h"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "1.0G")
	})
}
