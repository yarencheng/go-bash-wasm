package paste

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestPaste_Name(t *testing.T) {
	p := New()
	assert.Equal(t, "paste", p.Name())
}

func TestPaste_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/f1.txt", []byte("1\n2\n"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/f2.txt", []byte("a\nb\nc\n"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/zero.txt", []byte("x\x00y\x00"), 0644))

	tests := []struct {
		name       string
		args       []string
		stdin      string
		wantStatus int
		wantOut    string
		wantErr    string
	}{
		{
			name:       "parallel default",
			args:       []string{"/f1.txt", "/f2.txt"},
			wantStatus: 0,
			wantOut:    "1\ta\n2\tb\n\tc\n",
		},
		{
			name:       "parallel with custom delimiter",
			args:       []string{"-d", ",", "/f1.txt", "/f2.txt"},
			wantStatus: 0,
			wantOut:    "1,a\n2,b\n,c\n",
		},
		{
			name:       "parallel with multiple delimiters",
			args:       []string{"-d", ",:", "/f1.txt", "/f2.txt", "/f1.txt"},
			wantStatus: 0,
			wantOut:    "1,a:1\n2,b:2\n,c:\n",
		},
		{
			name:       "serial mode",
			args:       []string{"-s", "/f1.txt", "/f2.txt"},
			wantStatus: 0,
			wantOut:    "1\t2\na\tb\tc\n",
		},
		{
			name:       "serial mode with delimiter",
			args:       []string{"-s", "-d", ",", "/f1.txt", "/f2.txt"},
			wantStatus: 0,
			wantOut:    "1,2\na,b,c\n",
		},
		{
			name:       "zero terminated parallel",
			args:       []string{"-z", "/zero.txt", "/zero.txt"},
			wantStatus: 0,
			wantOut:    "x\tx\x00y\ty\x00",
		},
		{
			name:       "zero terminated serial",
			args:       []string{"-z", "-s", "/zero.txt", "/zero.txt"},
			wantStatus: 0,
			wantOut:    "x\ty\x00x\ty\x00",
		},
		{
			name:       "stdin no args",
			args:       []string{},
			stdin:      "s1\ns2\n",
			wantStatus: 0,
			wantOut:    "s1\ns2\n",
		},
		{
			name:       "stdin explicit",
			args:       []string{"/f1.txt", "-", "/f2.txt"},
			stdin:      "i1\ni2\n",
			wantStatus: 0,
			wantOut:    "1\ti1\ta\n2\ti2\tb\n\t\tc\n",
		},
		{
			name:       "invalid flag",
			args:       []string{"--invalid"},
			wantStatus: 1,
			wantErr:    "unknown flag",
		},
		{
			name:       "file not found",
			args:       []string{"/nonexistent"},
			wantStatus: 1,
			wantErr:    "file does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			errout := &bytes.Buffer{}
			env := &commands.Environment{
				FS:     fs,
				Cwd:    "/",
				Stdin:  io.NopCloser(strings.NewReader(tt.stdin)),
				Stdout: out,
				Stderr: errout,
			}

			p := New()
			status := p.Run(context.Background(), env, tt.args)

			assert.Equal(t, tt.wantStatus, status)
			if tt.wantOut != "" {
				assert.Equal(t, tt.wantOut, out.String())
			}
			if tt.wantErr != "" {
				assert.Contains(t, errout.String(), tt.wantErr)
			}
		})
	}
}

func TestScanNull(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		atEOF    bool
		wantAdv  int
		wantTok  []byte
		wantErr  bool
	}{
		{
			name:    "empty data at EOF",
			data:    []byte{},
			atEOF:   true,
			wantAdv: 0,
			wantTok: nil,
		},
		{
			name:    "data with NUL",
			data:    []byte("abc\x00def"),
			atEOF:   false,
			wantAdv: 4,
			wantTok: []byte("abc"),
		},
		{
			name:    "data without NUL not EOF",
			data:    []byte("abc"),
			atEOF:   false,
			wantAdv: 0,
			wantTok: nil,
		},
		{
			name:    "data without NUL at EOF",
			data:    []byte("abc"),
			atEOF:   true,
			wantAdv: 3,
			wantTok: []byte("abc"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adv, tok, err := scanNull(tt.data, tt.atEOF)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantAdv, adv)
				assert.Equal(t, tt.wantTok, tok)
			}
		})
	}
}
