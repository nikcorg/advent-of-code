package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	testInput2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}

func TestSolveFirst(t *testing.T) {
	v, err := solveFirst(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 13, v)
}

func TestSolveSecond(t *testing.T) {
	cases := []struct {
		input string
		exp   int
	}{
		{testInput, 1},
		{testInput2, 36},
	}
	for _, tc := range cases {
		v, err := solveSecond(tc.input)
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, v)
	}
}
