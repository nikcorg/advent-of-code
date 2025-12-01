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
	exp := [][]int{
		{190, 10, 19},
		{3267, 81, 40, 27},
		{83, 17, 5},
		{156, 15, 6},
		{7290, 6, 8, 6, 15},
		{161011, 16, 10, 13},
		{192, 17, 8, 14},
		{21037, 9, 7, 18, 13},
		{292, 11, 6, 16, 20},
	}

	i := parseInput(testInput)

	assert.Equal(t, exp, i)
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 3749, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInput)
	assert.Equal(t, 11387, sol)
}
