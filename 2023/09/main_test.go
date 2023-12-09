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
	expect := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
	assert.Equal(t, expect, parseInput(testInput))
}

func TestSolveFirst(t *testing.T) {
	seqs := parseInput(testInput)
	sol := solveFirst(seqs)
	assert.Equal(t, 114, sol)
}

func TestSolveSecond(t *testing.T) {
	seqs := parseInput(testInput)
	sol := solveSecond(seqs)
	assert.Equal(t, 2, sol)
}
