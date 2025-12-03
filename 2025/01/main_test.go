package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed input_test.txt
	testInput string

	//go:embed input.txt
	realInput string
)

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 3, sol)

	sol = solveFirst(realInput)
	assert.Equal(t, 1132, sol)
}

var (
	testCases = []struct {
		input    string
		expected int
	}{
		{`L50`, 1}, // 0
		{`L50
L1`, 1}, // 1
		{`R49`, 0}, // 2
		{`R50`, 1}, // 3
		{`R49
R1`, 1}, // 4
		{`R49
R1
R100`, 2}, // 5
		{`R49
R1
R203`, 3}, // 6
		{`L55
L95
L200`, 4}, // 7
		{testInput, 6},    // 8
		{realInput, 6623}, // 9
	}
)

func TestSolveSecond(t *testing.T) {
	for i, tc := range testCases {
		sol := solveSecond(tc.input)
		assert.Equal(t, tc.expected, sol, fmt.Sprintf("failed for %d", i))
	}
}
