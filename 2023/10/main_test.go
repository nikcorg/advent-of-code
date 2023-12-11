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
	parseInput(testInput)
}

func TestSolveFirst(t *testing.T) {
	i := parseInput(testInput)
	sol := solveFirst(i)
	assert.Equal(t, 8, sol)
}

func TestSolveSecond(t *testing.T) {
	i := parseInput(testInput)
	sol := solveSecond(i)
	assert.Equal(t, 0, sol)
}
