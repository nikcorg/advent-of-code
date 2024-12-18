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
	order, runs := parseInput(testInput)

	_ = order

	assert.Equal(t, order, map[int]map[int]struct{}{
		47: {13: struct{}{}, 29: struct{}{}, 53: struct{}{}, 61: struct{}{}},
		97: {13: struct{}{}, 29: struct{}{}, 47: struct{}{}, 53: struct{}{}, 61: struct{}{}, 75: struct{}{}},
		75: {13: struct{}{}, 29: struct{}{}, 47: struct{}{}, 53: struct{}{}, 61: struct{}{}},
		61: {13: struct{}{}, 29: struct{}{}, 53: struct{}{}},
		29: {13: struct{}{}},
		53: {13: struct{}{}, 29: struct{}{}},
	})

	assert.Equal(t, [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}, runs)
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 143, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInput)
	assert.Equal(t, 123, sol)
}
