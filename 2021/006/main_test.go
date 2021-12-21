package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `3,4,3,1,2`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	nums, err := parseInput(testInput)
	assert.NoError(t, err)
	assert.Len(t, nums, 5)
}

func TestSolveFirst(t *testing.T) {
	expected := 5934

	parsed, _ := parseInput(testInput)

	result, err := solveFirst(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSolveSecond(t *testing.T) {
	expected := 26984457539

	parsed, _ := parseInput(testInput)

	result, err := solveSecond(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
