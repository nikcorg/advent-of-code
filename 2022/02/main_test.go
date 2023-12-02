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

func TestFirst(t *testing.T) {
	cases := []struct {
		input string
		exp   int
	}{
		{testInput, 15},

		{"A X", 3 + 1},
		{"A Y", 6 + 2},
		{"A Z", 0 + 3},

		{"B X", 0 + 1},
		{"B Y", 3 + 2},
		{"B Z", 6 + 3},

		{"C X", 6 + 1},
		{"C Y", 0 + 2},
		{"C Z", 3 + 3},
	}

	for _, tc := range cases {
		v, err := solveFirst(tc.input)
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, v)
	}
}

func TestSecond(t *testing.T) {
	cases := []struct {
		input string
		exp   int
	}{
		{testInput, 12},

		{"A X", 0 + 3},
		{"B X", 0 + 1},
		{"C X", 0 + 2},

		{"A Y", 3 + 1},
		{"B Y", 3 + 2},
		{"C Y", 3 + 3},

		{"A Z", 6 + 2},
		{"B Z", 6 + 3},
		{"C Z", 6 + 1},
	}

	for _, tc := range cases {
		v, err := solveSecond(tc.input)
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, v)
	}
}
