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
	games := parseInput(testInput)
	assert.Equal(t, [][]int{{7, 9}, {15, 40}, {30, 200}}, games)
}

func TestSolveFirst(t *testing.T) {
	games := parseInput(testInput)
	solution := solveFirst(games)
	assert.Equal(t, 288, solution)
}

func TestSolveSecond(t *testing.T) {
	games := parseInput(testInput)
	solution := solveSecond(games)
	assert.Equal(t, 71503, solution)
}

func TestZip(t *testing.T) {
	as := []int{1, 2, 3}
	bs := []int{4, 5, 6, 7}

	cs := zip(as, bs)

	assert.Equal(t, [][]int{{1, 4}, {2, 5}, {3, 6}}, cs)
}
