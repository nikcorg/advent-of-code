package main

import (
	"bytes"
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testInput = `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`
	testInput2 = `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`
	testInput3 = `31meight5636eightxpd3
	1m9ninechrpncvqfone1
	nqdftnsevenonellvpsdhrnrtrdjhbqscpd78
	31mthree6sdnttwothree3
	threetxgc2htprtqqj5fouroneightlf`
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}

func TestSolveFirst(t *testing.T) {
	solution := solveFirst(strings.Split(testInput, "\n"))

	assert.Equal(t, 142, solution)
}

func TestSolveSecond(t *testing.T) {
	solution := solveSecond(strings.Split(testInput2, "\n"))

	assert.Equal(t, 281, solution)

	solution = solveSecond(strings.Split(testInput3, "\n"))

	assert.Equal(t, 33+11+78+33+38, solution)
}
