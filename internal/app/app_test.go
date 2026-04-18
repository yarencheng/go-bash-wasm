package app

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_ShowVersion(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	// Create a minimal environment for testing
	a := New(nil, &stdout, &stderr)

	a.ShowVersion()

	output := stdout.String()
	assert.Contains(t, output, "go-bash-wasm, version 5.3-rc")
	assert.Contains(t, output, "Copyright (C) 2026 go-bash-wasm team")
	assert.Contains(t, output, "License Apache-2.0: Apache License, Version 2.0")
}
