package main

import (
	"bufio"
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

func TestSolveFirst(t *testing.T) {
	v, err := solveFirst(testInput)
	assert.NoError(t, err)
	assert.Equal(t, "CMZ", v)
}
func TestSolveSecond(t *testing.T) {
	v, err := solveSecond(testInput)
	assert.NoError(t, err)
	assert.Equal(t, "MCD", v)
}

func TestParseStacks(t *testing.T) {
	stacks := parseStacks(bufio.NewScanner(strings.NewReader(testInput)))

	assert.Equal(t, 3, len(stacks))

	assert.Equal(t, 2, stacks[0].Size())
	assert.Equal(t, byte('N'), stacks[0].Peek(), "expected peek to return N")
	assert.Equal(t, []byte{'N', 'Z'}, stacks[0].PopN(2))

	assert.Equal(t, 3, stacks[1].Size())
	assert.Equal(t, byte('D'), stacks[1].Peek(), "expected peek to return D")
	assert.Equal(t, []byte{'D', 'C', 'M'}, stacks[1].PopN(3))

	assert.Equal(t, 1, stacks[2].Size())
	assert.Equal(t, byte('P'), stacks[2].Peek(), "expected peek to return P")
	assert.Equal(t, []byte{'P'}, stacks[2].PopN(1))
}
