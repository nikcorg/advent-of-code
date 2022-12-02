package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testInput string = ``

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	_, err := parseInput(testInput)
	assert.NoError(t, err)
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
