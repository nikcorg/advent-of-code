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
	left, right := parseInput(testInput)

	expRight := []int{4, 3, 5, 3, 9, 3}
	expLeft := []int{3, 4, 2, 1, 3, 3}

	assert.Equal(t, expLeft, left)
	assert.Equal(t, expRight, right)
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 11, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInput)
	assert.Equal(t, 31, sol)
}
