package date

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDate_Flags(t *testing.T) {
	d := New()
	
	t.Run("UTC", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{Stdout: &out, Stderr: io.Discard}
		status := d.Run(context.Background(), env, []string{"-u", "-d", "2023-01-01 12:00:00"})
		assert.Equal(t, 0, status)
		assert.Contains(t, out.String(), "UTC")
	})

	t.Run("RFC2822", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{Stdout: &out, Stderr: io.Discard}
		status := d.Run(context.Background(), env, []string{"-R", "-d", "2023-01-01 12:00:00"})
		assert.Equal(t, 0, status)
		// RFC2822 usually starts with weekday
		assert.Regexp(t, `^[A-Z][a-z]{2}, `, out.String())
	})

	t.Run("ISO8601", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{Stdout: &out, Stderr: io.Discard}
		status := d.Run(context.Background(), env, []string{"-I", "-d", "2023-01-01"})
		assert.Equal(t, 0, status)
		assert.Equal(t, "2023-01-01\n", out.String())
	})

	t.Run("CustomFormat", func(t *testing.T) {
		var out bytes.Buffer
		env := &commands.Environment{Stdout: &out, Stderr: io.Discard}
		status := d.Run(context.Background(), env, []string{"-d", "2023-05-10", "+%Y/%m/%d"})
		assert.Equal(t, 0, status)
		assert.Equal(t, "2023/05/10\n", out.String())
	})
}
