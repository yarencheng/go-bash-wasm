package groups

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestGroups_Name(t *testing.T) {
	g := New()
	assert.Equal(t, "groups", g.Name())
}

func TestGroups_Run(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		user       string
		groups     []int
		wantStatus int
		wantOut    string
		wantErr    string
	}{
		{
			name:       "default user",
			args:       []string{},
			user:       "wasm",
			groups:     []int{1000},
			wantStatus: 0,
			wantOut:    "wasm\n",
		},
		{
			name:       "specified current user",
			args:       []string{"wasm"},
			user:       "wasm",
			groups:     []int{1000},
			wantStatus: 0,
			wantOut:    "wasm\n",
		},
		{
			name:       "multiple groups",
			args:       []string{},
			user:       "wasm",
			groups:     []int{1000, 1001},
			wantStatus: 0,
			wantOut:    "wasm wasm\n",
		},
		{
			name:       "non-existent user",
			args:       []string{"other"},
			user:       "wasm",
			wantStatus: 1,
			wantErr:    "no such user",
		},
		{
			name:       "mix of users",
			args:       []string{"wasm", "other"},
			user:       "wasm",
			wantStatus: 1,
			wantOut:    "wasm\n",
			wantErr:    "no such user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout: out,
				Stderr: errout,
				User:   tt.user,
				Groups: tt.groups,
			}

			g := New()
			status := g.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantOut != "" {
				assert.Contains(t, out.String(), tt.wantOut)
			}
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}
		})
	}
}
