package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `2199943210
3987894921
9856789892
8767896789
9899965678
`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	pi, err := parseInput(testInput)

	assert.NoError(t, err)
	assert.Equal(t, pi.height, 5)
	assert.Equal(t, pi.width, 10)
	assert.Len(t, pi.tiles, 10*5)
}

func TestPiAt(t *testing.T) {
	pi, _ := parseInput(testInput)
	v, err := pi.At(point{0, 0})
	assert.NoError(t, err)
	assert.Equal(t, v, 2)

	v, err = pi.At(point{9, 0})
	assert.NoError(t, err)
	assert.Equal(t, v, 0)

	v, err = pi.At(point{3, 2})
	assert.NoError(t, err)
	assert.Equal(t, 6, v)

	_, err = pi.At(point{0, 100})
	assert.Error(t, err)
}

func TestSolveFirst(t *testing.T) {
	expected := 15

	parsed, _ := parseInput(testInput)

	result, err := solveFirst(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSolveSecond(t *testing.T) {
	expected := 1134

	parsed, _ := parseInput(testInput)

	result, err := solveSecond(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
