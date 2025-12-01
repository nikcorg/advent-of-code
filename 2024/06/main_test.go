package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"nikc.org/aoc2024/util"
)

var (
	//go:embed input_test.txt
	testInput string
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}

func TestParseInput(t *testing.T) {
	m := parseInput(testInput)

	assert.Equal(t, 8, m.objects.Size())
	assert.Equal(t, 10, m.width)
	assert.Equal(t, 10, m.height)
	assert.Equal(t, util.NewPoint(4, 6), m.gLoc)

	// ....#.....
	// .........#
	// ..........
	// ..#.......
	// .......#..
	// ..........
	// .#..^.....
	// ........#.
	// #.........
	// ......#...
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 41, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInput)
	assert.Equal(t, 0, sol)
}
