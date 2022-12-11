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
	n := solveFirst(testInput)

	assert.Equal(t, 10605, n)
}

func TestSolveSecond(t *testing.T) {
	n := solveSecond(testInput)

	assert.Equal(t, 2713310158, n)
}
