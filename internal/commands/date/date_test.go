package date

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestDate_Run(t *testing.T) {
	var stdout bytes.Buffer
	env := &commands.Environment{
		Stdout: &stdout,
		Stderr: io.Discard,
	}

	d := New()
	status := d.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, status)
	// Basic check that it outputted something looking like a date
	assert.NotEmpty(t, stdout.String())
}
