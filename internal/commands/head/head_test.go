package head

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

func TestHead_Name(t *testing.T) {
	h := New()
	assert.Equal(t, "head", h.Name())
}

func TestHead_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	require.NoError(t, afero.WriteFile(fs, "/f1.txt", []byte("1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/f2.txt", []byte("a\nb\nc\n"), 0644))
	require.NoError(t, afero.WriteFile(fs, "/zero.txt", []byte("x\x00y\x00z\x00"), 0644))

	tests := []struct {
		name       string
		args       []string
		stdin      string
		wantStatus int
		wantOut    string
		wantErr    string
	}{
		{
			name:       "default 10 lines",
			args:       []string{"/f1.txt"},
			wantStatus: 0,
			wantOut:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n",
		},
		{
			name:       "specified lines -n 2",
			args:       []string{"-n", "2", "/f1.txt"},
			wantStatus: 0,
			wantOut:    "1\n2\n",
		},
		{
			name:       "specified bytes -c 5",
			args:       []string{"-c", "5", "/f1.txt"},
			wantStatus: 0,
			wantOut:    "1\n2\n3",
		},
		{
			name:       "multiple files with headers",
			args:       []string{"/f2.txt", "/f2.txt"},
			wantStatus: 0,
			wantOut:    "==> /f2.txt <==\na\nb\nc\n\n==> /f2.txt <==\na\nb\nc\n",
		},
		{
			name:       "verbose header",
			args:       []string{"-v", "/f2.txt"},
			wantStatus: 0,
			wantOut:    "==> /f2.txt <==\na\nb\nc\n",
		},
		{
			name:       "quiet header",
			args:       []string{"-q", "/f2.txt", "/f2.txt"},
			wantStatus: 0,
			wantOut:    "a\nb\nc\na\nb\nc\n",
		},
		{
			name:       "zero terminated",
			args:       []string{"-z", "/zero.txt"},
			wantStatus: 0,
			wantOut:    "x\x00y\x00z\x00",
		},
		{
			name:       "zero terminated with -n 2",
			args:       []string{"-z", "-n", "2", "/zero.txt"},
			wantStatus: 0,
			wantOut:    "x\x00y\x00",
		},
		{
			name:       "stdin default",
			args:       []string{},
			stdin:      "s1\ns2\ns3\n",
			wantStatus: 0,
			wantOut:    "s1\ns2\ns3\n",
		},
		{
			name:       "stdin explicit",
			args:       []string{"-"},
			stdin:      "s1\ns2\ns3\n",
			wantStatus: 0,
			wantOut:    "s1\ns2\ns3\n",
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
			wantErr:    "cannot open '/nonexistent'",
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

			h := New()
			status := h.Run(context.Background(), env, tt.args)

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
