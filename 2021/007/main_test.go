package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `16,1,2,0,4,2,7,1,2,14`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	positions, err := parseInput(testInput)

	assert.NoError(t, err)
	assert.Len(t, positions, 10)
}

func TestSolveFirst(t *testing.T) {
	expected := 37

	parsed, _ := parseInput(testInput)

	result, err := solveFirst(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSolveSecond(t *testing.T) {
	expected := 168

	parsed, _ := parseInput(testInput)

	result, err := solveSecond(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
