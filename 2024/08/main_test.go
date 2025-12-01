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

	assert.Equal(t, 12, m.Width)
	assert.Equal(t, 12, m.Height)

	as := [][]int{
		{1, 8, '0'},
		{2, 5, '0'},
		{3, 7, '0'},
		{4, 4, '0'},
		{5, 6, 'A'},
		{8, 8, 'A'},
		{9, 9, 'A'},
	}

	for i, a := range as {
		p := util.NewPoint(a[1], a[0])
		assert.True(t, p.Equals(m.Antennae[i].Pos))
		assert.Equal(t, a[2], int(m.Antennae[i].Freq))
	}
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 14, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInput)
	assert.Equal(t, 34, sol)
}
