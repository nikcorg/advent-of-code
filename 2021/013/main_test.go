package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	pi, err := parseInput(testInput)

	assert.NoError(t, err)
	assert.Len(t, pi.Dots, 18)
	assert.Len(t, pi.Folds, 2)
}

func TestSolveFirst(t *testing.T) {
	expected := 17

	parsed, _ := parseInput(testInput)

	result, err := solveFirst(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSolveSecond(t *testing.T) {
	expected := `#####
#...#
#...#
#...#
#####
.....
.....
`

	parsed, _ := parseInput(testInput)

	result, err := solveSecond(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
