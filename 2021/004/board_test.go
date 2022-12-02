package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	parsed, _ := parseInput(testInput)

	b := parsed.Boards[0]

	assert.Equal(t, 11, b.At(point{3, 0}))
	assert.Equal(t, 9, b.At(point{1, 2}))

	assert.False(t, b.Winner())

	for _, n := range []int{11, 4, 16, 18, 15} {
		b.Mark(n)
	}

	assert.True(t, b.Winner())
}

func TestBoardMore(t *testing.T) {
	parsed, _ := parseInput(testInput)

	b := parsed.Boards[2]

	for _, n := range parsed.Nums[0:12] {
		b.Mark(n)
	}

	for i, n := range []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24} {
		assert.Equal(t, n, b.At(b.Hits[i]))
	}

	assert.True(t, b.Winner())
}
