package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInput string = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(testInput))
}

func TestParseInput(t *testing.T) {
	lines, err := parseInput(testInput)
	assert.NoError(t, err)
	assert.Len(t, lines, 10)
}

func TestSolveFirst(t *testing.T) {
	expected := 26397

	parsed, _ := parseInput(testInput)

	result, err := solveFirst(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestSolveSecond(t *testing.T) {
	expected := 288957

	parsed, _ := parseInput(testInput)

	result, err := solveSecond(parsed)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
