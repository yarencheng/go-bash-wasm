package uptime

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestUptime_Run(t *testing.T) {
	startTime := time.Now().Add(-2 * time.Hour).Add(-30 * time.Minute)

	tests := []struct {
		name           string
		args           []string
		expectedStatus int
		containsOutput []string
	}{
		{
			name:           "default output",
			args:           []string{},
			expectedStatus: 0,
			containsOutput: []string{"up 2:30", "1 user", "load average"},
		},
		{
			name:           "pretty output",
			args:           []string{"-p"},
			expectedStatus: 0,
			containsOutput: []string{"up 2 hours, 30 minutes"},
		},
		{
			name:           "since output",
			args:           []string{"-s"},
			expectedStatus: 0,
			containsOutput: []string{startTime.Format("2006-01-02 15:04:05")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout:    stdout,
				StartTime: startTime,
			}
			u := New()
			status := u.Run(context.Background(), env, tt.args)
			assert.Equal(t, tt.expectedStatus, status)

			for _, expected := range tt.containsOutput {
				assert.Contains(t, stdout.String(), expected)
			}
		})
	}
}

func TestUptime_Metadata(t *testing.T) {
	u := New()
	assert.Equal(t, "uptime", u.Name())
}
