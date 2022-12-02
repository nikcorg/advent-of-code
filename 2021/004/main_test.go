package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	parsed, err := parseInput(testInput)

	assert.NoError(t, err)
	assert.Equal(t, 27, len(parsed.Nums))
	assert.Equal(t, 3, len(parsed.Boards))

	for _, b := range parsed.Boards {
		assert.Equal(t, 25, len(b.Nums))
	}
}

func TestSolveFirst(t *testing.T) {
	expected := 4512

	parsed, _ := parseInput(testInput)

	result, err := solveFirst(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSolveSecond(t *testing.T) {
	expected := 1924

	parsed, _ := parseInput(testInput)

	result, err := solveSecond(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func BenchmarkSolveFirst(b *testing.B) {
	parsed, _ := parseInput(testInput)

	for n := 0; n < b.N; n++ {
		solveFirst(parsed)
	}
}

func BenchmarkSolveSecond(b *testing.B) {
	parsed, _ := parseInput(testInput)

	for n := 0; n < b.N; n++ {
		solveSecond(parsed)
	}
}
