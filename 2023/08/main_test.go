package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testInput string = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
	testInput2 string = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}

func TestParseInput(t *testing.T) {
	cursor, nodes := parseInput(testInput)
	for _, c := range []string{"L", "L", "R", "L", "L", "R"} {
		assert.Equal(t, c, cursor.C())
		cursor.Adv()
	}

	assert.Equal(t, nodemap{
		"AAA": {"BBB", "BBB"},
		"BBB": {"AAA", "ZZZ"},
		"ZZZ": {"ZZZ", "ZZZ"},
	}, nodes)
}

func TestSolveFirst(t *testing.T) {
	c, nm := parseInput(testInput)
	solution := solveFirst(c, nm)
	assert.Equal(t, 6, solution)
}

func TestSolveSecond(t *testing.T) {
	t.Log("using input 2")
	c, nm := parseInput(testInput2)
	solution := solveSecond(c, nm)
	assert.Equal(t, 6, solution)

	t.Log("using input 1")
	c, nm = parseInput(testInput)
	solution = solveSecond(c, nm)
	assert.Equal(t, 6, solution)
}
