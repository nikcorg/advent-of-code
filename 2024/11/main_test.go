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
	exp := []int64{125, 17}
	ns := parseInput(testInput)
	assert.Equal(t, exp, ns)
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 55312, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInput)
	assert.Equal(t, int64(55312), sol)
}
