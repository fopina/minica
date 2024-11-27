package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to capture output
func runMain(args ...string) (string, string, error) {
	var buf bytes.Buffer
	var buf2 bytes.Buffer
	originalArgs := os.Args
	originalStdout := os.Stdout

	r, w, _ := os.Pipe()
	r2, w2, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w2
	os.Args = args

	err := main2()

	w.Close()
	w2.Close()
	io.Copy(&buf, r)
	io.Copy(&buf2, r2)

	// restore
	os.Stdout = originalStdout
	os.Args = originalArgs

	return buf.String(), buf2.String(), err
}

func TestParseArgs(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdir")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	os.Chdir(tempDir)
	stdout, stderr, err := runMain("cmd", "-domains=test.example.com")
	assert.NoError(t, err)
	assert.Equal(t, "", stdout)
	assert.Equal(t, "", stderr)

	filesToCheck := []string{
		"minica.pem",
		"minica-key.pem",
		"test.example.com/cert.pem",
		"test.example.com/key.pem",
	}

	for _, fileToCheck := range filesToCheck {
		_, err = os.Stat(tempDir + "/" + fileToCheck)
		assert.NoError(t, err, fileToCheck+" not created")
	}
}
