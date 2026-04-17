package touch

import (
	"context"
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/yarencheng/go-bash-wasm/internal/commands"
)

func TestTouch_Run(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	tr := New()

	// Test basic touch (create file)
	status := tr.Run(context.Background(), env, []string{"newfile.txt"})
	assert.Equal(t, 0, status)
	_, err := fs.Stat("/newfile.txt")
	assert.NoError(t, err)

	// Test touch existing file
	status = tr.Run(context.Background(), env, []string{"newfile.txt"})
	assert.Equal(t, 0, status)
}

func TestTouch_NoCreate(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	tr := New()
	status := tr.Run(context.Background(), env, []string{"-c", "nocreate.txt"})
	assert.Equal(t, 0, status)
	_, err := fs.Stat("/nocreate.txt")
	assert.Error(t, err)
}

func TestTouch_Reference(t *testing.T) {
	fs := afero.NewMemMapFs()
	env := &commands.Environment{
		FS:     fs,
		Cwd:    "/",
		Stderr: io.Discard,
	}

	// Create a reference file
	refFile := "/ref.txt"
	_ = afero.WriteFile(fs, refFile, []byte("ref"), 0644)
	refInfo, _ := fs.Stat(refFile)
	refMtime := refInfo.ModTime()

	tr := New()
	// Create another file
	targetFile := "/target.txt"
	status := tr.Run(context.Background(), env, []string{"-r", "/ref.txt", targetFile})
	assert.Equal(t, 0, status)
	
	targetInfo, _ := fs.Stat(targetFile)
	assert.True(t, targetInfo.ModTime().Equal(refMtime))
}
