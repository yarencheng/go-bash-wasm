package expr

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestExpr_Name(t *testing.T) {
	e := New()
	assert.Equal(t, "expr", e.Name())
}

func TestExpr_Run(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantStatus int
		wantOut    string
		wantErr    string
	}{
		{
			name:       "missing operand",
			args:       []string{},
			wantStatus: 2,
			wantErr:    "missing operand",
		},
		{
			name:       "help",
			args:       []string{"--help"},
			wantStatus: 0,
			wantOut:    "Usage: expr",
		},
		{
			name:       "version",
			args:       []string{"--version"},
			wantStatus: 0,
			wantOut:    "expr (GNU coreutils)",
		},
		{
			name:       "length",
			args:       []string{"length", "hello"},
			wantStatus: 0,
			wantOut:    "5\n",
		},
		{
			name:       "index found",
			args:       []string{"index", "hello", "e"},
			wantStatus: 0,
			wantOut:    "2\n",
		},
		{
			name:       "index not found",
			args:       []string{"index", "hello", "x"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "substr",
			args:       []string{"substr", "hello", "2", "3"},
			wantStatus: 0,
			wantOut:    "ell\n",
		},
		{
			name:       "substr out of range lower",
			args:       []string{"substr", "hello", "0", "3"},
			wantStatus: 0,
			wantOut:    "\n",
		},
		{
			name:       "substr out of range upper",
			args:       []string{"substr", "hello", "10", "3"},
			wantStatus: 0,
			wantOut:    "\n",
		},
		{
			name:       "substr partial",
			args:       []string{"substr", "hello", "4", "10"},
			wantStatus: 0,
			wantOut:    "lo\n",
		},
		{
			name:       "logical or s1",
			args:       []string{"5", "|", "3"},
			wantStatus: 0,
			wantOut:    "5\n",
		},
		{
			name:       "logical or s2",
			args:       []string{"0", "|", "3"},
			wantStatus: 0,
			wantOut:    "3\n",
		},
		{
			name:       "logical and both",
			args:       []string{"5", "&", "3"},
			wantStatus: 0,
			wantOut:    "5\n",
		},
		{
			name:       "logical and one zero",
			args:       []string{"5", "&", "0"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "addition",
			args:       []string{"1", "+", "2"},
			wantStatus: 0,
			wantOut:    "3\n",
		},
		{
			name:       "subtraction",
			args:       []string{"5", "-", "2"},
			wantStatus: 0,
			wantOut:    "3\n",
		},
		{
			name:       "multiplication",
			args:       []string{"3", "*", "4"},
			wantStatus: 0,
			wantOut:    "12\n",
		},
		{
			name:       "division",
			args:       []string{"10", "/", "2"},
			wantStatus: 0,
			wantOut:    "5\n",
		},
		{
			name:       "division by zero",
			args:       []string{"10", "/", "0"},
			wantStatus: 2,
			wantErr:    "division by zero",
		},
		{
			name:       "modulo",
			args:       []string{"10", "%", "3"},
			wantStatus: 0,
			wantOut:    "1\n",
		},
		{
			name:       "modulo by zero",
			args:       []string{"10", "%", "0"},
			wantStatus: 2,
			wantErr:    "division by zero",
		},
		{
			name:       "greater than true",
			args:       []string{"5", ">", "3"},
			wantStatus: 0,
			wantOut:    "1\n",
		},
		{
			name:       "greater than false",
			args:       []string{"3", ">", "5"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "less than true",
			args:       []string{"3", "<", "5"},
			wantStatus: 0,
			wantOut:    "1\n",
		},
		{
			name:       "less than false",
			args:       []string{"5", "<", "3"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "greater equal true",
			args:       []string{"5", ">=", "5"},
			wantStatus: 0,
			wantOut:    "1\n",
		},
		{
			name:       "greater equal false",
			args:       []string{"3", ">=", "5"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "less equal true",
			args:       []string{"5", "<=", "5"},
			wantStatus: 0,
			wantOut:    "1\n",
		},
		{
			name:       "less equal false",
			args:       []string{"5", "<=", "3"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "equal true",
			args:       []string{"5", "=", "5"},
			wantStatus: 0,
			wantOut:    "1\n",
		},
		{
			name:       "equal false",
			args:       []string{"5", "=", "3"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "not equal true",
			args:       []string{"5", "!=", "3"},
			wantStatus: 0,
			wantOut:    "1\n",
		},
		{
			name:       "not equal false",
			args:       []string{"5", "!=", "5"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "string equal true",
			args:       []string{"abc", "=", "abc"},
			wantStatus: 0,
			wantOut:    "1\n",
		},
		{
			name:       "string equal false",
			args:       []string{"abc", "=", "def"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "string not equal true",
			args:       []string{"abc", "!=", "def"},
			wantStatus: 0,
			wantOut:    "1\n",
		},
		{
			name:       "string not equal false",
			args:       []string{"abc", "!=", "abc"},
			wantStatus: 0,
			wantOut:    "0\n",
		},
		{
			name:       "single argument",
			args:       []string{"hello"},
			wantStatus: 0,
			wantOut:    "hello\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				Stdout: out,
				Stderr: errout,
			}

			e := New()
			status := e.Run(context.Background(), env, tt.args)

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
