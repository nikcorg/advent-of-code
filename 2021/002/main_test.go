package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr())
}

const testInput = `forward 5
down 5
forward 8
up 3
down 8
forward 2
`

func TestParseInput(t *testing.T) {
	parsed, err := parseInput(testInput)

	assert.NoError(t, err)
	assert.Len(t, parsed, 6)

	for _, cmd := range parsed {
		assert.NotEqual(t, unset, cmd.dir)
	}
}

func TestSolveFirst(t *testing.T) {
	solution, err := solveFirst(testInput)

	assert.NoError(t, err)
	assert.Equal(t, 150, solution)
}

func TestSolveSecond(t *testing.T) {
	solution, err := solveSecond(testInput)

	assert.NoError(t, err)
	assert.Equal(t, 900, solution)
}
