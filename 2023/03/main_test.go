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
	m := parseInput(testInput)
	assert.Equal(t, 10, m.Width)
	assert.Equal(t, 100, len(m.Map))
}

func TestSolveFirst(t *testing.T) {
	m := parseInput(testInput)
	solution := solveFirst(m)
	assert.Equal(t, 4361, solution)
}

func TestSolveSecond(t *testing.T) {
	m := parseInput(testInput)
	solution := solveSecond(m)
	assert.Equal(t, 467835, solution)
}
