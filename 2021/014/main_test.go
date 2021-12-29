package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	_, err := parseInput(testInput)
	assert.NoError(t, err)
}

func TestSplitPolymer(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("NN"), []byte("NC"), []byte("CB")}, splitPolymer([]byte("NNCB")))
}

func TestSolveFirst(t *testing.T) {
	expected := uint64(1588)

	parsed, _ := parseInput(testInput)

	result, err := solveFirst(parsed)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSolveSecond(t *testing.T) {
	expected := uint64(2188189693529)

	parsed, _ := parseInput(testInput)

	result, err := solveSecond(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
