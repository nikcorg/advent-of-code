package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"nikc.org/aoc2024/util"
)

var (
	//go:embed input_test.txt
	testInput string
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}

func TestParseInput(t *testing.T) {
	m := parseInput(testInput)

	// 89010123
	// 78121874
	// 87430965
	// 96549874
	// 45678903
	// 32019012
	// 01329801
	// 10456732

	assert.Equal(t, 9, len(m.heads))
	assert.Equal(t, 8, m.width)
	assert.Equal(t, 8, m.height)
	assert.Equal(t, 64, len(m.grid))

	v, err := m.At(util.NewPoint(7, 7))
	assert.NoError(t, err)
	assert.Equal(t, 2, v)

	_, err = m.At(util.NewPoint(8, 8))
	assert.Error(t, err)

	v, err = m.At(util.NewPoint(0, 2))
	assert.NoError(t, err)
	assert.Equal(t, 8, v)

	v, err = m.At(util.NewPoint(0, 3))
	assert.NoError(t, err)
	assert.Equal(t, 9, v)
}

func TestVisit(t *testing.T) {
	m := parseInput(testInput)

	assert.Equal(t, 1, visit(m, 8, util.NewPoint(0, 2), util.NewSet[util.Point]()))
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInput)
	assert.Equal(t, 36, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInput)
	assert.Equal(t, 81, sol)
}
