package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	bits, parsed, err := parseInput(testInput)
	assert.NoError(t, err)
	assert.Equal(t, uint(5), bits)
	assert.Len(t, parsed, 12)
}

func TestOnesCount(t *testing.T) {

	bits, parsed, _ := parseInput(testInput)
	oc := onesCount(bits, parsed)

	assert.Equal(t, []int{7, 5, 8, 7, 5}, oc)
}

func TestSolveFirst(t *testing.T) {
	bits, parsed, _ := parseInput(testInput)
	solution, err := solveFirst(bits, parsed)
	assert.NoError(t, err)
	assert.Equal(t, 198, solution)
}
func TestKeep(t *testing.T) {
	bits, parsed, _ := parseInput(testInput)
	// filter on "first" (left-most) bit
	filtered := keep(parsed, 1<<(bits-1))

	assert.Equal(t, 7, len(filtered))
}

func TestDiscard(t *testing.T) {
	bits, parsed, _ := parseInput(testInput)
	// filter on "first" (left-most) bit
	filtered := discard(parsed, 1<<(bits-1))

	assert.Equal(t, 5, len(filtered))
}

func TestSolveSecond(t *testing.T) {
	bits, parsed, _ := parseInput(testInput)
	solution, err := solveSecond(bits, parsed)
	assert.NoError(t, err)
	assert.Equal(t, 230, solution)
}
