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
	cs := parseInput(testInput)

	cases := []card{
		{ID: 1, Copies: 1, Winners: map[string]struct{}{"41": {}, "48": {}, "83": {}, "86": {}, "17": {}}, Draw: []string{"83", "86", "6", "31", "17", "9", "48", "53"}},
		{ID: 2, Copies: 1, Winners: map[string]struct{}{"13": {}, "32": {}, "20": {}, "16": {}, "61": {}}, Draw: []string{"61", "30", "68", "82", "17", "32", "24", "19"}},
		{ID: 3, Copies: 1, Winners: map[string]struct{}{"1": {}, "21": {}, "53": {}, "59": {}, "44": {}}, Draw: []string{"69", "82", "63", "72", "16", "21", "14", "1"}},
		{ID: 4, Copies: 1, Winners: map[string]struct{}{"41": {}, "92": {}, "73": {}, "84": {}, "69": {}}, Draw: []string{"59", "84", "76", "51", "58", "5", "54", "83"}},
		{ID: 5, Copies: 1, Winners: map[string]struct{}{"87": {}, "83": {}, "26": {}, "28": {}, "32": {}}, Draw: []string{"88", "30", "70", "12", "93", "22", "82", "36"}},
		{ID: 6, Copies: 1, Winners: map[string]struct{}{"31": {}, "18": {}, "13": {}, "56": {}, "72": {}}, Draw: []string{"74", "77", "10", "23", "35", "67", "36", "11"}},
	}

	assert.GreaterOrEqual(t, len(cs), len(cases))

	for i, c := range cases {
		assert.Equal(t, c, cs[i])
	}
}

func TestSolveFirst(t *testing.T) {
	cards := parseInput(testInput)
	solution := solveFirst(cards)
	assert.Equal(t, 13, solution)
}

func TestSolveSecond(t *testing.T) {
	cards := parseInput(testInput)
	solution := solveSecond(cards)
	assert.Equal(t, 30, solution)
}
