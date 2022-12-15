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
}

func TestParseInput(t *testing.T) {
	m := parseInput(bufio.NewScanner(strings.NewReader(testInput)))

	assert.Equal(t, 14, len(m)) // one pair for each line in the input
}
