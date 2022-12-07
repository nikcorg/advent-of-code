package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed input_test.txt
	testInput string
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}

func TestSolveFirst(t *testing.T) {
	v, err := solveFirst(testInput)
	assert.NoError(t, err)
	assert.Equal(t, uint(95437), v)
}

func TestSolveSecond(t *testing.T) {
	v, err := solveSecond(testInput)
	assert.NoError(t, err)
	assert.Equal(t, uint(24933642), v)
}

func TestScanFS(t *testing.T) {
	sizes, used, err := scanFS(bufio.NewScanner(strings.NewReader(testInput)))

	assert.NoError(t, err)
	assert.Equal(t, uint(48381165), used)
	assert.Equal(t, 4, len(sizes))

	assert.Equal(t, uint(584), sizes["/a/e"])
	assert.Equal(t, uint(94853), sizes["/a"])
	assert.Equal(t, uint(24933642), sizes["/d"])
	assert.Equal(t, uint(48381165), sizes[""])
}
