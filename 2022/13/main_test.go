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

func TestSolveFirst(t *testing.T) {
	v := solveFirst(testInput)
	assert.Equal(t, 13, v)

	v = solveFirst(input)
	assert.Equal(t, 6428, v)
}

func TestSolveSecond(t *testing.T) {
	v := solveSecond(testInput)
	assert.Equal(t, 140, v)

	v = solveSecond(input)
	assert.Equal(t, 22464, v)
}
