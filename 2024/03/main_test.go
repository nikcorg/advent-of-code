package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testInput  string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	testInput2 string = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 161, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInput2)
	assert.Equal(t, 48, sol)
}
