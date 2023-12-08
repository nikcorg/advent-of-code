package main

import (
	"bytes"
	_ "embed"
	"strings"
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
	tests := []struct {
		Input  string
		Expect round
	}{
		{
			Input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`,
			Expect: round{
				ID:     1,
				Reveal: []game{{Blue: 3, Red: 4}, {Red: 1, Green: 2, Blue: 6}, {Green: 2}},
			},
		},
		{
			Input: `Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue`,
			Expect: round{
				ID:     2,
				Reveal: []game{{Blue: 1, Green: 2}, {Green: 3, Blue: 4, Red: 1}, {Green: 1, Blue: 1}},
			},
		},
		{
			Input: `Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red`,
			Expect: round{
				ID:     3,
				Reveal: []game{{Green: 8, Blue: 6, Red: 20}, {Blue: 5, Red: 4, Green: 13}, {Green: 5, Red: 1}},
			},
		},
		{
			Input: `Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red`,
			Expect: round{
				ID:     4,
				Reveal: []game{{Green: 1, Red: 3, Blue: 6}, {Green: 3, Red: 6}, {Green: 3, Blue: 15, Red: 14}},
			},
		},
		{
			Input: `Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			Expect: round{
				ID:     5,
				Reveal: []game{{Red: 6, Blue: 1, Green: 3}, {Blue: 2, Red: 1, Green: 2}},
			},
		},
	}

	for _, tc := range tests {
		out := parseInput([]string{tc.Input})
		assert.Equal(t, tc.Expect, out[0])
	}
}

func TestSolveFirst(t *testing.T) {
	input := parseInput(strings.Split(testInput, "\n"))
	setup := game{Red: 12, Green: 13, Blue: 14}
	solution := solveFirst(setup, input)
	assert.Equal(t, 8, solution)
}

func TestSolveSecond(t *testing.T) {
	input := parseInput(strings.Split(testInput, "\n"))
	solution := solveSecond(input)
	assert.Equal(t, 2286, solution)
}
