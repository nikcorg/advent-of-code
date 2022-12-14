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
	assert.Equal(t, 24, v)
}

func TestSolveSecond(t *testing.T) {
	v := solveSecond(testInput)
	assert.Equal(t, 93, v)
}
