package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string = `199
200
208
210
200
207
240
269
260
263
`

func TestParseInput(t *testing.T) {
	parsed, err := parseInput(testInput)

	expected := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	assert.NoError(t, err)
	assert.ElementsMatch(t, expected, parsed)
}

func TestCountIncrements(t *testing.T) {
	input, _ := parseInput(testInput)
	incs := countIncrements(input)

	assert.Equal(t, 7, incs)
}

func TestCountTripletIncrements(t *testing.T) {
	input, _ := parseInput(testInput)
	incs := countTripletIncrements(input)

	assert.Equal(t, 5, incs)
}
