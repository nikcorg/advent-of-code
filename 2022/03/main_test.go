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

func TestSolveFirst(t *testing.T) {
	v, err := solveFirst(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 157, v)
}

func TestSolveSecond(t *testing.T) {
	v, err := solveSecond(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 70, v)
}

func TestSplit(t *testing.T) {
	cases := []struct{ i, l, r string }{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
	}

	for _, c := range cases {
		l, r := split(c.i)
		assert.Equal(t, c.l, l)
		assert.Equal(t, c.r, r)
	}
}

func TestPriority(t *testing.T) {
	cases := []struct {
		ch byte
		p  int
	}{
		{'a', 1},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
	}

	for _, c := range cases {
		assert.Equal(t, c.p, priority(c.ch))
	}
}

func TestOverlap(t *testing.T) {
	cases := []struct {
		l, r, o string
	}{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp", "p"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL", "L"},
	}

	for _, c := range cases {
		assert.Equal(t, c.o, overlap(c.l, c.r))
	}
}
