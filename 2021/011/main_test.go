package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	pi, err := parseInput(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 10, pi.Width)
	assert.Equal(t, 10, pi.Height)
	assert.Len(t, pi.Grid, 100)
}

func TestSolveFirst(t *testing.T) {
	expected := 1656

	parsed, _ := parseInput(testInput)

	result, err := solveFirst(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSolveSecond(t *testing.T) {
	expected := 195

	parsed, _ := parseInput(testInput)

	result, err := solveSecond(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
