package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"nikc.org/aoc2022/12/dijkstra"
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
	assert.Equal(t, 31, v)
}

func TestSolveSecond(t *testing.T) {
	v, err := solveSecond(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 29, v)
}

func TestNewMap(t *testing.T) {
	m, err := newMap(bufio.NewScanner(strings.NewReader(testInput)))

	assert.NoError(t, err)
	assert.Equal(t, 8, m.Width())
	assert.Equal(t, dijkstra.NewPoint(0, 0), m.start)
	assert.Equal(t, dijkstra.NewPoint(5, 2), m.end)

	// x underflow
	_, err = m.At(dijkstra.NewPoint(-1, 0))
	assert.Error(t, err)

	// y underflow
	_, err = m.At(dijkstra.NewPoint(0, -1))
	assert.Error(t, err)

	// x overflow
	_, err = m.At(dijkstra.NewPoint(m.Width(), 0))
	assert.Error(t, err)

	// y overflow
	_, err = m.At(dijkstra.NewPoint(0, m.Height()))
	assert.Error(t, err)
}
