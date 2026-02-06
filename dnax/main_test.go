package main

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tmpDir := os.TempDir()

	inputFile := "../testseq.txt"
	outputFile := filepath.Join(tmpDir, "testresults.txt")
	args := []string{
		"-input",
		inputFile,
		"-output",
		outputFile,
	}
	assert.NoError(t, run(args, io.Discard))

	got, err := os.ReadFile(outputFile)
	assert.NoError(t, err)
	assert.Equal(t, "", string(got))

}
