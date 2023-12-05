package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed input_test.txt
	testInput string
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}

func TestParseInput(t *testing.T) {
	seeds, maps := parseInput(testInput)

	assert.Equal(t, []int{79, 14, 55, 13}, seeds)

	assert.Equal(t, sourceToDestMap{
		Name: "seed-to-soil",
		Ranges: []mapRange{
			// 98-99 -> 50-51
			{Width: 2, DestFrom: 50, SourceFrom: 98},
			// 50-97 -> 52-99
			{Width: 48, DestFrom: 52, SourceFrom: 50},
		}}, maps[0])

	assert.Equal(t, 7, len(maps))
}

func TestMapRange(t *testing.T) {
	// 98-99 -> 50-51
	a := mapRange{2, 50, 98}
	// 50-97 -> 52-99
	b := mapRange{48, 52, 50}
	l := listOfRanges{a, b}

	for _, c := range []struct {
		From, To int
		Ok       bool
	}{
		{10, 10, false},
		{98, 50, true},
		{99, 51, true},
		{100, 100, false},
	} {
		next, ok := a.MapSource(c.From)
		assert.Equal(t, c.Ok, ok)
		assert.Equal(t, c.To, next)
	}

	// Seed number 79 corresponds to soil number 81.
	assert.Equal(t, 81, l.MapSource(79))
	// Seed number 14 corresponds to soil number 14.
	assert.Equal(t, 14, l.MapSource(14))
	// Seed number 55 corresponds to soil number 57.
	assert.Equal(t, 57, l.MapSource(55))
	// Seed number 13 corresponds to soil number 13
	assert.Equal(t, 13, l.MapSource(13))
}

func TestSolveFirst(t *testing.T) {
	seeds, maps := parseInput(testInput)
	solution := solveFirst(seeds, maps)
	assert.Equal(t, 35, solution)
}

func TestSolveSecond(t *testing.T) {
	seeds, maps := parseInput(testInput)
	solution := solveSecond(seeds, maps)
	assert.Equal(t, 46, solution)
}
