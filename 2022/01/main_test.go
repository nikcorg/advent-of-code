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
	calories, err := solveFirst(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 24000, calories)
}

func TestSolveSecond(t *testing.T) {
	calories, err := solveSecond(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 45000, calories)
}
