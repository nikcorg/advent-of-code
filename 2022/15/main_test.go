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
	v := solveFirst(testInput, 10)
	assert.Equal(t, 26, v)

	v = solveFirst(input, 2_000_000)
	assert.Equal(t, 6124805, v)
}

func TestSolveSecond(t *testing.T) {
	v := solveSecond(testInput, 0, 20)
	assert.Equal(t, 56000011, v)

	v = solveSecond(input, 0, 4_000_000)
	assert.Equal(t, 12555527364986, v)
}

func TestParseInput(t *testing.T) {
	m, _ := parseInput(bufio.NewScanner(strings.NewReader(testInput)))

	assert.Equal(t, 14, len(m)) // one pair for each line in the input
}
