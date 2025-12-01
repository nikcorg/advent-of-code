package main

import (
	"bytes"
	_ "embed"
	"fmt"
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
	inp := parseInput(testInput)
	exp := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	assert.Equal(t, exp, inp)
}

func TestSolveFirst(t *testing.T) {
	assert.Equal(t, 2, solveFirst(testInput))
}

func TestSolveSecond(t *testing.T) {
	assert.Equal(t, 4, solveSecond(testInput))
}

func TestDescending(t *testing.T) {
	ts := []struct {
		I []int
		S int
		E bool
	}{
		{[]int{7, 6, 4, 2, 1}, 1, true},  // safe (desc)
		{[]int{1, 2, 7, 8, 9}, 1, false}, // unsafe
		{[]int{9, 7, 6, 2, 1}, 1, false}, // unsafe
		{[]int{1, 3, 2, 4, 5}, 1, false}, // safe (asc)
		{[]int{8, 6, 4, 4, 1}, 1, true},  // safe (desc)
		{[]int{1, 3, 6, 7, 9}, 1, false}, // safe (asc)
	}

	for _, tc := range ts {
		assert.Equal(t, tc.E, descending(tc.I, tc.S), fmt.Sprintf("I=%v", tc.I))
	}
}

func TestAscending(t *testing.T) {
	ts := []struct {
		E bool
		S int
		I []int
	}{
		{false, 1, []int{7, 6, 4, 2, 1}}, // safe (desc)
		{false, 1, []int{1, 2, 7, 8, 9}}, // unsafe
		{false, 1, []int{9, 7, 6, 2, 1}}, // unsafe
		{true, 1, []int{1, 3, 2, 4, 5}},  // safe (asc)
		{false, 1, []int{8, 6, 4, 4, 1}}, // safe (desc)
		{true, 1, []int{1, 3, 6, 7, 9}},  // safe (asc)
		{true, 1, []int{1, 3, 6, 9, 8}},  // safe (asc)
		{true, 1, []int{1, 3, 5, 9, 8}},  // safe (asc)
	}

	for _, tc := range ts {
		assert.Equal(t, tc.E, ascending(tc.I, tc.S), fmt.Sprintf("I=%v", tc.I))
	}
}
