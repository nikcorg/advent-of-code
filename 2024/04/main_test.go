package main

import (
	"bytes"
	_ "embed"
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

func TestParseInput(t *testing.T) {
	m, w := parseInput(testInput)

	assert.Equal(t, 100, len(m))
	assert.Equal(t, 10, w)
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 18, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInput)
	assert.Equal(t, 9, sol)
}
