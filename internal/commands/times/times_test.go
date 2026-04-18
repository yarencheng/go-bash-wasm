package times

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTimes(t *testing.T) {
	tm := New()
	stdout := &bytes.Buffer{}
	env := &commands.Environment{
		Stdout:    stdout,
		StartTime: time.Now().Add(-10 * time.Second),
	}

	code := tm.Run(context.Background(), env, []string{})
	assert.Equal(t, 0, code)
	
	output := stdout.String()
	lines := strings.Split(strings.TrimSpace(output), "\n")
	assert.Equal(t, 2, len(lines))
	// line 1 should have some time values
	assert.Contains(t, string(lines[0]), "m")
	// line 2 should be 0s
	assert.Contains(t, string(lines[1]), "0m0.000s 0m0.000s")
}
