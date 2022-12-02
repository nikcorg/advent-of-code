package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `D2FE28`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	parsed, err := parseInput(testInput)

	expected := []packet{
		{Version: 6, Type: 4, Value: 2021},
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, parsed)
}

func TestSolveFirst(t *testing.T) {
	expected := 0

	parsed, _ := parseInput(testInput)

	result, err := solveFirst(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSolveSecond(t *testing.T) {
	expected := 0

	parsed, _ := parseInput(testInput)

	result, err := solveSecond(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
